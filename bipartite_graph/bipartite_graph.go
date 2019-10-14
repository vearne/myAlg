package main

// 二分图问题
// 
import (
	"container/list"
	"fmt"
)

var edges []Edge
var cnt int
var head []int

type Graph struct {
	currentNode int
	visted      []bool
	match       []int
}

func NewGraph(n int) *Graph {
	g := Graph{}
	g.visted = make([]bool, (n+1)*2)
	g.match = make([]int, (n+1)*2)
	for i := 0; i < len(g.match); i++ {
		g.match[i] = -1
	}
	return &g
}

func (g *Graph) Clone() *Graph {
	x := *g
	return &x
}

type Edge struct {
	next int //同一起点的上一条边的编号
	to   int //第i条边的终点
	//w    int
}

func add(u, v int) {
	//edge[cnt].w = w
	edges[cnt].to = v
	edges[cnt].next = head[u]
	head[u] = cnt
	cnt++
}

func main() {
	edges = make([]Edge, 100)
	head = make([]int, 100)

	// 男生与女生数量相同
	// 男生的数量
	n := 4

	cnt = 0
	for i := 0; i < len(head); i++ {
		head[i] = -1
	}
	//input := [][]int{{1, 4}, {1, 6}, {2, 5}, {3, 5}}
	//input := [][]int{{1, 4}, {2, 4}}
	//input := [][]int{{2, 5}, {2, 7}, {2, 8}, {3, 6}, {3, 8}, {4, 8}}
	input := [][]int{{2, 5}, {2, 6}, {3, 5}}

	//input := [][]int{{1, 4}, {2, 6}, {2, 5}}

	for _, item := range input {
		add(item[0], item[1])
		//add(item[1], item[0])
	}

	fmt.Println("head", head[0:10])
	fmt.Println("edge", edges[0:10])
	for i := 1; i <= 6; i++ {
		findTheSameStart(i)
	}

	query(n)
}

func findTheSameStart(node int) {
	e := head[node]
	for e != -1 {
		fmt.Println("相同起始节点:", node, "-->", "终止节点", edges[e].to)
		e = edges[e].next
	}
	fmt.Println("------")
}

type Queue struct {
	innerList *list.List
}

func NewQueue() *Queue {
	q := Queue{}
	q.innerList = list.New()
	return &q
}

func (q *Queue) push(g *Graph) {
	q.innerList.PushBack(g)
}

func (q *Queue) Size() int {
	return q.innerList.Len()
}

func (q *Queue) isEmpty() bool {
	//fmt.Println("queue.Size", q.Size())
	return !(q.Size() > 0)
}

func (q *Queue) pop() *Graph {
	if q.innerList.Len() > 0 {
		e := q.innerList.Front()
		//fmt.Println("--pop1--")
		q.innerList.Remove(e) // Dequeue

		return e.Value.(*Graph)
	} else {
		//fmt.Println("--pop2--")
		return nil
	}
}

func query(n int) {
	queue := NewQueue()
	for u := 1; u <= n; u++ {
		g := NewGraph(n)
		g.currentNode = u
		g.visted[u] = true
		queue.push(g)
	}

	result := 0           // 增广路经的长度
	finalMatch := []int{} // 匹配关系

	for !queue.isEmpty() {
		g := queue.pop()

		// 看看是否是最优的结果
		value, match := extractMaxMatch(g)
		if value > result {
			result = value
			finalMatch = match
		}

		// 表示从 v  -> u 的路径
		if g.currentNode > n {
			for u := 1; u <= n; u++ {
				if !g.visted[u] {
					newG := g.Clone()
					newG.visted[u] = true
					newG.match[g.currentNode] = u
					newG.currentNode = u
					queue.push(newG)
				}
			}
		} else { // 表示从 u  -> v 的路径
			for i := head[g.currentNode]; i != -1; i = edges[i].next {
				v := edges[i].to
				if !g.visted[v] {
					newG := g.Clone()
					newG.visted[v] = true
					newG.match[g.currentNode] = v
					newG.currentNode = v
					queue.push(newG)
				}
			}
		}
	}

	count := 0
	for u := 1; u <= n; u++ {
		if finalMatch[u] != -1 {
			fmt.Println(u, "-->", finalMatch[u])
			count++
		}
	}
	fmt.Println("最大组合数", count)
}

func extractMaxMatch(g *Graph) (int, []int) {
	//fmt.Println(g.match)
	value := 0
	for _, item := range g.match {
		if item != -1 {
			value++
		}
	}
	return value, g.match
}
