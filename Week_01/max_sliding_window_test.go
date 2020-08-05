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
  // maxNums := maxSlidingWindow4(nums, 3)
  maxNums := maxSlidingWindow5(nums, 3)
  fmt.Println("input:", nums, 3)
  fmt.Println("output:", maxNums)
  fmt.Println("expect:", expect)
  fmt.Println("============End TestSlidingWindow=============")
  queue := []int{1,2,3,4}
  queue2 := queue[1:4]
  queue3 := append(queue2, 100)
  queue3[1] = 1000000
  queue3[2] = 2000000
  fmt.Println(queue2[1], queue3[1], queue[2])
  fmt.Println(queue2[1], queue3[2], queue[3])
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

// 方法5： 使用单调队列 最有解
//使用单调队列来求解，构建一个单调队列queue
//queue里面的元素是单调递减的，而且长度小于等于滑动窗口的长度
//时间复杂度: O(n)
//空间复杂度: O(n)
func maxSlidingWindow5(nums []int, k int) []int {
  if len(nums) == 0 { return nil}
  if k == 1 { return nums }
  queue := make([]int, 0, len(nums))
  for i := 0; i < k ; i++ {
    queue = add_value_to_queue(queue, nums[i])
  }
  res := make([]int, 0, len(nums) - k + 1)
  res = append(res, queue[0])
  for i, j := 1, k; j < len(nums); i, j = i + 1, j + 1 {
    if nums[i - 1] == queue[0] {queue = queue[1:] }
    queue = add_value_to_queue(queue, nums[j])
    res = append(res, queue[0])
  }
  return res
}

func add_value_to_queue(queue []int, value int) []int {
  for len(queue) > 0 && value > queue[len(queue) - 1] {
    queue = queue[:len(queue) - 1]
  }
  return  append(queue, value)
}

