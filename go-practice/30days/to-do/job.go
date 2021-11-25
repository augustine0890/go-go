package main

import "fmt"

type Job struct {
	State string
	done  chan struct{}
}

func (job *Job) Wait() {
	<-job.done
}

func (job *Job) Done() {
	job.State = "done"
	close(job.done)
}

func main() {
	// a channel is a pointer-like type
	ch := make(chan *Job)
	go func() {
		j := <-ch
		j.Done()
	}()

	job := Job{"ready", make(chan struct{})}
	ch <- &job
	job.Wait()
	fmt.Println(job.State)
}
