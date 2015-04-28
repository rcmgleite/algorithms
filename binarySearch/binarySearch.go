package binarySearch

// PS: This package requires that the array is SORTED BEFORE CALLING THE SEARCH FUNCTIONS

// Recursive - recursive binary search
// this version returns -1 if value not present
// and the position of the value in the array if present
func Recursive(sortedArray []int, valueToSearch int) int {
	return doRecursive(sortedArray, valueToSearch, len(sortedArray)-1, 0)
}

func doRecursive(sortedArray []int, valueToSearch int, hi int, lo int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	if valueToSearch < sortedArray[mid] {
		return doRecursive(sortedArray, valueToSearch, mid-1, lo)
	} else if valueToSearch > sortedArray[mid] {
		return doRecursive(sortedArray, valueToSearch, hi, mid+1)
	} else {
		return mid
	}
}

// Iterative - iterative binary search
// this version return -1 if value not present
// and the position of the value in the array if present
func Iterative(sortedArray []int, valueToSearch int) int {
	lo := 0
	hi := len(sortedArray) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if valueToSearch < sortedArray[mid] {
			hi = mid - 1
		} else if valueToSearch > sortedArray[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
