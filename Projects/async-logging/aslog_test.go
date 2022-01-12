package aslog

import (
	"bytes"
	"errors"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

const messageTimestampPattern = `\[\d{4}-\d{2}-\d{2}\ \d{2}:\d{2}:\d{2}] - `

// 01
func TestMessageChannel(t *testing.T) {
	alog := New(nil)
	if alog.msgCh == nil {
		t.Fatal("msgCh field not initialized. Should have type 'chan string' but it is currently nil")
	}
}

// 02
func TestErrorChannel(t *testing.T) {
	alog := New(nil)
	if alog.errorCh == nil {
		t.Fatal("errorCh field not initialized. Should have type 'chan error' but it is currently nil")
	}
}

// 03
func TestMessageChannelMethod(t *testing.T) {
	alog := New(nil)
	if alog.MessageChannel() != alog.msgCh {
		t.Fatal("MessageChannel method does not return the msgCh field")
	}
	// messageChannelDir := reflect.ValueOf(alog.MessageChannel()).Type().ChanDir()
	// if messageChannelDir != reflect.SendDir {
	// t.Fatal("MessageChannel does not return send-only channel")
	// }
}

// 04
func TestErrorChannelMethod(t *testing.T) {
	alog := New(nil)
	if alog.ErrorChannel() != alog.errorCh {
		t.Fatal("ErrorChannel method does not return the errorCh field")
	}
	// errorChannelDir := reflect.ValueOf(alog.ErrorChannel()).Type().ChanDir()
	// if errorChannelDir != reflect.RecvDir {
	// t.Fatal("ErrorChannel does not return receive-only channel")
	// }
}

// 05
func TestWritesToWriter(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(b)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	alog.write("test", wg)

	written := b.String()
	if written == "" {
		t.Fatal("Nothing written to log")
	}
	if !regexp.MustCompile(messageTimestampPattern + "test\n$").Match([]byte(written)) {
		t.Error("Properly formatted string not written to log. Did you pass the message to 'formatMessage'?")
	}
}

// 06
type errorWriter struct {
	b *bytes.Buffer
}

func (ew errorWriter) Write(data []byte) (int, error) {
	ew.b.Write(data)
	return 0, errors.New("error")
}
func TestWriteSendsErrorsToErrorChannel(t *testing.T) {
	alog := New(&errorWriter{bytes.NewBuffer([]byte{})})
	alog.errorCh = make(chan error, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	alog.write("test", wg)
	go func() {
		if (<-alog.errorCh).Error() != "error" {
			t.Error("Did not receive destination writer's error on errorCh")
		}
	}()
	time.Sleep(100 * time.Millisecond)
	alog.errorCh <- errors.New("")
	time.Sleep(100 * time.Millisecond)
}

// 07
type sleepingWriter struct {
	b *bytes.Buffer
}

func (sw sleepingWriter) Write(data []byte) (int, error) {
	sw.b.Write(data)
	time.Sleep(500 * time.Millisecond)
	sw.b.WriteString("write complete")
	return 0, nil
}

func TestStartHandlesMessages(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(sleepingWriter{b})
	alog.msgCh = make(chan string, 2)
	go alog.Start()
	alog.msgCh <- "test message"
	time.Sleep(100 * time.Millisecond)
	written := b.Bytes()
	// Code calling t.Fatal (or a similar method) from a created goroutine should be rewritten to signal the test failure using
	// t.Error and exit the goroutine early using an alternative method, such as using a return statement.
	if !regexp.MustCompile(messageTimestampPattern + "test message\n$").Match(written) {
		t.Error("Message not written to logger's destination")
		return
	}
	shouldRelock := false
	if alog.m != nil {
		mutexState := reflect.ValueOf(*alog.m).FieldByName("state").Int()
		if mutexState != 0 {
			alog.m.Unlock()
			shouldRelock = true
		}
	}
	alog.msgCh <- "second message"
	time.Sleep(100 * time.Millisecond)
	if alog.m != nil {
		if shouldRelock {
			alog.m.Lock()
		}
	}
	written = b.Bytes()
	if !regexp.MustCompile(messageTimestampPattern + "test message\n" + messageTimestampPattern + "second message\n").Match(written) {
		t.Error("write method not called as a goroutine")
	}
}

// 08
type panickingWriter struct {
	b *bytes.Buffer
}

func (pw panickingWriter) Write(data []byte) (int, error) {
	pw.b.Write(data)
	panic("panicking!")
}
func TestWriteSendsWriteRequestsSequentially(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(sleepingWriter{b})
	if alog.m == nil {
		t.Fatal("Alog's mutex field 'm' not initialized")
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go alog.write("test message", wg)
	time.Sleep(100 * time.Millisecond)
	go alog.write("second message", wg)
	time.Sleep(1000 * time.Millisecond)
	written := b.Bytes()
	if !regexp.MustCompile(messageTimestampPattern + "test message\nwrite complete" + messageTimestampPattern + "second message\n").Match(written) {
		t.Error("Mutex not protecting Alog.dest#Write from concurrent calls")
	}

	b = bytes.NewBuffer([]byte{})
	alog = New(panickingWriter{b})
	if alog.msgCh == nil {
		t.Fatal("msgCh field is nil")
	}
	if alog.m == nil {
		t.Fatal("mutex field 'm' is nil")
	}
	go func() {
		defer func() {
			recover()
		}()
		alog.write("test message", wg)
	}()
	time.Sleep(100 * time.Millisecond)
	go func() {
		defer func() {
			recover()
		}()
		alog.write("second message", wg)
	}()
	time.Sleep(1000 * time.Millisecond)
	written = b.Bytes()
	if !regexp.MustCompile(messageTimestampPattern + "test message\n" + messageTimestampPattern + "second message\n").Match(written) {
		t.Error("Mutex not unlocked when panicking")
	}
}

// 09
func TestWriteSendsErrorsAsynchronously(t *testing.T) {
	TestWriteSendsWriteRequestsSequentially(t)
	b := bytes.NewBuffer([]byte{})
	alog := New(&errorWriter{b})
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go alog.write("first", wg)
	time.Sleep(100 * time.Millisecond)
	go alog.write("second", wg)
	time.Sleep(100 * time.Millisecond)
	written := b.Bytes()
	if !regexp.MustCompile(`.*first.*\n.*second.*`).Match(written) {
		t.Fatal("Error messages not sent to error channel asynchronously")
	}
	errorReceived := false
	go func() {
		<-alog.errorCh
		errorReceived = true
	}()
	time.Sleep(100 * time.Millisecond)
	if !errorReceived {
		t.Fatal("Error messages not sent to error channel")
	}
}

func TestNewInitializesShutdownChannels(t *testing.T) {
	alog := New(nil)
	if alog.shutdownCh == nil {
		t.Error("shutdownCh field not initialized")
	}

	if alog.shutdownCompleteCh == nil {
		t.Error("shutdownCompleteCh field not initialized")
	}
}

func TestShutdownMethod(t *testing.T) {
	alog := New(nil)
	alog.shutdownCompleteCh = make(chan struct{}, 1)
	alog.shutdown()
	time.Sleep(100 * time.Millisecond)
	select {
	case _, ok := <-alog.msgCh:
		if ok {
			t.Error("msgCh not closed by shutdown() method")
		}
	default:
		t.Error("msgCh not closed by shutdown() method")
	}
	select {
	case <-alog.shutdownCompleteCh:
	default:
		t.Error("shutdown() doesn't send message to shutdownCompleteCh")
	}
}

func TestStartMethodCallsShutdown(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(b)
	alog.shutdownCh = make(chan struct{}, 1)
	alog.shutdownCompleteCh = make(chan struct{}, 1)
	go alog.Start()
	alog.shutdownCh <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	select {
	case _, ok := <-alog.msgCh:
		if ok {
			t.Error("Passing message to shutdownCh doesn't call shutdown()")
		}
	default:
		t.Error("Passing message to shutdownCh doesn't call shutdown()")
	}
	select {
	case <-alog.shutdownCompleteCh:
	default:
		t.Error("Passing message to shutdownCh doesn't call shutdown()")
	}
	if b.Len() != 0 {
		t.Error("Passing message to shutdownCh doesn't break out of the Start method's for loop. " +
			"Note that 'break' statements can be used for select and for loops so a label might be " +
			"required to break out the loop.")
	}
}

func TestStopMethod(t *testing.T) {
	alog := New(nil)
	alog.shutdownCh = make(chan struct{}, 1)
	alog.shutdownCompleteCh = make(chan struct{}, 1)
	alog.shutdownCompleteCh <- struct{}{}
	alog.Stop()
	select {
	case <-alog.shutdownCh:
	default:
		t.Error("Stop() method doesn't send signal to shutdownCh channel")
	}
	select {
	case <-alog.shutdownCompleteCh:
		t.Error("Stop() method doesn't wait for signal from shutdownCompleteCh channel")
	default:
	}
}

func TestWriteAllBeforeShutdown(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(sleepingWriter{b})
	alog.msgCh = make(chan string, 2)
	go alog.Start()
	alog.msgCh <- "first"
	alog.msgCh <- "second"
	time.Sleep(10 * time.Millisecond)
	doneCh := make(chan struct{})
	go func() {
		alog.Stop()
		written := b.String()
		if !strings.Contains(written, "first") || !strings.Contains(written, "second") {
			t.Error("Not all messages written before logger shutdown")
		}
		doneCh <- struct{}{}
	}()
	select {
	case <-time.Tick(1 * time.Second):
		t.Error("Test timed out, please check that the Done method on the wait group is being called in the write method")
	case <-doneCh:
	}
}
