# 学习笔记

## 动态规划

###  分治 + 回溯 + 递归 + 动态规划

- [递归代码模板](https://shimo.im/docs/EICAr9lRPUIPHxsH)
- [分治代码模板](https://shimo.im/docs/zvlDqLLMFvcAF79A)
- [动态规划定义](https://en.wikipedia.org/wiki/Dynamic_programming)

**本质:寻找重复性 —> 计算机指令集**

### 动态规划 Dynamic Programming

1. [Wiki 定义](https://en.wikipedia.org/wiki/Dynamic_programming)

2. “Simplifying a complicated problem by breaking it down into simpler sub-problems”

    (in a recursive manner)

3. Divide& Conquer + Optimal substructure 分治 + 最优子结构

### 动态规划 和分治 ,回溯 ,递归区别	

​	没有根本上的区别（关键看有无最优子结构）

​	共性：找到重复子问题

​	差异性:  最优子结构，中途可以淘汰次优解

###  动态规划关键点

1. 最优子结构 opt[n] = best_of(opt[n-1], opt[n-2], ...)

2. 储存中间状态:opt[i]

3. 递推公式(美其名曰:状态转移方程或者 DP 方程)

   Fib: opt[i] = opt[n-1] + opt[n-2]

   二维路径:opt[i,j] = opt[i+1][j] + opt[i][j+1] (且判断a[i,j]是否空

###  动态规划小结

1. 打破自己的思维惯性，形成机器思维 
2. 理解复杂逻辑的关键

3. 也是职业进阶的要点要领

### 实战题目1

- [不同路径](https://leetcode-cn.com/problems/unique-paths/)（Facebook、亚马逊、微软在半年内面试中考过）
- [不同路径 II ](https://leetcode-cn.com/problems/unique-paths-ii/)（谷歌、美团、微软在半年内面试中考过）
- [最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)（字节跳动、谷歌、亚马逊在半年内面试中考过）
- [MIT 动态规划课程最短路径算法](https://www.bilibili.com/video/av53233912?from=search&seid=2847395688604491997)

### 实战题目2

- [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/description/)（阿里巴巴、腾讯、字节跳动在半年内面试常考）

- [三角形最小路径和](https://leetcode-cn.com/problems/triangle/description/)（亚马逊、苹果、字节跳动在半年内面试考过）

- [三角形最小路径和高票回答](https://leetcode.com/problems/triangle/discuss/38735/Python-easy-to-understand-solutions-(top-down-bottom-up))

- [最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)（亚马逊、字节跳动在半年内面试常考）

- [乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/description/)（亚马逊、字节跳动、谷歌在半年内面试中考过）

- [零钱兑换](https://leetcode-cn.com/problems/coin-change/description/)（亚马逊在半年内面试中常考）

### 实战题目3

- [打家劫舍](https://leetcode-cn.com/problems/house-robber/)（字节跳动、谷歌、亚马逊在半年内面试中考过）
- [打家劫舍 II ](https://leetcode-cn.com/problems/house-robber-ii/description/)（字节跳动在半年内面试中考过）
- [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/#/description)（亚马逊、字节跳动、Facebook 在半年内面试中常考）
- [买卖股票的最佳时机 II ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)（亚马逊、字节跳动、微软在半年内面试中考过）
- [买卖股票的最佳时机 III ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)（字节跳动在半年内面试中考过）
- [最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)（谷歌、亚马逊在半年内面试中考过）
- [买卖股票的最佳时机 IV ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)（谷歌、亚马逊、字节跳动在半年内面试中考过）
- [买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)
- [一个方法团灭 6 道股票问题](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/)

### 高级 DP 实战题目

- [完全平方数](https://leetcode-cn.com/problems/perfect-squares/)（亚马逊、谷歌在半年内面试中考过）
- [编辑距离](https://leetcode-cn.com/problems/edit-distance/) **（重点）**
- [跳跃游戏](https://leetcode-cn.com/problems/jump-game/)（亚马逊在半年内面试中考过）
- [跳跃游戏 II ](https://leetcode-cn.com/problems/jump-game-ii/)（亚马逊、华为字节跳动在半年内面试中考过）
- [不同路径](https://leetcode-cn.com/problems/unique-paths/)（Facebook、亚马逊、微软在半年内面试中考过）
- [不同路径 II ](https://leetcode-cn.com/problems/unique-paths-ii/)（谷歌、美团、微软在半年内面试中考过）
- [不同路径 III ](https://leetcode-cn.com/problems/unique-paths-iii/)（谷歌在半年内面试中考过）
- [零钱兑换](https://leetcode-cn.com/problems/coin-change/)（亚马逊在半年内面试中常考）
- [零钱兑换 II ](https://leetcode-cn.com/problems/coin-change-2/)（亚马逊、字节跳动在半年内面试中考过）

### 本周作业

### 中等

- [最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)（亚马逊、高盛集团、谷歌在半年内面试中考过）
- [解码方法](https://leetcode-cn.com/problems/decode-ways)（亚马逊、Facebook、字节跳动在半年内面试中考过）
- [最大正方形](https://leetcode-cn.com/problems/maximal-square/)（华为、谷歌、字节跳动在半年内面试中考过）
- [任务调度器](https://leetcode-cn.com/problems/task-scheduler/)（Facebook 在半年内面试中常考）
- [回文子串](https://leetcode-cn.com/problems/palindromic-substrings/)（Facebook、苹果、字节跳动在半年内面试中考过）

### 困难

- [最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)（字节跳动、亚马逊、微软在半年内面试中考过）
- [编辑距离](https://leetcode-cn.com/problems/edit-distance/)（字节跳动、亚马逊、谷歌在半年内面试中考过）
- [矩形区域不超过 K 的最大数值和](https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/)（谷歌在半年内面试中考过）
- [青蛙过河](https://leetcode-cn.com/problems/frog-jump/)（亚马逊、苹果、字节跳动在半年内面试中考过）
- [分割数组的最大值](https://leetcode-cn.com/problems/split-array-largest-sum)（谷歌、亚马逊、Facebook 在半年内面试中考过）
- [学生出勤记录 II ](https://leetcode-cn.com/problems/student-attendance-record-ii/)（谷歌在半年内面试中考过）
- [最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/)（Facebook 在半年内面试中常考）
- [戳气球](https://leetcode-cn.com/problems/burst-balloons/)（亚马逊在半年内面试中考过）