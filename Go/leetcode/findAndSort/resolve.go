package findAndSort

func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i := 0; i < n; i++ {
        nums1[m+i] = nums2[i]
        for j := m - 1 + i; j >= 0; j-- {
            if nums1[j + 1] < nums1[j] {
                nums1[j + 1], nums1[j] = nums1[j], nums1[j+1]
            }
        }
    }
}