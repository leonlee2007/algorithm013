# 学习笔记

## 递归（Recursion） 
    递归-循环
    通过函数体来进行循环

### 递归代码模板

```go
func recur(level int, param int) {
  //recursion terminator(递归终结者)
  //把函数终止条件写上。否则会死循环，只有递没有归！！！
  if (level > MAX_LEVEL) {
    return
  }
  //process current logic（处理当前层逻辑）
  //完成本层需要的逻辑
  process(level, param)
  //drill down（下探到下一层）
    //需要用参数标记当前是哪一层并把参数放进去
  recur(level + 1, newParam)
  //reverse the current level status if needed(清理当前层)
  //如果递归完了，这一层有些东西可能需要清理
}
```
		1. terminator
		2. process
		3. drill down
		4. reverse status
### 思维要点

	1. 抵制人肉递归
		不要人肉递归（即用人类的思维一个一个去枚举），为难自己。这是递归的一大思维误区。
		当我们看递归代码时，总是忍不住跟着递归一遍又一遍。
		简单的场景还好，复杂一些的，很容易进得去出不来，迷失了。
		前期可以画出递归树，但后期要逐渐养成直接看函数本身就开始写。
	2. 找最近重复子问题（计算机最擅长计算具有重复性的问题）
		找到最小可重复的步骤
		找到终止条件
	3. 数学归纳思维
		从n成立推导出n+1也成立。没有反例即正确

### 实战练习题目

1. [爬楼梯 No.70]( https://leetcode-cn.com/problems/climbing-stairs/)
2. [ 括号生成 No.22](https://leetcode-cn.com/problems/generate-parentheses/)
3. [翻转二叉树 No.226]( https://leetcode-cn.com/problems/invert-binary-tree/description/)
4. [验证二叉搜索树 No.98]( https://leetcode-cn.com/problems/validate-binary-search-tree)
5. [ 二叉树的最大深度 No.104]( https://leetcode-cn.com/problems/maximum-depth-of-binary- tree)
6. [ 二叉树的最小深度 No.111]( https://leetcode-cn.com/problems/minimum-depth-of-binary-tree)
7. [二叉树的序列化与反序列化 No.297]( https://leetcode-cn.com/problems/serialize-and-deserialize- binary-tree/)
8. [二叉树的最近公共祖先 No.236]( https://leetcode-cn.com/problems/lowest-common-ancestor- of-a-binary-tree/)
9. [从前序与中序遍历序列构造二叉树 No.105]( https://leetcode-cn.com/problems/construct-binary-tree-from- preorder-and-inorder-traversal)
10. [组合 No.77](https://leetcode-cn.com/problems/combinations/)
11. [全排列 No.46]( https://leetcode-cn.com/problems/permutations/ https://leetcode-cn.com/problems/permutations-ii/)
12. [全排列 II No.47]( https://leetcode-cn.com/problems/permutations/ https://leetcode-cn.com/problems/permutations-ii/)



