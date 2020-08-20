# 学习笔记

## 数据结构

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

## 搜索-遍历

*	每个节点都要访问一次
*	每个节点仅仅要访问一次
*	对于节点的访问顺序不限
  - 深度优先：depth first search
  - 广度优先：breadth first search
  - 优先级优先：启发式搜索
  - 其他优先

## 深度优先搜索 Depth-First-Search（DFS 代码模板） 

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

##  广度优先搜索 Breadth-First-Search(BFS 代码模板)

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

## 实战题目

- [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/#/description)（字节跳动、亚马逊、微软在半年内面试中考过）
- [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/#/description)
- [括号生成](https://leetcode-cn.com/problems/generate-parentheses/#/description)（字节跳动、亚马逊、Facebook 在半年内面试中考过）
- [在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/#/description)（微软、亚马逊、Facebook 在半年内面试中考过）

## 课后作业

- [单词接龙](https://leetcode-cn.com/problems/word-ladder/description/)（亚马逊在半年内面试常考）
- [单词接龙 II ](https://leetcode-cn.com/problems/word-ladder-ii/description/)（微软、亚马逊、Facebook 在半年内面试中考过）
- [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)（近半年内，亚马逊在面试中考查此题达到 350 次）
- [扫雷游戏](https://leetcode-cn.com/problems/minesweeper/description/)（亚马逊、Facebook 在半年内面试中考过）

## 参考链接

- [DFS 代码模板（递归写法、非递归写法）](https://shimo.im/docs/UdY2UUKtliYXmk8t/)
- [BFS 代码模板](https://shimo.im/docs/ZBghMEZWix0Lc2jQ/)



