// 剑指 Offer 49. 丑数
// 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
// 示例:
// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
// 说明:  
// 1 是丑数。
// n 不超过1690。

package week_02

import(
  "fmt"
  "testing"
)


func TestNthUglyNumber(t * testing.T) {
  fmt.Println("============Start TestNthUglyNumber=============")
  input := 10
  output := nthUglyNumber(input)
  fmt.Println("input:", input)
  fmt.Println("output:", output)
  fmt.Println("============End TestNthUglyNumber=============")
}

// 解法一 动态规划思想
func nthUglyNumber(n int) int {
  uglyNums := make([]int, n, n)
  uglyNums[0] = 1
  p2, p3, p5 := 0, 0, 0
  for i := 1; i < n; i++ {
    // 更新数组
    uglyNum2 := uglyNums[p2] * 2
    uglyNum3 := uglyNums[p3] * 3
    uglyNum5 := uglyNums[p5] * 5
    minUglyNum := getMinNum(uglyNum2, uglyNum3, uglyNum5)
    uglyNums[i] = minUglyNum
    // 移动指针
    if minUglyNum == uglyNum2 {p2 ++}
    if minUglyNum == uglyNum3 {p3 ++}
    if minUglyNum == uglyNum5 {p5 ++}
  }
  return uglyNums[n - 1];
}

func getMinNum(numA int, numB int, numC int) int {
  minNum := numA
  if minNum > numB {minNum = numB}
  if minNum > numC {minNum = numC}
  return minNum
}
