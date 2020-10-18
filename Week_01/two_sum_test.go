package  week_01

import(
  "fmt"
  "testing"
)

//两数之和
func TestTwoSum(t * testing.T) {
  fmt.Println("============Start TestTwoSum=============")
  nums := []int{0, 2, 3, 2, 0, 8, 9}
  target := 11
  idxList := twoSum2(nums, target)
  fmt.Println("input:", nums)
  fmt.Println("output:", idxList)
  fmt.Println("============End TestTwoSum=============")
}


//方法一：暴力枚举
//复杂度分析
//时间复杂度：O(N^2) 其中 NN 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
//空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
  for i := 0; i < len(nums) - 1; i ++ {
    for j := i + 1; j < len(nums); j++ { 
      if nums[i] + nums[j] == target {
        return []int{i, j}
      }
    }
  }
  return nil
}

//方法二：哈希表
//复杂度分析
//时间复杂度：O(N)，其中 N 是数组中的元素数量。对于每一个元素 x，我们可以 O(1) 地寻找 target - x。
//空间复杂度：O(N)，其中 N 是数组中的元素数量。主要为哈希表的开销。
func twoSum2(nums []int, target int) []int {
  hashTable := map[int]int{}
  for i, x := range nums {
    if p, ok := hashTable[target - x]; ok {
      return []int{p, i}
    }
    hashTable[x] = i
  }
  return nil
}

