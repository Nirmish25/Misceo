package rga

type Element struct {
	ID      string
	Value   string
	PrevID  string
	Deleted bool
}

type RGA struct {
	Elements map[string]*Element
}
