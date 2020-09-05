// 91. 解码方法
// 一条包含字母 A-Z 的消息通过以下方式进行了编码：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。

// 示例 1:

// 输入: "12"
// 输出: 2
// 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
// 示例 2:

// 输入: "226"
// 输出: 3
// 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

package week_06

import(
  "fmt"
  "testing"
)

func TestNumDecodings(t * testing.T) {
  fmt.Println("============Start TestNumDecodings=============")
  s := "226"
  ret := numDecodings(s)
  fmt.Println("input:", s)
  fmt.Println("output:", ret)
  fmt.Println("============End TestNumDecodings=============")
}

// 方法一：
//动态规划
func numDecodings(s string) int {
    if s[0] == '0' {return 0}
    priv, now := 1, 1
    for i := 1; i < len(s); i++ {
        m := (s[i - 1]-'0') * 10 + (s[i] - '0')
        if (m == 0 || m > 26) && s[i] == '0' {
            return 0
        } else if s[i] == '0' {
            priv, now = now, priv
        } else if m >= 1 && m <= 26 && s [i - 1] != '0'{
            priv, now = now, now + priv
        } else {
            priv, now = now, now
        }
    }
    return now
}

