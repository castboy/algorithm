package main

import (
	"algorithm/stack"
	"fmt"
	"strings"
)

func match(name string) bool {
	return strings.HasSuffix(name, "m")
}

func findThom() bool {
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	mark := map[string]bool{}

	queue := stack.NewQueue()
	queue.Enqueue("you")

	for {
		e := queue.Dequeue()
		if e == nil {
			break
		}

		gh := graph[e.(string)]
		for i := range gh {
			if match(gh[i]) {
				return true
			}

			if !mark[gh[i]] {
				queue.Enqueue(gh[i])
				mark[gh[i]] = true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(findThom())
}
