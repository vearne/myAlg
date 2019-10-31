package main

import (
	"fmt"
	"math/rand"
)

func init(){
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
}
//const MaxLevel = 32
//const Probability = 0.25 // 基于时间与空间综合 best practice 值, 越上层概率越小

const MaxLevel = 10
const Probability = 0.5 // 基于时间与空间综合 best practice 值, 越上层概率越小

func randLevel() (level int) {

	for level = 1; rand.Float32() < Probability && level < MaxLevel; level++ {
		// fmt.Println(rand.Float32())
	}
	// fmt.Printf("up to %d level\n", level)
	return
}

type node struct {
	forward []*node
	key     int
}

type skipList struct {
	head  *node
	level int
}

func newNode(key, level int) *node {
	return &node{key: key, forward: make([]*node, level)}
}

func newSkipList() *skipList {
	return &skipList{head: newNode(0, MaxLevel), level: 1}
}

func (s *skipList) insert(key int) {
	current := s.head
	update := make([]*node, MaxLevel) // 新节点插入以后的前驱节点
	for i := s.level - 1; i >= 0; i-- {
		if current.forward[i] == nil || current.forward[i].key > key {
			update[i] = current
		} else {
			for current.forward[i] != nil && current.forward[i].key < key {
				current = current.forward[i] // 指针往前推进
			}
			update[i] = current
		}
	}

	level := randLevel()
	if level > s.level {
		// 新节点层数大于跳表当前层数时候, 现有层数 + 1 的 head 指向新节点
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}
	node := newNode(key, level)
	for i := 0; i < level; i++ {
		node.forward[i] = update[i].forward[i]
		update[i].forward[i] = node
	}
}

func (s *skipList) delete(key int) {
	current := s.head
	for i := s.level - 1; i >= 0; i-- {
		for current.forward[i] != nil {
			if current.forward[i].key == key {
				tmp := current.forward[i]
				current.forward[i] = tmp.forward[i]
				tmp.forward[i] = nil
			} else if current.forward[i].key > key {
				break
			} else {
				current = current.forward[i]
			}
		}
	}

}

func (s *skipList) search(key int) *node {
	// 类似 delete
	return nil
}

func (s *skipList) print() {
	fmt.Println()

	for i := s.level - 1; i >= 0; i-- {
		current := s.head
		for current.forward[i] != nil {
			fmt.Printf("%d ", current.forward[i].key)
			current = current.forward[i]
		}
		fmt.Printf("***************** Level %d \n", i+1)
	}
}

func main() {
	list := newSkipList()
	for i := 0; i < 20; i++ {
		list.insert(rand.Intn(100))
	}
	list.print()

	fmt.Println("\n--------------------------------------")

	list.delete(88)
	list.print()


	list.delete(6)
	list.print()

	list.delete(12)
	list.print()

	fmt.Println("\n--------------------------------------")
}
