package main


import (
	"fmt"

	"golang.org/x/example/hello/reverse"

)

// initial version of the main function, 
// without the use of numbers reversal
// func main() {
// 	fmt.Println(reverse.String("Hello"))
// }

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
