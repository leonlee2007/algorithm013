package week_01

import(
  "fmt"
  "testing"
)
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func TestMoveZeroes(t * testing.T) {
  fmt.Println("============Start TestMoveZeroes=============")
  nums := []int {0,1,0,3,12}
  fmt.Println("input:", nums)
  moveZeroes1(nums)
  // moveZeroes2(nums)
  fmt.Println("output:", nums)
  // if maxArea != 49 { panic("result_is_error") }
  fmt.Println("============End TestMoveZeroes=============")
}

// 方法一
// 时间复杂度:O(n)
// 空间复杂度:O(1)
func moveZeroes1(nums []int)  {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[j] = nums[i]
            j ++
        }
    }
    for ;j < len(nums); j ++ {
        nums[j] = 0
    }
}

// 方法二 最优解
// 时间复杂度:O(n)
// 空间复杂度:O(1)
func moveZeroes2(nums []int)  {
    zeroPos := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[zeroPos], nums[i] = nums[i], nums[zeroPos]
            zeroPos ++
        }
    }
}
