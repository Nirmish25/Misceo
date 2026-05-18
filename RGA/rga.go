package rga

import (
	"sort"
)

func NewRGA() *RGA {
	r := &RGA{
		Elements: make(map[string]*Element),
	}

	r.Elements["ROOT"] = &Element{
		ID: "ROOT",
	}

	return r
}

func (r *RGA) Insert(id, value, prevID string) {
	r.Elements[id] = &Element{
		ID:     id,
		Value:  value,
		PrevID: prevID,
	}
}

func (r *RGA) Delete(id string) {
	if ele, ok := r.Elements[id]; ok {
		ele.Deleted = true
	}
}

func (r *RGA) Merge(other *RGA) {
	for id, ele := range other.Elements {
		if exists, ok := r.Elements[id]; !ok {
			r.Elements[id] = ele
		} else if exists.Deleted {
			r.Elements[id].Deleted = true
		}
	}
}

func (r *RGA) ToString() string {
	children := map[string][]*Element{}

	for _, elem := range r.Elements {
		if elem.ID == "ROOT" {
			continue
		}

		children[elem.PrevID] = append(
			children[elem.PrevID],
			elem,
		)
	}

	for _, elems := range children {
		sort.Slice(elems, func(i, j int) bool {
			return elems[i].ID < elems[j].ID
		})
	}

	result := ""

	var walk func(string)

	walk = func(id string) {
		for _, child := range children[id] {
			if !child.Deleted {
				result += child.Value
			}

			walk(child.ID)
		}
	}

	walk("ROOT")

	return result
}
