// import fmt
package test

type GCounter struct {
	count map[string]int
}

type PNCounter struct {
	P map[string]int
	N map[string]int
}

type ORSet struct {
	add    map[string]map[string]bool
	remove map[string]map[string]bool
}

func NewORSet() *ORSet {
	return &ORSet{
		add:    make(map[string]map[string]bool),
		remove: make(map[string]map[string]bool),
	}
}

func NewCounter() *GCounter {
	return &GCounter{
		count: make(map[string]int),
	}
}

func NewPNCounter() *PNCounter {
	return &PNCounter{
		P: make(map[string]int),
		N: make(map[string]int),
	}
}

func (g *GCounter) Increment(nodeID string) {
	g.count[nodeID]++
}

func (or *ORSet) Add(nodeID string, value string) {
	if or.add[nodeID] == nil {
		or.add[nodeID] = make(map[string]bool)
	}
	or.add[nodeID][value] = true
}

func (or *ORSet) Remove(nodeID string, value string) {
	if or.remove[nodeID] == nil {
		or.remove[nodeID] = make(map[string]bool)
	}
	or.remove[nodeID][value] = true
}

func (or *ORSet) MergeOR(other *ORSet) {
	for node, ids := range other.add {
		if or.add[node] == nil {
			or.add[node] = make(map[string]bool)
		}
		for id := range ids {
			or.add[node][id] = true
		}
	}

	for node, ids := range other.remove {
		if or.remove == nil {
			or.remove[node] = make(map[string]bool)
		}
		for id := range ids {
			or.remove[node][id] = true
		}
	}

}

func (p *PNCounter) PNIncrement(nodeID string) {
	p.P[nodeID]++
}

func (p *PNCounter) PNDecrement(nodeID string) {
	p.N[nodeID]++
}

func (p *PNCounter) Value() int {
	pTotal, nTotal := 0, 0
	for _, v := range p.P {
		pTotal += v
	}

	for _, v := range p.N {
		nTotal += v
	}

	return pTotal - nTotal

}

func (p *PNCounter) MergePN(other *PNCounter) {
	for node, val := range other.P {
		if curr, ok := p.P[node]; !ok || val > curr {
			p.P[node] = val
		}
	}

	for node, val := range other.N {
		if curr, ok := p.N[node]; !ok || val > curr {
			p.N[node] = val
		}
	}

	for k, v := range other.N {
		if curr, ok := p.N[k]; !ok || v > curr {
			p.N[k] = v
		}
	}
}

func (other *ORSet) Contains(element string) bool {
	return true
}

func (g *GCounter) totalValue() int {
	total := 0
	for _, v := range g.count {
		total += v
	}

	return total
}

func (g *GCounter) Merge(nodes *GCounter) {
	for k, val := range nodes.count {
		if curr, ok := g.count[k]; !ok || val > curr {
			g.count[k] = val
		}
	}
}

// GCounter
/*
func main(){
	nodeA := NewCounter();
	nodeB := NewCounter();
	i := 0;
	nodeA.Increment("Hey");

	for i<5 {
		nodeA.Increment("How")
		i++;
	}

	for i < 8{
		nodeB.Increment("Are")
		nodeB.Increment("You")
		i++
	}

	nodeA.Increment("Fine")

	nodeA.Merge(nodeB)
	nodeB.Merge(nodeA)

	fmt.Println(nodeA.count)

	fmt.Println(nodeA.totalValue());
	fmt.Println(nodeB.totalValue());

}
*/

//PNCounter
/*
func main() {
	nodeA := NewPNCounter()
	nodeB := NewPNCounter()

	nodeA.PNDecrement("H")
	nodeA.PNDecrement("E")

	nodeB.PNDecrement("L")
	nodeB.PNIncrement("L")
	nodeB.PNDecrement("O")
	nodeB.PNDecrement("O")

	fmt.Println("Before Merge: ")
	fmt.Println(nodeA.P)
	fmt.Println(nodeA.N)
	fmt.Println(nodeB.P)
	fmt.Println(nodeB.N)
	fmt.Println("A:", nodeA.Value())
	fmt.Println("B:", nodeB.Value())
	fmt.Println(nodeA)
	fmt.Println(nodeB)

	nodeA.MergePN(nodeB)
	nodeB.MergePN(nodeA)

	fmt.Println("After Merge: ")
	fmt.Println(nodeA.P)
	fmt.Println(nodeA.N)
	fmt.Println(nodeB.P)
	fmt.Println(nodeB.N)
	fmt.Println("A:", nodeA.Value())
	fmt.Println("B:", nodeB.Value())
	fmt.Println(nodeA)
	fmt.Println(nodeB)
}
*/

// OR-Set
func main() {

}
