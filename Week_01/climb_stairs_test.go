package week_01

import(
  "fmt"
  "testing"
)
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func TestClimbStairs(t * testing.T) {
  fmt.Println("============Start TestClimbStairs=============")
  fmt.Println("input:", 10)
  // Ret := climbStairs3(10)
  Ret := climbStairs4(10)
  fmt.Println("output:", Ret)
  fmt.Println("============End TestClimbStairs=============")
}

// 方法一 暴力求解 不推荐
// 时间复杂度:O(2的n次方)
// 空间复杂度:O(n)
func climbStairs1(n int) int {
  if n == 1 {return 1}
  if n == 2 {return 2}
  return climbStairs1(n - 1) + climbStairs1(n - 2)
}

// 方法二 记忆化递归
// 思路 把求出的中间结构存储，避免重复运算
// 时间复杂度:O(n)
// 空间复杂度:O(n)
func climbStairs2(n int) int {
    mem := map[int]int{1:1, 2:2}
    return climbStairsMem(n, mem)
}

func climbStairsMem(n int, mem map[int]int) int {
    value, ok := mem[n]
    if ok {return value}
    mem[n] = climbStairsMem(n - 1, mem) + climbStairsMem(n - 2, mem)
    return mem[n]
}

// 方法三 动态规划
// 思路: 用数组记录每个状态 
// 时间复杂度:O(n)
// 空间复杂度:O(n)
func climbStairs3(n int) int {
  ret := make([]int, 0, n + 1)
  ret = append(ret, 0)
  ret = append(ret, 1)
  ret = append(ret, 2)
  for i := 3; i < n + 1; i++ {
    ret = append(ret, ret[i-2] + ret[i-1])
  }
  return ret[n]
}

// 方法四 动态规划 优化
// 思路: 计算是仅需要2个状态，这样就可以把空间复杂度优化到O(1)
// 时间复杂度:O(n)
// 空间复杂度:O(1)
func climbStairs4(n int) int {
  if n == 1 {return 1}
  first, second := 1, 2
  for i := 3; i < n + 1; i++ {
      first, second = second, first + second
  }
  return second
}

