package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

// How do we run the code? go command line
// package == project == workspace
// all files in main need to have package main at the top
// two types of packages
// 	executable - when compiled, generates a file that we can run
// package main
// 	defines a package that can be compiled and then executed
// 	must have a fcn called main
// 	reusable - code used as helpers, good place to put reusable logic
// 		package calculator - defines a package that can be used as a dependency (helper code)
// 		package uploader - defines a package that can be used as a dependency (helper code)

// package main > go build > main.exe

// the word main is used to make an executable package
// 	main is important, if we named it something else, it would not be an executable

// fmt is short for format; used to form a link to other packages
