package  week_01

import(
  "fmt"
  "testing"
)

//两数之和
func TestTwoSum(t * testing.T) {
  fmt.Println("============Start TestTwoSum=============")
  nums := []int{0, 2, 3, 2, 0, 8, 9}
  target := 11
  idxList := twoSum(nums, target)
  fmt.Println("input:", nums)
  fmt.Println("output:", idxList)
  fmt.Println("============End TestTwoSum=============")
}


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
