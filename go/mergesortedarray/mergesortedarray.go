package mergesortedarrays

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	// p1 points to the last element of nums1 that needs to be considered.
	p1 := m - 1
	// p2 points to the last element of nums2.
	p2 := n - 1
	// p points to the end of nums1, where the merged element will be placed.
	p := m + n - 1

	// Iterate until either p1 or p2 becomes negative (meaning one of the arrays is exhausted).
	for p1 >= 0 && p2 >= 0 {
		// Compare the elements at p1 and p2.
		if nums1[p1] > nums2[p2] {
			// If nums1[p1] is larger, place it at the end of nums1.
			nums1[p] = nums1[p1]
			p1-- // Move p1 to the left.
		} else {
			// Otherwise, place nums2[p2] at the end of nums1.
			nums1[p] = nums2[p2]
			p2-- // Move p2 to the left.
		}
		p-- // Move p to the left (towards the beginning of nums1).
	}

	// If nums2 still has elements left (nums1 is exhausted), copy them into nums1.
	// This handles the case where all elements of nums1 are greater than elements of nums2
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}

	// No need to handle if nums1 still has elements (nums2 is exhausted) because the initial
	// m elements of nums1 are already in the correct place.
	// (delete me)
	// (also delete me)
}
