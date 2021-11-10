package main

import (
	"bufio"
	"fmt"
	"os"
)

type circle struct {
	radius float32
	border float32
}

func main() {
	fmt.Print("Go Standard Library\n")

	fmt.Println("This is end of the line")

	var answer int = 23
	fmt.Println("The correct answer is", answer)

	items := []int{5, 10, 15, 20}
	fmt.Println("The items are:", items)
	fmt.Printf("The items are:%v\n", items)

	x := 20
	f := 123.456789
	fmt.Printf("x: %d\n", x)
	fmt.Printf("x: %x\n", x) // base 16
	fmt.Printf("f: %f\n", f)
	fmt.Printf("%t\n", x > 10)

	fmt.Printf("%[2]d %[1]d\n", 55, 45)

	c := circle{radius: 20.0, border: 121.2}
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)

	// Sprintf returns a string
	s := fmt.Sprintf("%[2]d %[1]d\n", 55, 45)
	fmt.Printf("%s", s)

	fmt.Printf("%.2f\n", f)
	fmt.Printf("%10f\n", f)   // width 10 precision
	fmt.Printf("%10.2f\n", f) // padding and precision
	fmt.Printf("%+10.2f\n", f)
	fmt.Printf("%010.2f\n", f)

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Printf("%s\n", str)
}
