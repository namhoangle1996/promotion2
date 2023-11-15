package main

import "fmt"

func main() {
	//conf := config.NewConfig()
	//_ = redis.NewRedisConn(conf.Redis)
	//
	//db := postgresql.NewConn(conf.Postgresql)
	//
	////http handler
	//httpHandler.NewHttpServer(conf.HTTP, db)
	//
	//select {}

	fmt.Println("get center points ", findMinHeightTrees(4, [][]int{[]int{1, 2}, []int{1, 3}, []int{0, 1}}))
}

func isParadollineNumb() {
	var rn int

	var input = 142
	tmp := input
	for tmp > 0 {
		rn = rn*10 + tmp%10
		tmp = tmp / 10
	}
	fmt.Println("is palindrome number", isEqualNumbers(input, rn))
}

func isEqualNumbers(originNum, reverseNum int) bool {
	return originNum == reverseNum
}

func findMinHeightTrees(node int, edges [][]int) []int {
	if node == 1 {
		return []int{0}
	}

	queue := []int{}
	//graph := make(map[int][]int)
	graph := make([][]int, node)
	degree := make([]int, node)

	for _, e := range edges {
		fmt.Println("get edge ", e)
		na, nb := e[0], e[1]
		graph[na] = append(graph[na], nb)
		graph[nb] = append(graph[nb], na)
		fmt.Println(graph)
		degree[na]++
		degree[nb]++
	}

	for i := range degree {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	// topological sort
	for node > 2 {
		size := len(queue)
		node -= size

		for size > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, v := range graph[node] {
				degree[v]--
				if degree[v] == 1 {
					queue = append(queue, v)
				}
			}

			size--
		}
	}

	return queue
}
