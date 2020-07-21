package main

import "fmt"

type nodes []string

func findShortestPath(network map[string]nodes, start, end string) []string {
	if _, ok := network[start]; !ok {
		fmt.Println("Path Not found or doesn't exist")
		return []string{}
	}

	if start == end {
		return []string{start}
	}

	cache := map[string]bool{}
	queue := []string{start}
	route := map[string]string{}
	fOL := 0

	for len(queue) > fOL {
		node := queue[fOL]
		fOL++

		if node == end {
			break
		}

		for _, name := range network[node] {
			if _, ok := cache[name]; !ok {
				cache[name] = true
				queue = append(queue, name)
				route[name] = node
			}
		}
	}

	if _, ok := route[end]; !ok {
		fmt.Println("Path Not found or doesn't exist")
		return []string{}
	}

	curr := end
	path := []string{curr}
	for i := 0; i < len(cache); i++ {
		if curr != start {
			curr = route[curr]
			path = append(path, curr)
		} else {
			break
		}
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func main() {
	network := map[string]nodes{
		"Min":     nodes{"William", "Jayden", "Omar"},
		"William": nodes{"Min", "Noam"},
		"Jayden":  nodes{"Min", "Amelia", "Ren", "Noam"},
		"Ren":     nodes{"Jayden", "Omar"},
		"Amelia":  nodes{"Jayden", "Adam", "Miguel"},
		"Adam":    nodes{"Amelia", "Miguel", "Sofia", "Lucas"},
		"Miguel":  nodes{"Amelia", "Adam", "Liam", "Nathan"},
		"Noam":    nodes{"Nathan", "Jayden", "William"},
		"Omar":    nodes{"Ren", "Min", "Scott"},
	}

	fmt.Println(findShortestPath(network, "Jayden", "Jayden"))
}
