package main

type pathQueue struct {
	head *pathItem
	tail *pathItem
}

type pathItem struct {
	path *path
	next *pathItem
}

func (p *pathQueue) push(path *path) {
	i := &pathItem{path: path}

	if p.head == nil {
		p.head = i
	} else if p.tail != nil {
		p.tail.next = i
	}

	p.tail = i
}

func (p *pathQueue) pop() *path {
	if p.head == nil {
		return nil
	}

	path := p.head.path
	p.head = p.head.next

	return path
}

func (p *pathQueue) peek() *path {
	if p.head == nil {
		return nil
	}

	return p.head.path
}
