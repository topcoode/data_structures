package array

import "fmt"

func Binary_main() {
	// binary search /Iterative  Binary Search Algorithm
	data := Binary_Searching_Iterative()
	fmt.Println("data:", data)
	if data == -1 {
		fmt.Println("data in not found in the array")
	} else {
		fmt.Println("data is present the index is :", data)
	}
	// binary search /Recursive  Binary Search Algorithm
	Binary_Searching_Iterative()
}
func Binary_Searching_Iterative() int {
	myArray := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	x := 12
	low := 0
	high := len(myArray) - 1

	for low <= high {
		mid := (low + high) / 2

		if myArray[mid] == x {
			return mid
		} else if myArray[mid] < x {
			low = mid + 1
			fmt.Println("low--->", low)
		} else {
			high = mid - 1
			fmt.Println("high--->", high)
		}
	}

	return -1
}

func binarySearch_recursive(arr []int, x int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if arr[mid] == x {
		return mid
	} else if arr[mid] < x {
		fmt.Println("coming--->")
		return binarySearch_recursive(arr, x, mid+1, high)
	} else {
		return binarySearch_recursive(arr, x, low, mid-1)
	}
}

func binarySearch_main() {
	myArray := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	x := 38
	index := binarySearch_recursive(myArray, x, 0, len(myArray)-1)
	if index == -1 {
		fmt.Println("Element not found")
	} else {
		fmt.Println("Element found at index", index)
	}
}
