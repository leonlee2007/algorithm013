# 学习笔记

## 字典树和并查集

###  字典树 Trie

#### 字典树的数据结构

字典树，即Trie树，又称单词查找树或键树，是一种树形结构。
典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。

它的优点是：最大限度地减少无谓的字符串比较，查询效率比哈希表高。

#### 字典树的核心思想

Trie树的核心思想是空间换时间

利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的

#### 字典树的基本性质

* 结点本身不存在完整单词
* 从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串
* 每个节点的所有子节点路径代表的字符都不相同

#### 参考链接

- [二叉树的层次遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)
- [实现 Trie](https://leetcode-cn.com/problems/implement-trie-prefix-tree/solution/)
- [Tire 树代码模板](https://shimo.im/docs/DP53Y6rOwN8MTCQH)

#### 实战题目

- [实现 Trie (前缀树) ](https://leetcode-cn.com/problems/implement-trie-prefix-tree/#/description)（亚马逊、微软、谷歌在半年内面试中考过）
- [单词搜索 II ](https://leetcode-cn.com/problems/word-search-ii/)（亚马逊、微软、苹果在半年内面试中考过）
- 分析单词搜索 2 用 Tire 树方式实现的时间复杂度

### 并查集 Disjoint Set

#### 适用场景

* 组团和配对问题
* group or not ？

#### 基本操作

* makeSet(s): 建立一个新的并查集，其中包含s个单元素集合
* unionSet(x, y): 把元素x和元素y所在的集合合并，要求x和y所在的集合不相交，如果相交则不合并
* find(x)：找到元素x所在的集合的代表，该操作也可以用于判断两个元素是否位于同一个集合，只要将他们各自的代表比较一下就可以了。
  - 初始化
  - 查询，合并
  - 路径压缩

#### 参考链接

- [并查集代码模板](https://shimo.im/docs/VtcxL0kyp04OBHak)

#### 实战题目

- [朋友圈](https://leetcode-cn.com/problems/friend-circles)（亚马逊、Facebook、字节跳动在半年内面试中考过）
- [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)（近半年内，亚马逊在面试中考查此题达到 361 次）
- [被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/)（亚马逊、eBay、谷歌在半年内面试中考过）

## 高级搜索

### 初级搜索

1. 朴素搜索

2. 优化方式：不重复（fibonacci）,剪枝（生成括号问题）

3. 搜索方向：

   DFS：depth first search 深度优先搜索

   BFS：breadth first search 广度优先搜索

   双向搜索，启发式搜索

### 剪枝的实现和特性

#### 回溯法

回溯法采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确解答的时候，它将取消上一步甚至是上几步的计算，在通过其他可能的分步解答再次尝试寻找问题的答案。

回溯法 通常用最简单的递归方法来实现，在反复重复上述的步骤后可能出现两种情况：

* 找到一个可能存在的正确的答案
* 在尝试了所有可能的分步方法后宣告该问题没有答案

在最坏的情况下，回溯法会导致一次复杂度为指数时间的计算

#### 参考链接

- [DFS 代码模板](https://shimo.im/docs/UdY2UUKtliYXmk8t)
- [BFS 代码模板](https://shimo.im/docs/ZBghMEZWix0Lc2jQ)
- [AlphaZero Explained](https://nikcheerla.github.io/deeplearningschool/2018/01/01/AlphaZero-Explained/)
- [棋类复杂度](https://en.wikipedia.org/wiki/Game_complexity)

#### 实战题目

- [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)（阿里巴巴、腾讯、字节跳动在半年内面试常考）
- [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)（亚马逊、Facebook、字节跳动在半年内面试中考过）
- [N 皇后](https://leetcode-cn.com/problems/n-queens/)（亚马逊、苹果、字节跳动在半年内面试中考过）
- [有效的数独](https://leetcode-cn.com/problems/valid-sudoku/description/)（亚马逊、苹果、微软在半年内面试中考过）
- [解数独](https://leetcode-cn.com/problems/sudoku-solver/#/description)（亚马逊、华为、微软在半年内面试中考过）

####  LeetCode 讨论区代码剖析

• https://leetcode.com/problems/n-queens/discuss/19808/ Accepted-4ms-c%2B%2B-solution-use-backtracking-and-bitmask- easy-understand

• https://leetcode.com/problems/n-queens/discuss/19810/Fast-short- and-easy-to-understand-python-solution-11-lines-76ms

### 双向BFS

#### 实战题目

- [单词接龙](https://leetcode-cn.com/problems/word-ladder/)（亚马逊、Facebook、谷歌在半年内面试中考过）
- [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)（谷歌、Twitter、腾讯在半年内面试中考过）

### 启发式所搜索（A*）

#### 估计函数

​	启发式函数：h(n), 它用来评价哪些结点最有希望的是一个我们要找的结点，h(n)会返回一个非负实数，也可以认为是从结点n的目标结点路径的估计成本。

​	启发式函数是一种告知搜索的方向的方法。它提供了一种明智的方法来猜测哪个邻居结点会导向一个目标。

#### 参考链接

- [A* 代码模板](https://shimo.im/docs/8CzMlrcvbWwFXA8r)
- [相似度测量方法](https://dataaspirant.com/2015/04/11/five-most-popular-similarity-measures-implementation-in-python/)
- [二进制矩阵中的最短路径的 A* 解法](https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/313347/A*-search-in-Python)
- [8 puzzles 解法比较](https://zxi.mytechroad.com/blog/searching/8-puzzles-bidirectional-astar-vs-bidirectional-bfs/)

#### 实战题目

- [二进制矩阵中的最短路径](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)（亚马逊、字节跳动、Facebook 在半年内面试中考过）
- [滑动谜题](https://leetcode-cn.com/problems/sliding-puzzle/)（微软、谷歌、Facebook 在半年内面试中考过）
- [解数独](https://leetcode-cn.com/problems/sudoku-solver/)（微软、华为、亚马逊在半年内面试中考过）

####  Shortest Path

1. BFS: 经典的BFS代码

2. A* search

   估价函数:

   h(current_point) = dist(current_point, destination_point)

   https://dataaspirant.com/2015/04/11/five-most-popular-similarity- measures-implementation-in-python/

3. https://leetcode.com/problems/shortest-path-in-binary-matrix/ discuss/313347/A*-search-in-Python

####  Sliding Puzzle

1. BFS: 经典的BFS代码: https://leetcode-cn.com/problems/sliding-puzzle/submissions/

2. A* search

   估价函数:

   h(current_state) = distance(current_state, target_state)

3. https://zxi.mytechroad.com/blog/searching/8-puzzles- bidirectional-astar-vs-bidirectional-bfs/

## 高级树，AVL树和红黑树

### 参考链接

- [维基百科：平衡树](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree)

###  二叉搜索树 Binary Search Tree

二叉搜索树，也称二叉搜索树、有序二叉树(Ordered Binary Tree)、排 序二叉树(Sorted Binary Tree)，是指一棵空树或者具有下列性质的二叉 树:

1. 左子树上所有结点的值均小于它的根结点的值;

2. 右子树上所有结点的值均大于它的根结点的值;

3. 以此类推:左、右子树也分别为二叉查找树。 (这就是 重复性!)

   中序遍历:升序排列

###  保证性能的关键

1. 保证二维维度! —> 左右子树结点平衡(recursively)

2. Balanced

3. https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree

### 如何平衡

###  AVL 树

1. 发明者 G. M. Adelson-Velsky 和 Evgenii Landis

2. Balance Factor(平衡因子): 是它的左子树的高度减去它的右子树的高度(有时相反)。 balancefactor={-1, 0, 1}

3. 通过旋转操作来进行平衡(四种)
   - 左旋 
   - 右旋 
   - 左右旋
   - 右左旋
   - [带有子树的旋转 ](https://zhuanlan.zhihu.com/p/63272157)

4. https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree

5.  AVL 总结

* 平衡二叉搜索树

* 每个结点存balancefactor={-1,0,1}

* 四种旋转操作

  不足:结点需要存储额外信息、且调整次数频繁

### 红黑树 Red-black Tree

​	红黑树是一种近似平衡的二叉搜索树(Binary Search Tree)，它能够确保任何一 个结点的左右子树的高度差小于两倍。具体来说，红黑树是满足如下条件的二叉 搜索树:

*  每个结点要么是红色，要么是黑色

*  根结点是黑色

*  每个叶结点(NIL结点，空结点)是黑色的。

* 不能有相邻接的两个红色结点

* 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点。

   关键性质 从根到叶子的最长的可能路径不多于最短的可能路径的两倍长



### AVL Trees和Red Black Trees 对比

* AVL trees provide faster lookups than Red Black Trees because they are more strictly balanced.

* Red Black Trees provide faster insertion and removal operations than AVL trees as fewer rotations are done due to relatively relaxed balancing.

* AVL trees store balance factors or heights with each node, thus requires storage for an integer per node whereas Red Black Tree requires only 1 bit of information per node.

*  Red Black Trees are used in most of the language libraries

  like map, multimap, multisetin C++whereas AVL trees are used in databases where faster retrievals are required.