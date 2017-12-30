package main

type rule struct {
	from, to uint32
}

type byFrom []rule

func (r byFrom) Len() int           { return len(r) }
func (r byFrom) Less(i, j int) bool { return r[i].from < r[j].from }
func (r byFrom) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
