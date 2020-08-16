
// 46. 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
// 示例:

// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

package week_03

import(
  "fmt"
  "testing"
)

func TestPermute(t * testing.T) {
  fmt.Println("============Start TestPermute=============")
  nums := []int{1,2,3}
  ret := permute1(nums)
  fmt.Println("input:", nums)
  fmt.Println("output:", ret)
  fmt.Println("============End TestPermute=============")
}

// 方法一：
// 解题思路 回溯
// 套用回溯公式:
// result = []
// func backtrack(路径，选择列表) {
// 	if 满足结束条件 {
// 		result.add(路径)
// 	}
// 	return

// 	for 选择 in 选择列表 {
// 		做选择
// 		backtrack(路径，选择列表)
// 		撤销选择
// 	}
// }

// 最终结果
var result [][]int

func permute1(nums []int) [][]int {
  var pathNums []int
  var used = make([]bool, len(nums))
  result = [][]int{}
  backtrack(nums, pathNums, used)
  return result
}


func backtrack(nums, pathNums []int, used[]bool) {
  //terminator 结束条件
  if len(nums) == len(pathNums) {
    tmp := make([]int, len(nums))
    copy(tmp, pathNums)
    result = append(result, tmp)
    return
  }

  for i:=0; i<len(nums); i++ {
    if !used[i] {
      used[i] = true
      //process 做选择, 做操作
      pathNums = append(pathNums, nums[i])
      //drilldown 下探到下一层
      backtrack(nums,pathNums,used)
      //reverse status 撤销的选择，恢复操作
      pathNums = pathNums[:len(pathNums) -1]
      used[i] = false
    }
  }
}

