// 49. 字母异位词分组
// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

// 示例:
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出:
// [
// ["ate", "eat", "tea"]
// ["nat", "tan"],
// ["bat"]
// ]

// 说明：
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。

package week_02

import(
  "fmt"
  "testing"
  "sort"
)

func TestGroupAnagrams(tt * testing.T) {
  fmt.Println("============Start TestGroupAnagrams=============")
  strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
  ret := groupAnagrams(strs)
  fmt.Println("input:", strs)
  fmt.Println("output:", ret)
  fmt.Println("============End TestGroupAnagrams=============")
}

// 方法一 golang中数组是可比较类型，可直接用作映射的key。
// 思路: 计算每个单词的key, 将具有相同的key放入同一个数组中。
// 时间复杂度: O(NK) N是strs的长度，K是strs字符串的最大长度
// 空间复现度: O(NK)
func groupAnagrams1(strs []string) [][]string {
  m := make(map[[26]int][]string)
  for _, str := range strs {
    k := strArray(str)
    s, ok := m[k]
    if !ok {
      s = make([]string, 0)
    }
    s = append(s, str)
    m[k] = s
  }
  ret := make([][]string, 0, len(m))
  for _, v := range m {
    ret = append(ret, v)
  }
  return ret
}

func strArray(s string) [26]int {
  ret := [26]int{}
  for _, v := range s {
    ret[v - 'a']++
  }
  return ret
}

// 方法二 用质数表示26个字母，把字符串的各个字母相乘，这样可保证字母异位词的乘积必定是相等的
// 思路: 计算每个单词的key, 将具有相同的key放入同一个数组中。
// 时间复杂度: O(NK) N是strs的长度，K是strs字符串的最大长度
// 空间复现度: O(NK)
func groupAnagrams2(strs []string) [][]string {
  m := make(map[int][]string)
  for _, str := range strs {
    k := gen_key(str)
    s, ok := m[k]
    if !ok {
      s = make([]string, 0)
    }
    s = append(s, str)
    m[k] = s
  }
  ret := make([][]string, 0, len(m))
  for _, v := range m {
    ret = append(ret, v)
  }
  return ret
}

// 26个质数集合
var primes = [26]int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}
func gen_key(s string) int {
  ret := 0
  for i := 0; i < len(s); i++ {
    if ret == 0 {
      ret = primes[s[0] - 'a']
    }else {
      ret *= primes[s[i] - 'a']
    }
  }
  return ret
}

// 方法三
// 思路： 
// 1: 遍历数组，然后对数组中的字符串排序
// 2: 用排序后的字符串作为key，记录在数组中出现的位置
// 3: 若该字符串未出现过记录, 加到ret后面, 否则加入到ret中已经出现的地方
// 时间复杂度: O(NKlogK) N是strs的长度，K是strs字符串的最大长度
// 空间复现度: O(NK)

func groupAnagrams3(strs []string) [][]string {
  ret := [][]string{}
  m := make(map[string]int)
  for _, str := range strs {
    kByte := bytes(str)
    //将字符串排序
    sort.Sort(kByte) 
    k := string(kByte) 
    if idx, ok := m[k]; !ok {
      //记录拍完序的字符串以及字符串在ret的位置
      m[k] = len(ret) 
      ret = append(ret, []string{str})
    } else { 
      //已经出现过，将其放入出现在ret中，在ret的位置为idx
      ret[idx] = append(ret[idx], str)
    }
  }
  return ret
}

type bytes []byte
// 自定义排序规则
func (b bytes) Len() int {
  return len(b)
}
func (b bytes) Less(i,j int) bool {
  return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
  b[i], b[j] = b[j], b[i]
}
