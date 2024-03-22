package main

import (
	// "io"
	// "os"
	// "strings"

	// "fmt"
	// "go-sample-codes/functions/clidemo"
	// "time"
	// "go-sample-codes/functions/simpleserver"
	"go-sample-codes/functions"
)

func main() {
	// fmt.Println("Hi, Mom!")
	// fmt.Println("The time is ", time.Now())
	// fmt.Println(functions.Sqrt(3))
	// simpleserver.MyServer()
	// clidemo.CliDemo()

	//rot13Reader exercise
	// s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// r := functions.Rot13Reader{R: s}
	// io.Copy(os.Stdout, &r)

	//Typed parameter functions example
	// Index works on a slice of ints
	// si := []int{10, 20, 15, -10}
	// fmt.Println(functions.Index(si, 15))

	// Index also works on a slice of strings
	// ss := []string{"foo", "bar", "baz"}
	// fmt.Println(functions.Index(ss, "hello"))

	//generic type linked list implementation example
	head := functions.List[int32]{Next: nil, Val: 1}
	head.Next = &functions.List[int32]{Next: nil, Val: 2}
	head.Display()

}
