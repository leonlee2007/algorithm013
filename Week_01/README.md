## 学习笔记

#### 方法套路
- chunk it up 切碎知识点
- deliberate practicing 刻意练习
- feedback 反馈

#### 切题四件套：
 - Clarification
 - Possible solutions 
   - compare（time/space） 
   - optimal（加强）
- Coding（多写）
- Test cases

#### 五毒刷题法
- 刷题第一遍：
    - 5分钟：读题+思考
    - 直接看解法：多解法，比较解法优劣
    - 背诵 ，默写好的解法

- 刷题第二遍：
    - 马上直接写 --> leetcode提交
    - 多种解法比较，体会 --> 优化

- 刷题第三遍：
    - 24小时后，再重复做题
    - 不同解法的熟练程度 --> 专项练习

- 刷题第四遍：
    - 过了一周：反复回来练习相同题目

- 刷题第五遍：
    - 面试前一周恢复性训练

#### Big O notation
- O(1): Constant Complexity 常数复杂度
- O(log n): Logarithmic Complexity 对数复杂度
- O(n): Linear Complexity 线性时间复杂度
- O(n^2):N square Complexity 平方
- O(n^3): N cubic Complexity 立方
- O(2^n):Exponential Growth 指数
- O(n!):Factorial 阶乘
注意：只看最高复杂度的运算

#### 数组 Array
定义：是一种线性表数据结构，用一组连续的内存空间，来存储一组具有相同类型的数据。

特点：支持随机访问，根据下标随机访问的时间复杂度为 O(1)

缺点：想要在数组中插入一个元素或者删除一个元素，为了保证其连续性，就需要做大量的数据搬移工作。

时间复杂度: 插入和删除是O(n)

#### 链表 Linked List
定义：链表就是通过指针将一组零散的内存块串连在一起。

特点：
1、非连续的内存空间，为了能将所有节点串起来，每个结点都记录了下一个结点的内存地址，即后继指针
2、链表中插入或者删除一个数据，我们并不需要为了保持内存的连续性而搬移结点，因为链表的存储空间本身就不是连续的。所以，在链表中插入和删除一个数据是非常快速的。

缺点：
链表要想随机访问第 k 个元素，就没有数组那么高效了。因为链表中的数据并非连续存储的，所以无法像数组那样，根据首地址和下标，通过寻址公式就能直接计算出对应的内存地址，而是需要根据指针一个结点一个结点地依次遍历，直到找到相应的结点。需要 O(n) 的时间复杂度。

时间复杂度: 查询 O(n)

#### 双向链表 Double Linked List

#### 跳表(Skip List)
跳表，对标的是平衡树和二分查找，是一种插入/删除/搜索都是O(log n) 的数据结构，跳表里面的元素始终是有序的。

只能用于元素有序的情况
用于解决链表搜索的问题，给有序的链表加速

优势：原理简单，容易实现，方便扩展，效率最高

#### 实战练习题目 - Array
[盛最多水的容器 No.11](https://leetcode-cn.com/problems/container-with-most-water/)
[移动零 No.283](https://leetcode-cn.com/problems/move-zeroes/)
[爬楼梯 No.70](https://leetcode-cn.com/problems/climbing-stairs/)
[三数之和 No.15](https://leetcode-cn.com/problems/3sum/)

#### 实战练习题目 - Linked List
[反转链表 No.206](https://leetcode-cn.com/problems/reverse-linked-list/)
[两两交换链表中的节点 No.24](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)
[环形链表 No.141](https://leetcode-cn.com/problems/linked-list-cycle/)
[环形链表 II No.142](https://leetcode-cn.com/problems/linked-list-cycle-ii/)
[K 个一组翻转链表 No.25](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)

#### Homework
[删除排序数组中的重复项 No.26](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)
[旋转数组 No.189](https://leetcode-cn.com/problems/rotate-array/)
[合并两个有序链表 No.21](https://leetcode-cn.com/problems/merge-two-sorted-lists/)
[合并两个有序数组 No.88](https://leetcode-cn.com/problems/merge-sorted-array/)
[两数之和 No.1](https://leetcode-cn.com/problems/two-sum/)
[移动零 No.283](https://leetcode-cn.com/problems/move-zeroes/)
[加一 No.66](https://leetcode-cn.com/problems/plus-one/)


#### 栈(Stack)
先入后出；添加删除皆为O(1)

#### 队列(Queue)
先入先出；添加删除皆为O(1)

#### 双端队列(Deque:Double-End Queue)
1. 简单理解：两端可以进出的Queue
2. 插入和删除都是O(1)的操作

#### 优先队列(Priority Queue)
1. 插入操作O(1)
2. 取出操作O(log(N)) 按照元素优先级取出
3. 底层实现数据结构较为多样和复杂:heap，bst，treap


#### 预习题目
[有效的括号 No.20](https://leetcode-cn.com/problems/valid-parentheses/description/)
[最小栈 No.155](https://leetcode-cn.com/problems/min-stack/)

#### 实战题目
[柱状图中最大的矩形 No.84](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)
[滑动窗口最大值 No.239](https://leetcode-cn.com/problems/sliding-window-maximum/)

#### Homework
[Design Circular Deque No.641](https://leetcode.com/problems/design-circular-deque/)
[Trapping Rain Water No.42](https://leetcode.com/problems/trapping-rain-water/)