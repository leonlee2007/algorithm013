
//二分查找算法模板
package week_04

import(
  "fmt"
  "testing"
)

func TestBinarySearch(t * testing.T) {
  fmt.Println("============Start TestBinarySearch=============")
  array := [] int {1, 2, 3,4 ,5 , 10 ,30, 40, 60, 79, 100}
  target := 3
  ret := binarySearch1(array, target)
  fmt.Println("input:", array, target)
  fmt.Println("output:", ret)
  fmt.Println("============End TestBinarySearch=============")
}

// 方法一：
func binarySearch1 (array []int, target int) int {
  left, right := 0, len(array) -1
  for left <= right {
    mid := (right + left)/ 2;
    if array[mid] == target{
      return mid
    } else if array[mid] < target {
      left = mid + 1
    } else {
      right = mid -1
    }
  }
  return -1
}
