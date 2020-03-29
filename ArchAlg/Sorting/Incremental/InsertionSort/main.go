package main

import (
	"fmt"
	)

// Algorithm Type		:		Sorting
// Algorithm			: 		Insertion Sort
// Formal Definition	:
//								Input	: 	A sequence of n numbers {a1, a2, ..... an}
//								Output	:	A permutation (reordering) {a1, a2, ...... an} of the input sequence each that
//								            a1 <= a2 <= ...... <= an

// Legend:
// 	c represents a constant.
// 	O represent BigO notation
// 	n represents the number of items of input.

// Below is the pseudocode  of the insertion sort. The insertion sort will run in O(c*n^2) time.
// 		for j = 2 to A.length
//  		key = A[j]
//  		i = j - 1
//  		while i > 0 and A[i] > key
//    			A[i + 1] = A[i]
//    			i = i - 1
//  		A[i + 1] = key

// Below is a textual representation of the process of the insertion sort.

// Original
//       1 Based Index
// 	[1   2  3  4  5   6   7]
// 	{6, 34, 5, 1, 3, 100, 44}

// First Pass
//       1 Based Index
// 	[1   2  3  4  5   6   7]
// 	{6, 34, 5, 1, 3, 100, 44}
//                                              j = 2
// 												A.length = 7,
// 												A[j] = A[2] = 34,
//												key = 34
// 												i = j - 1 = 2 - 1 = 1,
// 												A[i] = A[0] = 6
//												A[i + 1] = A[1 + 1] = = A[2] 34

// 		for j = 2 to A.length				|	 for j = 2 to 7					|		           j			       			A.length
//  		key = A[j]						| 		key = 34					|		           v				     		   v
//  		i = j - 1						| 		i = 2 - 1 = 0				|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 1 > 0 and 6 > 34		|	   { 6,    34,   5,    1,    3,    100,    44 }
//    			A[i + 1] = A[i]				| 			DOESN'T ENTER HERE		|	  	 ^     ^
//    			i = i - 1					|
//  		A[i + 1] = key					| 		A[1 + 1] = 34				|	    A[i]  A[i+1]

// Second Pass
//       1 Based Index
// 	[1     2    3    4    5     6     7 ]  -UPDATED
// 	{6,   34,   5,   1,   3,   100,   44}
//                                              j = 3
// 												A.length = 7,
// 												A[j] = A[3] = 5,
//												key = A[j] = 5
// 												i = j - 1 = 3 - 1 = 2,
// 												A[i] = A[2] = 34
//												A[i + 1] = A[2 + 1] = A[3] = 5

// 		for j = 2 to A.length				|	 for j = 3 to 7					|		       i     j			       	    A.length
//  		key = A[j]						| 		key = 5 					|		       v     v				     	   v
//  		i = j - 1						| 		i = 3 - 1				 	|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 2 > 0 and 34 > 5		|	   { 6,    5,   34,    1,    3,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[3] = A[2]     		|	  	       ^  -> ^
//    			i = i - 1					|           i = 2 - 1				|			  A[i]  A[i+1]
//  		A[i + 1] = key					| 		A[2 + 1] = A[3] = 5			|
// --------------------------------------------------------------------------------------------------------------------------------------
//											   			WHILE LOOP 2
// 											|	 	i = 1   					|		 i           j			       	     A.length
//  										| 									|		 v           v				     	   v
//  										| 								 	|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 1 > 0 and 6 > 5		|	   { 6,    6,   34,   1,    3,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[2] = A[1]     		|	  	 ^ ->  ^
//    			i = i - 1					|           i = 1 - 1				|		A[i] A[i+1]
//  										| 									|
//											|		i = 0 STOP WHILE LOOP       |
// --------------------------------------------------------------------------------------------------------------------------------------
// 			A[i + 1] = key					|       A[0 + 1] = 5                |	   { 5,    6,   34,   1,    3,    100,    44 }

// Third Pass
//       1 Based Index
// 	[1    2    3     4    5     6     7 ]
// 	{5,   6,   34,   1,   3,   100,   44} -UPDATED
//                                              j = 4
// 												A.length = 7,
// 												A[j] = A[4] = 1,
//												key = A[j] = 1
// 												i = j - 1 = 4 - 1 = 3,
// 												A[i] = A[3] = 34
//												A[i + 1] = A[3 + 1] = A[4] = 1

// 		for j = 2 to A.length				|	 for j = 4 to 7					|		       i <-  i     j			    A.length
//  		key = A[j]						| 		key = 1 					|		             v     v				   v
//  		i = j - 1						| 		i = 4 - 1 = 3			 	|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 3 > 0 and 34 > 1		|	   { 5,    6,   34,    34,    3,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[4] = A[3]     		|	  	            ^      ^
//    			i = i - 1					|           i = 3 - 1
//  		A[i + 1] = key					| 		A[2 + 1] = 5				|	      	       A[i]   key
// --------------------------------------------------------------------------------------------------------------------------------------

//											   			WHILE LOOP PASS 2
// 											|	 	i = 2   					|		 i <-  i           j			    A.length
//  										| 		key = 1						|		       v           v				   v
//  										| 								 	|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 2 > 0 and 6 > 1		|	   { 5,    6,    6,   34,    3,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[3] = A[2]     		|	  	       ^    ^
//    			i = i - 1					|           i = 2 - 1				|		     A[i] A[i+1]
//  										| 									|
// --------------------------------------------------------------------------------------------------------------------------------------

//											   			WHILE LOOP PASS 3
// 											|	 	i = 1   					|		 i <-  i           j	   	        A.length
//  										| 		key = 1						|		       v           v				   v
//  										| 								 	|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 1 > 0 and 5 > 1		|	   { 5,    5,    6,   34,    3,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[2] = A[1]     		|	  	       ^    ^
//    			i = i - 1					|           i = 1 - 1				|		     A[i] A[i+1]
//  										| 			STOP WHILE LOOP			|
// --------------------------------------------------------------------------------------------------------------------------------------
// 			A[i + 1] = key					|       A[0 + 1] = 1                |	   { 1,    5,   6,   34,    3,    100,    44 }

// Fourth Pass
//       1 Based Index
// 	[1    2    3     4    5     6     7 ]
// 	{1,   5,   6,   34,   3,   100,   44} -UPDATED
//                                              j = 5
// 												A.length = 7,
// 												A[j] = A[5] = 3,
//												key = A[j] = A[5] = 3
// 												i = j - 1 = 5 - 1 = 4,
// 												A[i] = A[4] = 34
//												A[i + 1] = A[4 + 1] = A[5] = 3

// 		for j = 2 to A.length				|	 for j = 5 to 7					|		             i <-  i      j			A.length
//  		key = A[j]						| 		key = 3 					|		                   v      v			   v
//  		i = j - 1						| 		i = 5 - 1 = 4			 	|	   [ 1     2     3     4      5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 4 > 0 and 34 > 3		|	   { 1,    5,    6,    34,    34,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[5] = A[4]     		|	  	                   ^      ^
//    			i = i - 1					|           i = 3 - 1
//  		A[i + 1] = key					| 		A[2 + 1] = 5				|	      	              A[i]   key
// --------------------------------------------------------------------------------------------------------------------------------------
//											   			WHILE LOOP PASS 2
// 											|	 	i = 4   					|		             i <-  i     j			A.length
//  										| 		key = 3						|		                   v     v			   v
//  										| 								 	|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 4 > 0 and 34 > 3		|	   { 1,    5,    6,   6,    34,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[4] = A[3]     		|	  	                  ^      ^
//    			i = i - 1					|           i = 4 - 1				|		                 A[i]  A[i+1]
//  										| 									|
// --------------------------------------------------------------------------------------------------------------------------------------
//											   			WHILE LOOP PASS 2
// 											|	 	i = 3   					|		             i <-  i     j			A.length
//  										| 		key = 3						|		                   v     v			   v
//  										| 								 	|	   [ 1     2     3     4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 3 > 0 and 6 > 3		|	   { 1,    5,    6,   6,    34,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[4] = A[3]     		|	  	                  ^      ^
//    			i = i - 1					|           i = 4 - 1				|		                 A[i]  A[i+1]
//  										| 									|
// --------------------------------------------------------------------------------------------------------------------------------------
//											   			WHILE LOOP PASS 3
// 											|	 	i = 2   					|		       i <-  i          j	       A.length
//  										| 		key = 3						|		             v          v	          v
//  										| 								 	|	   [ 1     2     3    4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 2 > 0 and 5 > 3		|	   { 1,    5,    5,   6,    34,    100,    44 }
//    			A[i + 1] = A[i]				| 			A[3] = A[2]     		|	  	       ^     ^
//    			i = i - 1					|           i = 2 - 1				|		      A[i] A[i+1]
//  										| 									|
// --------------------------------------------------------------------------------------------------------------------------------------
//											   			WHILE LOOP PASS 4
// 											|	 	i = 1   					|		 i       	            j 	     	A.length
//  										| 		key = 3						|		 v                      v			  v
//  										| 								 	|	   [ 1     2     3    4     5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 1 > 0 and 1 > 3		|	   { 1,    3,    5,   6,    34,    100,    44 }
//    			A[i + 1] = A[i]				| 			 DOESN'T ENTER STOP 	|	  	 ^     ^
//    			i = i - 1					|                   				|		A[i] A[i+1]
//  		A[i + 1] = key					| 		A[1 + 1] = A[2] = 5				|	      	              A[i]

// Fifth Pass
//       1 Based Index
// 	[1    2    3     4    5     6     7 ]
// 	{1,   3,   5,   6,   34,   100,   44} -UPDATED
//                                              j = 6
// 												A.length = 7,
// 												A[j] = A[6] = 100,
//												key = A[j] = A[5] = 100
// 												i = j - 1 = 6 - 1 = 5,
// 												A[i] = A[5] = 34
//												A[i + 1] = A[5 + 1] = A[6] = 100

// 		for j = 2 to A.length				|	 for j = 6 to 7					|		                    i <-  i      j   A.length
//  		key = A[j]						| 		key = 100 					|		                          v      v	    v
//  		i = j - 1						| 		i = 6 - 1 = 5			 	|	   [ 1     2     3     4      5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 5 > 0 and 34 > 100    |	   { 1,    3,    5,    6,    34,    100,    44 }
//    			A[i + 1] = A[i]				| 		  DOESN'T ENTER STOP 	    |	  	                          ^      ^
//    			i = i - 1					|
//  		A[i + 1] = key					| 		A[5 + 1] = A[6] = 100	    |	      	                     A[i]   key
// --------------------------------------------------------------------------------------------------------------------------------------

// Sixth Pass
//       1 Based Index
// 	[1    2    3     4    5     6     7 ]
// 	{1,   3,   5,   6,   34,   100,   44} -UPDATED
//                                              j = 7
// 												A.length = 7,
// 												A[j] = A[7] = 44,
//												key = A[j] = A[7] = 44
// 												i = j - 1 = 7 - 1 = 6,
// 												A[i] = A[6] = 100
//												A[i + 1] = A[6 + 1] = A[7] = 44

// 		for j = 2 to A.length				|	 for j = 7 to 7					|		                          i <-   i   j, A.length
//  		key = A[j]						| 		key = 44 					|		                                 v	    v
//  		i = j - 1						| 		i = 7 - 1 = 6			 	|	   [ 1     2     3     4      5      6      7 ]
//  		while i > 0 and A[i] > key		| 		while 6 > 0 and 100 > 44    |	   { 1,    3,    5,    6,    34,    100,    100 }
//    			A[i + 1] = A[i]				| 			A[6 + 1] = A[7] = A[6] 	|	  	                                 ^      ^
//    			i = i - 1					|			i = 6 - 1				|	      	                            A[i]   A[i + 1]
// --------------------------------------------------------------------------------------------------------------------------------------
//											   			WHILE LOOP PASS 2
// 											|	 	i = 5   					|		                   i               j, A.length
//  										| 		key = 44					|		                   v                     v
//  										| 								 	|	   [ 1     2     3     4     5       6       7 ]
//  		while i > 0 and A[i] > key		| 		while 6 > 0 and 34 > 44		|	   { 1,    3,    5,    6,    34,    44,    100 }
//    			A[i + 1] = A[i]				| 		  DOESN'T ENTER STOP 	    |	  	                          ^      ^
//    			i = i - 1					|           			         	|		                         A[i]  A[i+1]
//  		A[i + 1] = key					| 		A[5 + 1] = A[6] = 44	    |	      	                     A[i]   key
// --------------------------------------------------------------------------------------------------------------------------------------
// Seventh Pass
//       1 Based Index
// 	[1    2    3     4    5     6     7 ]
// 	{1,   3,   5,   6,   34,   44,   100} -UPDATED
//                                              j = 8
// 												A.length = 7,
// 												A[j] = A[8] = %ERROR%,
//												key = A[j] = A[8] = %ERROR%
// 												i = j - 1 = 8 - 1 = 7,
// 												A[i] = A[7] = 44
//												A[i + 1] = A[7 + 1] = A[8] = %ERROR%

// 		for j = 2 to A.length				|	 for j = 78 to 7				|		                          i <-   i   j, A.length
//  		key = A[j]						| 		DOESN'T ENTER STOP			|		                                 v	    v
//  		i = j - 1						| 				 					|	   [ 1     2     3     4      5      6      7 ]
//  		while i > 0 and A[i] > key		| 		    						|	   { 1,    3,    5,    6,    34,    100,    100 }
//    			A[i + 1] = A[i]				| 			 						|	  	                                 ^      ^
//    			i = i - 1					|									|	      	                            A[i]   A[i + 1]
//  		A[i + 1] = key					| 		A[5 + 1] = A[6] = 44	    |	      	                     A[i]   key
// --------------------------------------------------------------------------------------------------------------------------------------
// SORTED LIST {1, 3, 5, 6, 34, 44, 100}
//

func main() {
	xi := []int{1, 4, 5, 778, 221, 389, 420, 69, 331}
	InsertionSort(xi)
	fmt.Println(xi)
}

func InsertionSort(xi []int) {
	sz := len(xi)
	fmt.Println(xi)
	for j := 1; j < sz; j++ {
		i := j  - 1
		mk := xi[j]
		for i > 0 && xi[i] > mk {
			xi[i + 1] = xi[i]
			i--
		}
		xi[i + 1] = mk
	}
}

// Quotes: You have to learn to lose, to learn to win.
//		   I'm so scared to die, I have too much to prove.