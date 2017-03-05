package pullman

import "fmt"

type tree struct {
	v int
	l *tree
	r *tree
}

func(t *tree) Sum() int {
	if t == nil {
		return 0
	}
	return t.v + t.l.Sum() + t.r.Sum()
}

func(t *tree) String() string {
	if t == nil {
		return ""
	}
	return fmt.Sprint(t.l, t.v, t.r)
}

// nil receivers are useful: Find
func (t *tree) Find(v int) bool {
	if t == nil {
		return false
	}
	return t.v == v || t.l.Find(v) || t.r.Find(v)
}
