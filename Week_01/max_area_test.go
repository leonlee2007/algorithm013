package week_01

import(
  "fmt"
  "testing"
)

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
func TestMaxArea(t * testing.T) {
  fmt.Println("============Start TestMaxArea=============")
  nums := []int {1,8,6,2,5,4,8,3,7}
  // maxArea := maxArea1(nums)
  maxArea := maxArea2(nums)
  fmt.Println("input:", nums)
  fmt.Println("output:", maxArea)
  if maxArea != 49 { panic("result_is_error") }
  fmt.Println("============End TestMaxArea=============")
}


// 方法一：暴力解法
//时间复杂度: O(n2)
//空间复杂度: O(1)
func maxArea1(height []int) int {
    maxArea := 0
    for i := 0; i < len(height) - 1; i++ {
        for j := i + 1; j < len(height); j ++ {
            h := min(height[i], height[j])
            maxArea = max(maxArea, (j - i) * h)
        }
    }
    return maxArea
}

// 方法2：双指针法
//时间复杂度: O(n)
//空间复杂度: O(1)
func maxArea2(height []int) int {
  maxArea, l, r := 0, 0, len(height) - 1
  for l < r {
    minH := min(height[l], height[r])
    maxArea = max(maxArea, minH * (r - l))
    if height[l] > height[r] {
      r --
    }else {
      l ++
    }
  }
  return maxArea
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}

