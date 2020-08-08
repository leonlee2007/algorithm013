package week_02

import(
  "fmt"
  "testing"
  "strings"
  "sort"
)

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 示例 1:
// 输入: s = "anagram", t = "nagaram"
// 输出: true

// 示例 2:
// 输入: s = "rat", t = "car"
// 输出: false

// 说明:
// 你可以假设字符串只包含小写字母。
// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

func TestIsAnagram(tt * testing.T) {
  fmt.Println("============Start TestIsAnagram=============")
  s, t := "anagram", "nagaram"
  ret := isAnagram5(s, t)
  fmt.Println("input:", s, t)
  fmt.Println("output:", ret)
  fmt.Println("============End TestIsAnagram=============")
}


// 方法一： 一个个拿出来对比下，如果有，就删掉原来的字符串的字母
func isAnagram1(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  for _, elem := range t {
    if strings.Contains(s, string(elem)) {
      s = strings.Replace(s, string(elem), "", 1)
    }
  }
  return s == ""
}


// 方法2：排序对比
// 时间复杂度O(nlogn)
// 空间复杂度O(n)
func isAnagram2(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  sl := []byte(s)
  tl := []byte(t)
  sort.Slice(sl, func(i, j int) bool {
    return sl[i] < sl[j]
  })
  sort.Slice(tl, func(i, j int) bool {
    return tl[i] < tl[j]
  })
  for i, elem := range sl {
    if elem != tl[i] {
      return false
    }
  }
  return true
}

// 方法3：哈希表1
func isAnagram3(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  count := make(map[uint8]int)
  for i := 0; i < len(s); i++ {
    count[s[i]]++
    count[t[i]]--
  }
  for _, elem := range count {
    if elem != 0 {
      return false
    }
  }
  return true
}


// 方法4：哈希表2
func isAnagram4(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  count := make(map[uint8]int)
  for i := 0; i < len(s); i++ {
    count[s[i]]++
  }
  for i := 0; i < len(t); i++ {
    count[t[i]]--
    if count[t[i]] < 0 {
      return false
    }
  }
  return true
}

// 方法5：使用数组
// 创建一个数组，一共26个格子，对应26个字母
// 如果s有就在那个位置+1，t就-1
func isAnagram5(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  m := make([]int, 26)
  for i := 0; i < len(s); i++ {
    m[s[i]-'a']++
    m[t[i]-'a']--
  }
  for _, elem := range m {
    if elem != 0 {
      return false
    }
  }
  return true
}

