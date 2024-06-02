package mergesort

import "sync"

func sequentialMergesort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	sequentialMergesort(s[:middle]) // First half
	sequentialMergesort(s[middle:]) // Second half
	merge(s, middle)                // Merges the two halves
}

func parallelMergesortV1(s []int) {
	if len(s) <= 1 {
		return
	}
	middle := len(s) / 2
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		parallelMergesortV1(s[:middle])
	}()
	go func() {
		defer wg.Done()
		parallelMergesortV1(s[middle:])
	}()
	wg.Wait()
	merge(s, middle)
}

const max = 2048

func parallelMergesortV2(s []int) {
	if len(s) <= 1 {
		return
	}
	if len(s) <= max {
		sequentialMergesort(s)
	} else {
		middle := len(s) / 2
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			parallelMergesortV2(s[:middle])
		}()
		go func() {
			defer wg.Done()
			parallelMergesortV2(s[middle:])
		}()
		wg.Wait()
		merge(s, middle)
	}
}

func merge(s []int, middle int) {
	left := append([]int(nil), s[:middle]...)
	right := append([]int(nil), s[middle:]...)

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			s[k] = left[i]
			i++
		} else {
			s[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		s[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		s[k] = right[j]
		j++
		k++
	}
}
