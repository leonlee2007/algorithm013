package  week_01

import(
  "fmt"
  "testing"
)

//删除排序数组中的重复项
func TestRemoveDuplicates(t * testing.T) {
  fmt.Println("============Start TestRemoveDuplicates=============")
  nums := []int {0,0,1,1,1,2,2,3,3,4}
  newLen := removeDuplicates(nums)
  fmt.Println("input:", nums)
  fmt.Println("output:", newLen)
  fmt.Println("============End TestRemoveDuplicates=============")
}

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

