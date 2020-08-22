// 33. 搜索旋转排序数组
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

// 你可以假设数组中不存在重复的元素。

// 你的算法时间复杂度必须是 O(log n) 级别。

// 示例 1:

// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4
// 示例 2:

// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1

package week_04

import(
  "fmt"
  "testing"
)

func TestSearch(t * testing.T) {
  fmt.Println("============Start TestBinarySearch=============")
  array := [] int {4, 5, 6, 7, 0, 1, 2}
  target := 0
  ret := search1(array, target)
  fmt.Println("input:", array, target)
  fmt.Println("output:", ret)
  fmt.Println("============End TestBinarySearch=============")
}

// 方法一：
func search1(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := left + (right - left) / 2
        switch {
            case nums[mid] == target:
                return mid
            case nums[0] <= nums[mid] && (nums[mid] < target || nums[0] > target) :
                left = mid + 1
            case target < nums[0] && target > nums[mid] :
                left = mid + 1
            default:
                right = mid - 1
        }
    }
    return -1
 }

