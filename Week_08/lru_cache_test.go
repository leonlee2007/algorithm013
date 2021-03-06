// 146. LRU缓存机制
// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 Get 和 写入数据 Put 。

// 获取数据 Get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
// 写入数据 Put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

// 进阶:

// 你是否可以在 O(1) 时间复杂度内完成这两种操作？


// 示例:

// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

// cache.Put(1, 1);
// cache.Put(2, 2);
// cache.Get(1);       // 返回  1
// cache.Put(3, 3);    // 该操作会使得关键字 2 作废
// cache.Get(2);       // 返回 -1 (未找到)
// cache.Put(4, 4);    // 该操作会使得关键字 1 作废
// cache.Get(1);       // 返回 -1 (未找到)
// cache.Get(3);       // 返回  3
// cache.Get(4);       // 返回  4
package week_08

import(
  "fmt"
  "testing"
)

func TestLRUCache(t * testing.T) {
  fmt.Println("============Start TestLRUCache=============")
  cache := Constructor(2/* 缓存容量 */)
  cache.Put(1, 1)
  cache.Put(2, 2)
  cache.Get(1)      // 返回  1
  cache.Put(3, 3)   // 该操作会使得关键字 2 作废
  cache.Get(2)      // 返回 -1 (未找到)
  cache.Put(4, 4)   // 该操作会使得关键字 1 作废
  cache.Get(1)      // 返回 -1 (未找到)
  cache.Get(3)      // 返回  3
  cache.Get(4)      // 返回  4
  fmt.Println("cache:", cache)
  fmt.Println("============End TestLRUCache=============")
}

type LRUCache struct {
    size int
    capacity int
    cache map[int]*DLinkedNode
    head, tail *DLinkedNode
}

type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
    return &DLinkedNode{
        key: key,
        value: value,
    }
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDLinkedNode(0, 0),
        tail: initDLinkedNode(0, 0),
        capacity: capacity,
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    node := this.cache[key]
    this.moveToHead(node)
    return node.value
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; !ok {
        node := initDLinkedNode(key, value)
        this.cache[key] = node
        this.addToHead(node)
        this.size++
        if this.size > this.capacity {
            removed := this.removeTail()
            delete(this.cache, removed.key)
            this.size--
        }
    } else {
        node := this.cache[key]
        node.value = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
    node := this.tail.prev
    this.removeNode(node)
    return node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
