# Merge Sorted Array

You are given two sorted integer arrays nums1 and nums2, merged into a single sorted array.

- nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are initially 0 and should be ignored.
- nums2 has a length of n.

Merge nums2 into nums1 as one sorted array. The final sorted array should be stored in nums1. Do this in-place.

# Example Inputs and Expected Outputs

1) Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with nums1 as the target.

2) Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1] with nums1 as the target.


# Constraints

- 0 <= m, n <= 200

- 1 <= m + n <= 200

- -10^9 <= nums1[i], nums2[i] <= 10^9
