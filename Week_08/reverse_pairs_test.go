// 493. 翻转对
// 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。

// 你需要返回给定数组中的重要翻转对的数量。

// 示例 1:

// 输入: [1,3,2,3,1]
// 输出: 2
// 示例 2:

// 输入: [2,4,3,5,1]
// 输出: 3
// 注意:

// 给定数组的长度不会超过50000。
// 输入数组中的所有数字都在32位整数的表示范围内。

package week_08

import(
  "fmt"
  "testing"
)

func TestReversPairsgc(t * testing.T) {
  fmt.Println("============Start TestReversPairsgc=============")
  m := [] int { 1, 3, 2, 3, 1}
  ret := reversePairs(m)
  fmt.Println("input:", m)
  fmt.Println("output:", ret)
  fmt.Println("============End TestReversPairsgc=============")
}

func reversePairs(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	tmp := make([]int, len(nums))
	return sortAndCount(nums, 0, len(nums)-1, tmp)
}
func merge(arr []int, low, mid, high int, tmp []int) {
	i, j, k := low, mid+1, 0
	for i <= mid && j <= high {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i, k = i + 1, k + 1
		} else {
			tmp[k] = arr[j]
			k, j = k + 1, j + 1
		}
	}
	for i <= mid {
		tmp[k] = arr[i]
		k, i = k + 1, i + 1
	}
	for j <= high {
		tmp[k] = arr[j]
		k, j = k + 1, j + 1
	}
	copy(arr[low : high + 1], tmp)
}

func count(arr []int, low, mid, high int) int {
	i, j := low, mid+1
	cnt := 0
	for i <= mid && j <= high {
		if arr[i] <= 2 * arr[j] {
			i++
		} else {
			cnt += (mid - i + 1)
			j++
		}
	}
	return cnt
}

func sortAndCount(arr []int, low, high int, tmp []int) int {
	cnt := 0
	if low < high {
		mid := low + (high-low)/2
		cnt += sortAndCount(arr, low, mid, tmp)
		cnt += sortAndCount(arr, mid+1, high, tmp)
		cnt += count(arr, low, mid, high)
        merge(arr, low, mid, high, tmp)
	}
	return cnt
}

