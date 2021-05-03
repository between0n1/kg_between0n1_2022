// created by Sunhyeok Jung on May 3
// Kargo-2021-Summer-Internship-assessment

package main

import (
	"fmt"
	"os"
)

func main() {
	var argument_size int
	argument_size = len(os.Args) - 1 // arguments size
	for i:= 0; i < argument_size; i++{
		solution(os.Args[i+1])
		if i < argument_size -1{
			fmt.Printf(",")
		} else{
			fmt.Printf("\n")
		}
	}
}

func solution(input_string string){
	var english = [10]string{"Zero","One","Two","Three","Four","Five","Six","Seven","Eight","Nine"}
	var length int
	length = len(input_string)
	for i := 0; i < length; i++{
		fmt.Printf(english[input_string[i] - 48])
	}
}
