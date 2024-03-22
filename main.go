package main

import (
	"io"
	"os"
	"strings"

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
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := functions.Rot13Reader{R: s}
	io.Copy(os.Stdout, &r)

}
