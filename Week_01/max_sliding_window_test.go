package week_01

import(
  "fmt"
  "testing"
)

//输入：{1,3,-1,-3,5,3,6,7}
//输出: {3,3,5,5,6,7}
func TestSlidingWindow(t * testing.T) {
  fmt.Println("============Start TestSlidingWindow=============")
  nums := []int {1,3,-1,-3,5,3,6,7}
  expect := []int {3,3,5,5,6,7}
  // maxNums := maxSlidingWindow2(nums, 3)
  // maxNums := maxSlidingWindow3(nums, 3)
  maxNums := maxSlidingWindow4(nums, 3)
  fmt.Println("input:", nums, 3)
  fmt.Println("output:", maxNums)
  fmt.Println("expect:", expect)
  fmt.Println("============End TestSlidingWindow=============")
}


// 方法一：暴力解法
//时间复杂度: O(kn)
//空间复杂度: O(n)
func maxSlidingWindow1(nums []int, k int) []int {
    lenNums := len(nums)
    if lenNums < k || k <= 0 {return nil}
    ret := make([]int, lenNums - k + 1)
    for i := 0; i < lenNums - k + 1; i++ {
        ret[i] = nums[i]
        for j := i + 1; j < i + k; j++ {
            ret[i] = max(nums[j], ret[i])
        }
    }
    return ret
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

// 方法二：暴力解法
//时间复杂度: O(kn)
//空间复杂度: O(n)
func maxSlidingWindow2(nums []int, k int) []int {
    if len(nums) == 0 || k <= 0 || k > len(nums) {return nil}
    var maxNums []int
    for i:= 0; i <= len(nums) - k; i++ {
        maxNums = append(maxNums, maxValue(nums[i:i+k]))
    }
    return maxNums
}

func maxValue(nums []int) int {
    left := 0
    right := len(nums)-1
    for left < right{
        if nums[left] > nums[right]{
            right--
        }else{
            left++
        }
    }
    return nums[left]
  }


// 方法3： 比较下标
//时间复杂度: O(n)
//空间复杂度: O(n)
func maxSlidingWindow3(nums []int, k int) []int {
    if len(nums) == 0 || k <= 0 || k > len(nums){
        return nil
    }
    var maxNums []int
    max := -1
    for i:=0; i <= len(nums) - k; i++ {
        left := i
        right := i + k - 1
        if max == -1 || max == left - 1 {
            max = getMax(nums, left, right)
        }else {
            if nums[right] > nums[max] {
                max = right
            }
        }
        maxNums = append(maxNums, nums[max])
    }
    return maxNums
}

func getMax(nums []int,left, right int) int {
    for left < right {
        if nums[left] > nums[right]{
            right--
        }else {
            left++
        }
    }
    return left
}

// 方法3： 比较值本身
//时间复杂度: O(n)
//空间复杂度: O(n)
func maxSlidingWindow4(nums []int, k int) []int {
  if len(nums) == 0 || k <= 0 || k > len(nums){
    return nil
  }
  var maxNums []int
  maxNow := maxValue1(nums[0:k])
  maxNums = append(maxNums,maxNow)
  for i := 1; i <= len(nums) - k; i++ {
    if nums[i - 1] == maxNow {
      maxNow = maxValue1(nums[i:i + k])
      maxNums = append(maxNums, maxNow)
    }else{
      if nums[i+ k - 1] <= maxNow{
        maxNums = append(maxNums, maxNow)
      }else {
        maxNow = nums[i + k - 1]
        maxNums = append(maxNums, maxNow)
      }
    }
  }
  return maxNums
}

func maxValue1(nums []int) int {
    left:=0
    right:=len(nums)-1
    for left<right{
        if nums[left]>nums[right]{
            right--
        }else{
            left++
        }
    }
    return nums[left]
}

