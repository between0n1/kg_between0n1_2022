// created by Sunhyeok Jung on May 3
// Kargo-2021-Summer-Internship-assessment
// This solution convert integer into corresponding english word
// ex: 732 -> SevenThreeTwo
//     123 -> OneTwoThree

package main

import (
	"fmt"
	"os"
)

//driver code
func main() {
	var argument_size = len(os.Args) - 1 // arguments size
	for i:= 0; i < argument_size; i++{
		solution(os.Args[i+1])
		if i < argument_size -1{
			fmt.Printf(",")
		} else{
			fmt.Printf("\n")
		}
	}
}

// Required Solution
func solution(input_string string){
	var english = [10]string{"Zero","One","Two","Three","Four","Five","Six","Seven","Eight","Nine"}
	var length = len(input_string) // length of a string from input seperated by space
	for i := 0; i < length; i++{
		fmt.Printf(english[input_string[i] - 48])
	}
	// could we do this without using for loop?
}
