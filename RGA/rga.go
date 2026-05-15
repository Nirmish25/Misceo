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
	r := &RGA{
		elements: make(map[string]*Element),
	}
	
	r.elements["ROOT"] = &Element{
		ID: "ROOT"
	}

	return r
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


func (r *RGA) ToString() string {
	child := make(map[string][]*Element)

	for _, ele := range r.elements{
		if ele.ID == "ROOT"{
			continue
		}

		child[ele.PrevID] = append()


	return string
	}

}


