package rga

type Element struct {
	ID      string
	Value   string
	PrevID  string
	Deleted bool
}

type RGA struct {
	elements map[string]*Element
}

func NewRGA() *RGA {
	return &RGA{
		elements: make(map[string]*Element),
	}
}

func (r *RGA) Insert(id, value, prevID string) {
	r.elements[id] = &Element{
		ID:     id,
		Value:  value,
		PrevID: prevID,
	}
}

func (r *RGA) Delete(id string) {
	if ele, ok := r.elements[id]; ok {
		ele.Deleted = true
	}
}

func (r *RGA) Merge(other *RGA) {
	for id, ele := range other.elements {
		if exists, ok := r.elements[id]; !ok {
			r.elements[id] = ele
		} else if exists.Deleted {
			r.elements[id].Deleted = true
		}
	}
}
