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
  moveZeroes(nums)
  fmt.Println("output:", nums)
  // if maxArea != 49 { panic("result_is_error") }
  fmt.Println("============End TestMoveZeroes=============")
}

func moveZeroes(nums []int)  {
    zeroPos := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[zeroPos], nums[i] = nums[i], nums[zeroPos]
            zeroPos ++
        }
    }
}
