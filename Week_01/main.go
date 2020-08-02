package main

import("fmt")

func main() {
  //删除排序数组中的重复项
  fmt.Println("============RemoveDuplicates=============")
  nums1 := []int {0,0,1,1,1,2,2,3,3,4}
  newLen := removeDuplicates(nums1)
  fmt.Println("input:", nums1)
  fmt.Println("output:", newLen)
  fmt.Println("============TowSum=============")
  //两数之和
  nums2 := []int{0, 2, 3, 2, 0, 8, 9}
  target := 11
  idxList := twoSum(nums2, target)
  fmt.Println("input:", nums2)
  fmt.Println("output:", idxList)
}

//删除排序数组中的重复项
func removeDuplicates(nums []int) int {
  numsSize := len(nums)
  if numsSize < 2  {
    return numsSize
  }
  max := nums[numsSize - 1]
  j := 0
  for i := 0; i < numsSize; i++ {
    if nums[i] > nums[j] {
      j ++
      nums[j] = nums[i]
      if (nums[j] == max) { return j + 1 }
    }
  }
  return j + 1
}

//两数之和
func twoSum(nums []int, target int) []int {
  m := make(map[int]int)
  for index, num := range nums {
    key, ok := m[target - num]
    if !ok {
      m[num] = index
    }else {
      return []int{key, index}
    }
  }
  return nil
}
