import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	name   string
	cap    int
	used   int
}

type Pod struct {
	name string
	req  int
}

func filterNodes(nodes []Node, pod Pod) []Node {
	var ok []Node
	for _, n := range nodes {
		if n.cap-n.used >= pod.req {
			ok = append(ok, n)
		}
	}
	return ok
}

func rlSelect(nodes []Node, q map[string]float64, eps float64) Node {
	if rand.Float64() < eps {
		return nodes[rand.Intn(len(nodes))]
	}
	best := nodes[0]
	for _, n := range nodes {
		if q[n.name] > q[best.name] {
			best = n
		}
	}
	return best
}

func schedule(pods []Pod, nodes []Node, eps float64) time.Duration {
	q := make(map[string]float64)
	start := time.Now()
	for _, p := range pods {
		ok := filterNodes(nodes, p)
		if len(ok) == 0 {
			continue
		}
		chosen := rlSelect(ok, q, eps)
		for i := range nodes {
			if nodes[i].name == chosen.name {
				nodes[i].used += p.req
				q[nodes[i].name] += 0.1 // reward
			}
		}
	}
	return time.Since(start)
}

func main() {
	nodes := []Node{{"N1", 10, 0}, {"N2", 8, 0}, {"N3", 12, 0}}
	pods := []Pod{{"P1", 3}, {"P2", 4}, {"P3", 5}, {"P4", 2}}
	lat := schedule(pods, nodes, 0.2)
	fmt.Printf("RL Scheduler total latency: %v\n", lat)
}

