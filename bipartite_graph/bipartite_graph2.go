package main

import "fmt"

const MaxN = 10

var visted []bool
var match []int
var g [MaxN][]int

func init() {
	visted = make([]bool, MaxN)
	match = make([]int, MaxN)
}

func main() {
	var t, n, m, u, v int
	fmt.Scanf("%d", &t)
	//fmt.Println(t)
	for t > 0 {
		fmt.Scanf("%d%d", &n, &m)

		for i := 0; i < MaxN; i++ {
			match[i] = -1
			g[i] = make([]int, 0)
		}

		for i := 0; i < m; i++ {
			fmt.Scanf("%d%d", &u, &v)
			//g[u] = append(g[u], n+v)
			g[u] = append(g[u], v)
		}
		for i := 1; i <= n; i++ {
			fmt.Println("i", i, g[i])
		}

		fmt.Println("------")

		ans := 0
		for i := 1; i <= n; i++ {
			visted = make([]bool, MaxN)
			if dfs(i) {
				//fmt.Println("match", match)
				ans++
			}
		}
		fmt.Printf("%d\n", ans)
		t--
	}
}

func dfs(u int) bool {
	fmt.Println("u", u)
	for i := 0; i < len(g[u]); i++ {
		if !visted[g[u][i]] {
			visted[g[u][i]] = true
			if match[g[u][i]] == -1 || dfs(match[g[u][i]]) {
				match[g[u][i]] = u
				match[u] = g[u][i]
				fmt.Println(match)
				printMatch(3, match)
				return true
			}
		}
	}
	return false
}

func printMatch(n int, match []int) {
	for u := 1; u <= n; u++ {
		if match[u] != -1 {
			fmt.Print(u, "-->", match[u])
			fmt.Print("-->")
		}
	}
	fmt.Println()
}
