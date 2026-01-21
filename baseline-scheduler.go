package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Node struct{ cpu, mem int }
type Pod struct{ cpu, mem int }

func createNodes(n int) []*Node {
	nodes := make([]*Node, n)
	for i := range nodes {
		nodes[i] = &Node{cpu: 200 + rand.Intn(200), mem: 512 + rand.Intn(512)}
	}
	return nodes
}

func newPod() Pod {
	return Pod{cpu: 10 + rand.Intn(50), mem: 32 + rand.Intn(256)}
}

func filter(nodes []*Node, p Pod) []*Node {
	time.Sleep(time.Millisecond * time.Duration(1+rand.Intn(3)))
	var c []*Node
	for _, n := range nodes {
		if n.cpu >= p.cpu && n.mem >= p.mem {
			c = append(c, n)
		}
	}
	return c
}

func score(c []*Node) *Node {
	time.Sleep(time.Millisecond * time.Duration(2+rand.Intn(5)))
	if len(c) == 0 {
		return nil
	}
	best := c[0]
	for _, n := range c[1:] {
		if n.cpu+n.mem > best.cpu+best.mem {
			best = n
		}
	}
	return best
}

func bind(n *Node, p Pod) {
	time.Sleep(time.Millisecond * time.Duration(1+rand.Intn(2)))
	if n != nil {
		n.cpu -= p.cpu
		n.mem -= p.mem
	}
}

func schedule(p Pod, nodes []*Node) float64 {
	start := time.Now()
	c := filter(nodes, p)
	n := score(c)
	bind(n, p)
	return float64(time.Since(start).Milliseconds())
}

func percentile(s []float64, p float64) float64 {
	r := p / 100 * float64(len(s)-1)
	l, u := int(r), int(math.Ceil(r))
	if l == u {
		return s[l]
	}
	return s[l]*(1-(r-float64(l))) + s[u]*(r-float64(l))
}

func run(nodesCount, pods int) {
	nodes := createNodes(nodesCount)
	lat := make([]float64, pods)
	for i := 0; i < pods; i++ {
		lat[i] = schedule(newPod(), nodes)
	}
	sort.Float64s(lat)
	var sum float64
	for _, v := range lat {
		sum += v
	}
	fmt.Printf("Nodes:%d Pods:%d Avg:%.2f P50:%.2f P95:%.2f\n",
		nodesCount, pods, sum/float64(pods), percentile(lat, 50), percentile(lat, 95))
}

func main() {
	for _, n := range []int{3, 5, 7, 9, 11} {
		run(n, 500)
	}
}
