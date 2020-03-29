package main

import "fmt"

// Functions have a name, list of parameters, optional list of results and a body.
// func <name> (<parameters>) <results>{
//	<body>
//}

// Functions Have A Type And Can It Is Defined By Its Signature.
// Go Support First Class Functions As Well (Can Declare A Function As A Variable).

// This Function Main, Also Known As The Entry Point Has No Parameters Or Results
func main() {
	WithParameters(1, "With Parameters Only")

	a, b := WithParametersAndResults(2, "With Parameters and 2 Returns")
	fmt.Println(a, b)

	a, b = WithParametersAndNamedResults(3, "With Parameters and 2 Named Returns")
	fmt.Println(a, b)
}

func WithParameters(a int, b string) {
	fmt.Println(a, b)
}

// Thi Function Has 2 Parameters and Returns 2 Results
func WithParametersAndResults(a int, b string) (int, string) {
	return a, b
}

// This Function Will Initialize The Named Return To Its Zero Value
func WithParametersAndNamedResults(a int, b string) (i int, s string) {
	i = a
	s = b
	return i, s
}
