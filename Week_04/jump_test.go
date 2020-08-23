// 45. 跳跃游戏 II
// 给定一个非负整数数组，你最初位于数组的第一个位置。

// 数组中的每个元素代表你在该位置可以跳跃的最大长度。

// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。

// 示例:

// 输入: [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 说明:

// 假设你总是可以到达数组的最后一个位置。


package week_04

import(
  "fmt"
  "testing"
)

func TestJump(t * testing.T) {
  fmt.Println("============Start TestJump=============")
  nums := [] int{2,3,1, 1, 4}
  ret := jump2(nums)
  fmt.Println("input:", nums)
  fmt.Println("output:", ret)
  fmt.Println("============End TestJump=============")
} 
// 方法一：顺藤摸瓜
func jump1(nums []int) int {
    end, maxPosition, steps := 0, 0, 0
    for i :=0; i < len(nums) - 1; i++ {
        //找能跳的最远的
        maxPosition = max(maxPosition, nums[i] + i)
        //遇到边界，就更新边界，并且步数加一
        if i == end {
            end = maxPosition
            steps++
        }
    }
    return steps
}

func max(a int, b int) int {
    if a > b {return a}
    return b
}

//方法二: 顺瓜摸藤
func jump2(nums []int) int {
    position, steps := len(nums) - 1, 0 //要找的位置
    for position != 0 { //是否到了第 0 个位置
        for i := 0; i < position; i++ {
            if nums[i] >= position - i {
                position = i //更新要找的位置
                steps++
                break
            }
        }
    }
    return steps
}

