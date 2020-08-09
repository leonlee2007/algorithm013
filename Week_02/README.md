## 学习笔记

#### 哈希表（Hash table） 
    也叫散列表，是根据关键码值（Key value）而直接进行访问的数据结构
  通过关键码映射到表中一个位置来访问记录，以加快查找的速度
  这个映射函数叫作散列函数（Hash Function)
  存放记录的数组叫作哈希表（或散列表）



#### 实战练习题目
1. [有效的字母异位词 No.242](https://leetcode-cn.com/problems/valid-anagram/description/)
2. [字母异位词分组 No.49](https://leetcode-cn.com/problems/group-anagrams/)
3. [两数之和 No.1](https://leetcode-cn.com/problems/two-sum/description/)

#### 树、二叉树、二叉搜索树的实现和特性
    Linked List 是特殊化的树
    Tree 是特殊化的Graph

##### 二叉树遍历
    1.前序（Pre-order）根-左-右
    2.中序（In-order）左-根-右
    3.后序（Post-order）左-右-根
##### 二叉树示例代码
    C ++
    Struct TreeNode {
      int val;
      TreeNode *left;
      TreeNode *right;
      TreeNode(int X) : val(x), left(NULL), right(NULL){}
    }

    Python
    前序遍历
    def preorder(self, root):
      if root;
        self.traverse_path.append(root.val)
        self.preorder(root.left)
        self.preorder(root.right)

    中序遍历
    def inorder(self, root):
      if root;
        self.inorder(root.left)
        self.traverse_path.append(root.val)
        self.inorder(root.right)

    后序遍历
    def postorder(self, root):
      if root;
        self.postorder(root.left)
        self.postorder(root.right)
        self.traverse_path.append(root.val)


##### 二叉搜索树
    二叉搜索树叶也称为二叉排序树，
    有序二叉树（Order Binary Tree）
    排序二叉树（sorted Binary Tree）
    是指一课空树或者具有以下性质的二叉树:
    1. 左子树上所有节点上的值均小于它的根节点的值
    2. 右子树上所有节点上的值均大于它的根节点的值
    3. 以此类推：左右子树也分别为二叉查找树

    中序遍历就是升序排列
[二叉搜索树 Demo](https://visualgo.net/zh/bst)

#### 实战练习题目
1. [二叉树的中序遍历 No.94](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
2. [二叉树的前序遍历 No.144](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)
3. [N叉树的后序遍历 No.590](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)
4. [N叉树的前序遍历 No.589](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/)

5. [N叉树的层序遍历 No.429](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

#### 堆（Heap） 
    Heap:可以迅速找到一堆数中的最大或者最小值的数据结构
    根节点最大的叫大根堆，根节点最小的叫小根堆
    常见的有二叉堆，斐波那契堆
    常见操作（API）
    find-max: O(1)
    delete-max: O(logN)
    insert(creat): O(logN) or O(1)
#### 二叉堆
    通过完全二叉树来实现（不是二叉搜索树）
    二叉堆（大顶）满足下列性质：
    1. 是一棵完全树
    2. 树中任意节点的值总是 >= 其子节点的值 
[维基百科：堆（Heap）](https://en.wikipedia.org/wiki/Heap_(data_structure))
[堆的实现代码](https://shimo.im/docs/Lw86vJzOGOMpWZz2/read)
#### 二叉堆实现细节
    1.二叉堆一般通过数组实现
    2.第一个元素索引为0的情况下，父子节点的关系如下
        索引为i的左孩子的索引是2 * i+ 1
        索引为i的右孩子的索引是2 * i+ 2
        索引为i的父节点的索引是floor((i-1)/2)
    3. 二叉堆是堆的一种简单实现；但不是最优实现

#### 实战练习题目
1. [剑指 Offer 40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)
2. [滑动窗口最大值 No.239](https://leetcode-cn.com/problems/sliding-window-maximum/)

#### Homework
1. [HeapSort ：自学](https://www.geeksforgeeks.org/heap-sort/)
2. [剑指 Offer 49. 丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)
3. [前 K 个高频元素 No.347](https://leetcode-cn.com/problems/top-k-frequent-elements/)

#### 图的属性（Graph） 
 * Graph(V, E)
 * V - vertex;点
      1. 度 - 入度和出度
      2. 点与点是否连通
 * E - edge: 边
      1. 有向和无向（单行线）
      2. 权重（边长）

#### 图的表示
  * 临接矩阵 Adjacency matrix
  * 临接表 Adjacence list 

#### 图的分类 
  * 有向有权
  * 有向无权
  * 无向有权
  * 无向无权

#### 图的常见算法
  * DFS代码-递归写法
  ```
    visited = set() #和树中的DFS最大的区别, 因为树没有环路
    def dfs(node, visted)
      if node in visited:#terminator
        #already visted
        return
        visited.add(node)
        #process current node here.
        ...
        for next_node in node.childred():
          if not next_node in visited:
            dfs(next_node, visitd)
  ```
  * BFS代码
  ```
    def BFS(graph, start, end)
      queue = []
      queue.append([start])
      visited = set() #和树中的BFS的最大区别
      whild queue:
        node = queue.pop()
        visited.add(node)
        process(node)
        nodes = generate_related_nodes(node)
        queue.push(nodes)
  ```
####图的高级算法
[连通图个数-岛屿数量 No.200](https://leetcode-cn.com/problems/number-of-islands/)
[拓扑排序（Topological Sorting](https://zhuanlan.zhihu.com/p/34871092)
[最短路径（Shortest Path)](https://www.bilibili.com/video/av25829980?from=search&seid=13391343514095937158)
[最小生成树（Minimum Spanning Tree）](https://www.bilibili.com/video/av25829980?from=search&seid=13391343514095937158)
