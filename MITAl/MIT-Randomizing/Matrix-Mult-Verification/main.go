package main

// MIT Lecture notes for Randomizing Algorithms.

// A Randomized Algorithm is an algorithm that generates a random number. It
// will make decisions off of the value that it finds.

// Monte Carlo: Probably Correct Algorithms.
// Las Vegas: Probably Fast Algorithms.

// Multiplication takes the computer longer to compute compared to addition.

// Matrix Product Notes:
//	Equation: 			c = a * B
//	Simple Algorithm:	O(n^3) multiplications
//  Strassen:			Multiply 2 2x2 matrices using 7 multiplications.
//						O(n^(logv2 7) = O(n^2.81)