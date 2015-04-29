package customSort

import (
	"fmt"
	"math/rand"
	"time"
)

// BySelection ... Simplest one - O(n²/2)
func BySelection(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				Swap(array, j, i)
			}
		}
	}
}

// ByInsertion ... executes sort by insertion - worst case - n²
// this algorithm depends on the initial state of the array
// if its partially sorted, this algorithm takes advantage of that
// and executes less operations than the selection one
// on the other side, if the array is sorted in the oposite order,
// it will take more time than the selection algorithm
func ByInsertion(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				Swap(array, j, j-1)
			}
		}
	}
}

// ByShell ...with 3x+1 increment sequence.
// Similar to insertion, but uses 'h'
func ByShell(array []int) {
	h := 1
	for h < len(array)/3 {
		h = 3*h + 1 // 3x + 1 increment sequence
	}

	for h >= 1 {
		//Insertion sort is the algorithm of the inside of this loop
		for i := h; i < len(array); i++ {
			j := i
			for j >= h {
				if array[j] < array[j-h] {
					Swap(array, j, j-h)
				}
				j -= h
			}
		}
		// end of the sort by insretion
		h /= 3 // update h and repeat
	}
}

// MergeSort ... O(n*log(n)) - will scale!
// 	1) Divide array into to halves
//	2) Recursively sort each half
// 	3) Merge the two halves
func MergeSort(array []int) {
	aux := make([]int, len(array))
	recMergeSort(array, aux, 0, len(array)-1)
}

// recursive auxiliar function for MergeSort
func recMergeSort(array, aux []int, lo, hi int) {
	// one improvment could be to execute insertion sort if the length
	// of the current array is small...

	if hi <= lo { // recursion stop
		return
	}

	mid := lo + (hi-lo)/2
	recMergeSort(array, aux, lo, mid)   // sort left
	recMergeSort(array, aux, mid+1, hi) // sort rigth

	// Minor improvment.. if the last element of the first sub-array
	// is less then the first element of the second sub-array
	// We don't need the merge call
	if array[mid] < array[mid+1] {
		return
	}

	Merge(array, aux, lo, mid, hi) // Finally, merge the sub-arrays
}

// Merge two halves of the same array that must be sorted when calling this function
// precondition:
// 	1) array[lo..mid]   sorted
//  2) array[mid+1..hi] sorted
func Merge(array, aux []int, lo int, mid int, hi int) {
	for k := lo; k <= hi; k++ { //Copy values to aux
		aux[k] = array[k]
	}

	// i = index of the leftmost item of the first half
	// j = index of the leftmost item of the second half
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ { // merge
		if i > mid { // if first half already empty, use second half
			array[k] = aux[j]
			j++
		} else if j > hi { // if second half already empty, use first half
			array[k] = aux[i]
			i++
		} else if aux[j] < aux[i] { // if a[j] < a[k] we use a[k]
			array[k] = aux[j]
			j++
		} else {
			array[k] = aux[i]
			i++
		}
	}
}

// Shuffle uniformly shuffle an array
func Shuffle(array []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(array); i++ {
		r := rand.Intn(i + 1)
		Swap(array, i, r)
	}
}

// Swap two elements of an array
func Swap(array []int, i int, j int) {
	a := array[i]
	array[i] = array[j]
	array[j] = a
}

// IsSorted verify if array is indeed sorted or not
func IsSorted(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			return false
		}
	}

	return true
}

// QuickSort ... recursive implementation - no need for extra space(better than MergeSort)
// Quick-sort and Merge-sort are both N*log(N) but quicksort is faster
// PS: Quick-sort is NOT STABLE
func QuickSort(array []int) {
	Shuffle(array) // VERY IMPORTANT TO SHUFFLE THE ARRAY BEFORE BEGINNING
	fmt.Println("Shuffled array = ", array)
	recQuickSort(array, 0, len(array)-1)
}

func recQuickSort(array []int, lo, hi int) {
	if hi <= lo { //recursive stop condition
		return
	}

	j := Partition(array, lo, hi)
	recQuickSort(array, lo, j-1) // Sort left part of array
	recQuickSort(array, j+1, hi) // Sort rigth part of array
}

// Partition step for Quick-sort
// return the index of the partition center
// the result of a partition is:
//	The array will have values < partition center on the left
//		and values > partition center on the right.
// ps: the two halves divided by the partition center may NOT BE SORTED
func Partition(array []int, lo, hi int) int {
	i := lo + 1
	j := hi

	for {
		// search for element < partition center (partition center = array[0])
		for array[i] < array[lo] {
			i++
			if i == hi { // just to be sure not get outside boundary
				break
			}
		}

		// search for element > partition center
		for array[j] > array[lo] {
			j--
			if j == lo { // just to be sure not get outside boundary
				break
			}
		}

		// When i cross j, stop while(true)
		if i >= j {
			break
		}

		Swap(array, i, j) // execute the swap
	}

	Swap(array, lo, j) // Swap the partition center with the rigthmost element < partition center
	return j           //return the index of the partition center
}
