# 学习笔记

## 广度优先和深度优先

### 数据结构

```python
class TreeNode:
  def __init__(self, val):
    self.val = val
    self.left, self.right = None, None
```

```c++
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x):val(x), left:(NULL), right(NULL){}
}
```

```go
type TreeNode struct {
	val   int
	left, right *TreeNode
}

```

### 搜索-遍历

*	每个节点都要访问一次
*	每个节点仅仅要访问一次
*	对于节点的访问顺序不限
  - 深度优先：depth first search
  - 广度优先：breadth first search
  - 优先级优先：启发式搜索
  - 其他优先

### 深度优先搜索 Depth-First-Search（DFS 代码模板） 

#### 递归写法

```python
#Python
visited = set() 

def dfs(node, visited):
    if node in visited: # terminator
    	# already visited 
    	return 

	visited.add(node) 

	# process current node here. 
	...
	for next_node in node.children(): 
		if next_node not in visited: 
			dfs(next_node, visited)
```
```c++
//C/C++
//递归写法：
map<int, int> visited;
void dfs(Node* root) {
  // terminator
  if (!root) return ;

  if (visited.count(root->val)) {
    // already visited
    return ;
  }

  visited[root->val] = 1;
  // process current node here. 
  // ...
  for (int i = 0; i < root->children.size(); ++i) {
    dfs(root->children[i]);
  }
  return ;
}
```
```go
//Go DFS
//visited := make(map[int]int)
var visited map[*Node]int
func dfs(root *Node) {
  // terminator
  if root == nil {return}
    // already visited 
  if visited[root] == 1 {return}
  visited[root] = 1
  // process current node here. 
  // ...
  for i := 0; i <  len(root.children); i ++) {
    dfs(root.children[i])
  }
  return
}
```



#### 非递归写法

```python
#Python
def DFS(self, tree): 

	if tree.root is None: 
		return [] 
	visited, stack = [], [tree.root]
	while stack: 
		node = stack.pop() 
		visited.add(node)

		process (node) 
		nodes = generate_related_nodes(node) 
		stack.push(nodes) 

	# other processing work 
	...
```

```c++
//C/C++
//非递归写法：
void dfs(Node* root) {
  map<int, int> visited;
  if(!root) return ;

  stack<Node*> stackNode;
  stackNode.push(root);

  while (!stackNode.empty()) {
    Node* node = stackNode.top();
    stackNode.pop();
    if (visited.count(node->val)) continue;
    visited[node->val] = 1;

    for (int i = node->children.size() - 1; i >= 0; --i) {
        stackNode.push(node->children[i]);
    }
  }

  return ;
}
```

###  广度优先搜索 Breadth-First-Search(BFS 代码模板)

```python
# Python
def BFS(graph, start, end):
    visited = set()
	queue = [] 
	queue.append([start]) 
	while queue: 
		node = queue.pop() 
		visited.add(node)
		process(node) 
		nodes = generate_related_nodes(node) 
		queue.push(nodes)
	# other processing work 
	...
```

```c
// C/C++
void bfs(Node* root) {
  map<int, int> visited;
  if(!root) return ;

  queue<Node*> queueNode;
  queueNode.push(root);

  while (!queueNode.empty()) {
    Node* node = queueNode.top();
    queueNode.pop();
    if (visited.count(node->val)) continue;
    visited[node->val] = 1;

    for (int i = 0; i < node->children.size(); ++i) {
        queueNode.push(node->children[i]);
    }
  }
  return ;
}
```

```go
// GO BFS
//visited := make(map[int]int)
var visited map[*Node]int
func bfs(root *Node) { 
  if(root == nil) {return} ;	
  queue := []*Node{root}
  while (len(queue) > 0) {
    Node* node = queue[0]
    queue := queue[1:len(queue)]
    if visited[node] == 1 {continue}
    visited[node] = 1;
    for (int i = 0; i < len(node.children); i ++) {
      queue = append(queue, node.children[i])
    }
  }
  return
}
```

### 实战题目

- [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/#/description)（字节跳动、亚马逊、微软在半年内面试中考过）
- [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/#/description)
- [括号生成](https://leetcode-cn.com/problems/generate-parentheses/#/description)（字节跳动、亚马逊、Facebook 在半年内面试中考过）
- [在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/#/description)（微软、亚马逊、Facebook 在半年内面试中考过）

### 课后作业

- [单词接龙](https://leetcode-cn.com/problems/word-ladder/description/)（亚马逊在半年内面试常考）
- [单词接龙 II ](https://leetcode-cn.com/problems/word-ladder-ii/description/)（微软、亚马逊、Facebook 在半年内面试中考过）
- [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)（近半年内，亚马逊在面试中考查此题达到 350 次）
- [扫雷游戏](https://leetcode-cn.com/problems/minesweeper/description/)（亚马逊、Facebook 在半年内面试中考过）

### 参考链接

- [DFS 代码模板（递归写法、非递归写法）](https://shimo.im/docs/UdY2UUKtliYXmk8t/)
- [BFS 代码模板](https://shimo.im/docs/ZBghMEZWix0Lc2jQ/)

##  贪心算法 Greedy

* 贪心算法是一种在每一步选择中都采取在当前状态下最优选择，从而希望导致结果是全局最优的算法

* 贪心算法与动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能

* 贪心法可以解决一些最优化问题，如：求图中的最小生成树，求哈夫曼编码等。

* 一旦一个问题可以通过贪心法来解决，那么贪心法一般是解决这个问题的最好办法。

* 由于贪心法的高效性以及所求得的答案比较接近最优，贪心法也可用作辅助算法或者直接解决一些要求结果不特别精确的问题

  ### 适合贪心算法的场景

  ​	问题能够分解成子问题来解决，子问题的最优解能递推到最终问题的最优解。

  ​	这种子问题最优解称为最优子结构

  ​	贪心法难在如何证明能用贪心法。可以从前向后贪心也可以从后向前贪心

  ### 参考链接

  - [coin change 题目](https://leetcode-cn.com/problems/coin-change/)
  - [动态规划定义](https://zh.wikipedia.org/wiki/动态规划)

  ### 课后作业

  - [柠檬水找零](https://leetcode-cn.com/problems/lemonade-change/description/)（亚马逊在半年内面试中考过）
  - [买卖股票的最佳时机 II ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/)（亚马逊、字节跳动、微软在半年内面试中考过）
  - [分发饼干](https://leetcode-cn.com/problems/assign-cookies/description/)（亚马逊在半年内面试中考过）
  - [模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/description/)
  - [跳跃游戏](https://leetcode-cn.com/problems/jump-game/) （亚马逊、华为、Facebook 在半年内面试中考过）
  - [跳跃游戏 II ](https://leetcode-cn.com/problems/jump-game-ii/)（亚马逊、华为、字节跳动在半年内面试中考过）



## 二分查找

### 二分查找的前提

1. 目标函数单调性（单调递增或者递减）
2. 存在上下界（bounded）
3. 能够通过索引访问（index accessible）

### 代码模板

```c
//C/C++
int binarySearch(const vector<int>& nums, int target) {
	int left = 0, right = (int)nums.size()-1;	
	while (left <= right) {
	int mid = left + (right - left)/ 2;
	if (nums[mid] == target) return mid;
	else if (nums[mid] < target) left = mid + 1;
	else right = mid - 1;
	}	
	return -1;
}
```

```go
//Go
func binarySearch (array []int, target int) int {
  left, right := 0, len(array) -1
  for left <= right {
    mid := (right + left)/ 2;
    if array[mid] == target{
      return mid
    } else if array[mid] < target {
      left = mid + 1
    } else {
      right = mid -1
    }
  }
  return -1
}
```

### 参考链接

- [二分查找代码模板](https://shimo.im/docs/xvIIfeEzWYEUdBPD/)
- [Fast InvSqrt() 扩展阅读 ](https://www.beyond3d.com/content/articles/8/)

### 实战题目

- [x 的平方根](https://leetcode-cn.com/problems/sqrtx/)（字节跳动、微软、亚马逊在半年内面试中考过）牛顿迭代法
- [有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)（亚马逊在半年内面试中考过）

### 课后作业

- [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)（Facebook、字节跳动、亚马逊在半年内面试常考）
- [搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)（亚马逊、微软、Facebook 在半年内面试中考过）
- [寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)（亚马逊、微软、字节跳动在半年内面试中考过）
- 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方

