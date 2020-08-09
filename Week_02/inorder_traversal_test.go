
// 94. 二叉树的中序遍历
// 给定一个二叉树，返回它的中序 遍历。

package week_02

import(
  "fmt"
  "testing"
)

// Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func TestInorderTraversal(t * testing.T) {
  fmt.Println("============Start TestInorderTraversal=============")
  root := new(TreeNode)
  root.Val = 1

  node1 := new(TreeNode)
  node1.Val = 2

  node2 := new(TreeNode)
  node2.Val = 3

  node3 := new(TreeNode)
  node3.Val = 4
  node4 := new(TreeNode)
  node4.Val = 5

  root.Left = node1
  root.Right = node2
  node1.Left = node3
  node1.Right = node4

  ret := inorderTraversal4(root)
  fmt.Println("input:", root)
  fmt.Println("output:", ret)
  fmt.Println("============End TestInorderTraversal=============")
}

// 解法一 递归
var res []int

func inorderTraversal1(root *TreeNode) []int {
  res = make([]int, 0)
  dfs(root)
  return res
}

func dfs(root *TreeNode) {
  if root != nil {
    dfs(root.Left)
    res = append(res, root.Val)
    dfs(root.Right)
  }
}

// 解法二 迭代（stack）
// 逻辑与 解法一 递归 一致
// 自建 stack 负责 root 出栈入栈
func inorderTraversal2(root *TreeNode) []int {
  var res []int
  var stack []*TreeNode
  //root != nil 只为了第一次root判断，必须放最后
  for 0 < len(stack) || root != nil {
    for root != nil {
      //入栈
      stack = append(stack, root)
      //移至最左
      root = root.Left
    }
    //栈顶
    index := len(stack) - 1
    //中序输出
    res = append(res, stack[index].Val)
    //右节点会进入下次循环，如果 =nil，继续出栈
    root = stack[index].Right
    //出栈
    stack = stack[:index]
  }
  return res
}

// 解法三 Morris 破坏树结构 （二叉树转单向升序链表）

// 在 解法一 递归 解法二 迭代（stack) 中，必须借助 stack 来实现 中序遍历，增加空间复杂度 O(n)
// Morris 则在现有节点上进行节点关联，从而避免了 stack 空间复杂度 O(n) 的问题

// 寻找左子树 最大节点 指向当前节点
// 砍掉 当前节点 的 左子树
func inorderTraversal3(root *TreeNode) []int {
  var res []int
  var max *TreeNode
  for root != nil {
    if root.Left == nil {
      res = append(res, root.Val)
      root = root.Right
    } else {
      //寻找左树最大节点
      max = root.Left
      for max.Right != nil {
        max = max.Right
      }

      //最大节点指向root
      max.Right = root

      //root = root.Left :移动root到root.left
      //root.Left = nil  :砍左子树，避免下一次遍历到root时，再进入到左子树
      root, root.Left = root.Left, nil
    }
  }
  return res
}

// 解法四 Morris 保持树结构
// 在 解法三 Morris 破坏树结构 中
// 我们为了避免下一次遍历到root时，再进入到左子树，直接砍左子树
// 为了解决砍树问题，我们可以在 解法三 Morris 破坏树结构的基础上，增加

// 在下次遍历到 root 时，直接把 root加入结果
// 移动到root.Right 就可以避免再进入到左子树的死循环

func inorderTraversal4(root *TreeNode) []int {
  var res []int
  var max *TreeNode
  for root != nil {
    if root.Left == nil {
      res = append(res, root.Val)
      root = root.Right
    } else {
      max = root.Left
      for max.Right != nil && max.Right != root {
        max = max.Right
      }

      if max.Right == nil {
        // 未连接,连接到root
        max.Right = root
        root = root.Left
      } else {
        // 已连接,断开连接
        max.Right = nil
        res = append(res, root.Val)
        root = root.Right
      }
    }
  }
  return res
}
