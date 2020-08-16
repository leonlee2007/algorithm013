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



## 分治（Divide & Conquer）

- 最近重复性（分治，回溯，递归）
- 最优重复性（动态规划）
- 化解多个子问题
- Problem -> Sub-Problem (Divide)-> Sub-Problem(Conquer) -> Solution(Merge)

### 分治代码模板

```python
def devide_conquer(problem, param1, param2, ...):
  # recursion terminator 问题解决了
  if problem is None:
    print_result
    return
  # prepare data 处理当前逻辑
  data = prepare_data(problem)
  subproblems = split_problem(problem, data)
  # conquer subproblems 下探到下一层
  subresult1 = self.devide_conquer(subproblems[0], p1, ...)
  subresult2 = self.devide_conquer(subproblems[1], p2, ...)
	subresult3 = self.devide_conquer(subproblems[3], p3, ...)
  ...
  # process and generate the final result 组装结果
  result = process_result(subresult1, subresult2, subresult3, ...)
  # revert the current level states
  
```

1. terminator
2. process (split your big problem)
3. drill down (subproblemes),
4. merge(subsult)
5. revers states

自顶向下，当前层解决当前的问题

## 回溯（Backtracking）

回溯法采用试错的思想，它尝试分布的去解决一个问题。 通过尝试，发现分布答案不能得到有效的正确解答时，它将取消上一步甚至上几步的计算，在通过其他的可能的分布解答再次尝试寻找问题的答案。

回溯法通常用最简单的递归方法来实现，在反复上述的步骤后可能出现两种情况：

* 找到一个可能存在的正确答案；
* 在尝试了所有可能的分布方法后宣告问题没有答案

到每一层去尝试，本质上还是递归的思想

### 回溯公式

```go
result := []
func backtrack(路径，选择列表) {
	if 满足结束条件 {
		result.add(路径)
	}
	return
	for 选择 in 选择列表 {
		做选择
		backtrack(路径，选择列表)
		撤销选择
	}
}
```

### 实战练习题目

- [Pow(x, n) ](https://leetcode-cn.com/problems/powx-n/)（Facebook 在半年内面试常考）

- [子集](https://leetcode-cn.com/problems/subsets/)（Facebook、字节跳动、亚马逊在半年内面试中考过）

### 参考链接

  - [牛顿迭代法原理](http://www.matrix67.com/blog/archives/361)
  - [牛顿迭代法代码](http://www.voidcn.com/article/p-eudisdmk-zm.html)

### 实战题目

  - [多数元素](https://leetcode-cn.com/problems/majority-element/description/) （亚马逊、字节跳动、Facebook 在半年内面试中考过）
  - [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)（亚马逊在半年内面试常考）
  - [N 皇后](https://leetcode-cn.com/problems/n-queens/)（字节跳动、苹果、谷歌在半年内面试中考过）
