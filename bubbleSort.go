package main

import 
	(
	"fmt"
	)

func main() {
	var x int
	var arr []int
	fmt.Print("\n\t-----Enter your numbers 10 times:----- ")
	for i := 0; i < 10; i++ {
		
		fmt.Print("\n\tEnter a no. : ")
		fmt.Scanln(&x)
		arr = append(arr, x)
	}
	fmt.Print( "\n\n\t  -----That is your Shorting Values-----") 
	sort := bubbleSort(arr)
	fmt.Print("\n\n\t", sort)
	fmt.Print("\n\n\t ----- Program closed **THANKU** and **HAVE A NICE DAY** :P -----")
}

func bubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
	
}