// 680. 验证回文字符串 Ⅱ
// 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

// 示例 1:

// 输入: "aba"
// 输出: True
// 示例 2:

// 输入: "abca"
// 输出: True
// 解释: 你可以删除c字符。
// 注意:

// 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
package week_09

import(
  "fmt"
  "testing"
)

func TestValidPalindrome(t * testing.T) {
  fmt.Println("============Start TestValidPalindrome=============")
  s := "abca"
  ret := validPalindrome(s)
  fmt.Println("input:", s)
  fmt.Println("output:", ret)
  fmt.Println("============End TestValidPalindrome=============")
}


func validPalindrome(s string) bool {
    if len(s) == 0 {return true}
    low, high := 0, len(s) - 1
    for low < high{
        if s[low] == s[high] {
            low += 1
            high -= 1
        } else {
            return isPalindrome(s, low + 1, high) || isPalindrome(s, low, high - 1)
        }
    }
    return true
}

func isPalindrome(s string, low int, high int) bool{
    for low < high{
        if s[low] != s[high] {return false}
        low += 1
        high -= 1
    }
    return true
 }

