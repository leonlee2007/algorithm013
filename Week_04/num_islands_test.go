
// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

// 此外，你可以假设该网格的四条边均被水包围。



// 示例 1:

// 输入:
// [
// ['1','1','1','1','0'],
// ['1','1','0','1','0'],
// ['1','1','0','0','0'],
// ['0','0','0','0','0']
// ]
// 输出: 1
// 示例 2:

// 输入:
// [
// ['1','1','0','0','0'],
// ['1','1','0','0','0'],
// ['0','0','1','0','0'],
// ['0','0','0','1','1']
// ]
// 输出: 3
// 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

package week_04

import(
  "fmt"
  "testing"
)

func TestNumIslands(t * testing.T) {
  fmt.Println("============Start TestNumIslands=============")
  grid := [][] byte {
    { '1','1','1','1','0'},
    { '1','1','0','1','0'},
    {'1','1','0','0','0'},
    { '0','0','0','0','0'},
  }
  ret := numIslands1(grid)
  fmt.Println("input:", grid)
  fmt.Println("output:", ret)
  m := make(map[int]int)
  x:= m[111]
  fmt.Println("======:", x, 20 / 3)
  fmt.Println("============End TestNumIslands=============")
}

// 方法一：
func numIslands1(grid [][]byte) int {
    m := len(grid)
    if m == 0 {return 0}
    n := len(grid[0])
    if n == 0 {return 0}
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                count ++
                handle_grid(i, j, m, n, grid)
            }
        }
    }
    return count
}

func handle_grid(i int, j int, m int, n int, grid [][]byte) {
    grid[i][j] = '0'
    find_grid(i, j - 1, m, n, grid)
    find_grid(i , j + 1, m, n, grid)
    find_grid(i + 1, j, m, n, grid)
    find_grid(i - 1, j, m, n, grid)
}

func find_grid(i int, j int, m int, n int, grid [][]byte) {
    if i >= 0 && i < m && j >=0 && j < n && grid[i][j] == '1' {
        handle_grid(i, j, m, n, grid)
    }
}

