// 221. 最大正方形
// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

// 示例:

// 输入:
// {
//     { '1','1','1','1','0'},
//     { '1','1','0','1','0'},
//     {'1','1','0','0','0'},
//     { '0','0','0','0','0'},
//   }

// 输出: 4
package week_06

import(
  "fmt"
  "testing"
)

func TestMaximalSquare(t * testing.T) {
  fmt.Println("============Start TestMaximalSquare=============")
  matrix := [][] byte {
    { '1','1','1','1','0'},
    { '1','1','0','1','0'},
    {'1','1','0','0','0'},
    { '0','0','0','0','0'},
  }
  ret := maximalSquare(matrix)
  fmt.Println("input:", matrix)
  fmt.Println("output:", ret)
  fmt.Println("============End TestMaximalSquare=============")
}

// 方法一：动态规划
func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0{ return 0 }
    m, n, ret := len(matrix), len(matrix[0]), 0
    dp := make([][]int, m+1)
    for i:=0; i < m + 1; i++ {dp[i] = make([]int, n + 1) }
    for i:=1;i<m+1;i++ {
        for j:=1;j < n+1; j++ {
            if matrix[i - 1][j - 1] == '1'{
                dp[i][j] = min(min(dp[i - 1][j],dp[i][j - 1]),dp[i - 1][j - 1]) + 1
                ret = max(ret, dp[i][j])
            }
        }
    }
    return ret * ret
}

func min(a, b int)int{
  if a < b {return a}
  return b
}

func max(a, b int) int{
    if a > b { return a }
    return b
}

