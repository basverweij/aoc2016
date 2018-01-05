package main

type password struct {
	value []byte
}

func newPassword(value string) *password {
	return &password{[]byte(value)}
}

func (p *password) String() string {
	return string(p.value)
}

func (p *password) swap(x, y int) {
	p.value[x], p.value[y] = p.value[y], p.value[x]
}

func (p *password) rotate(rotateLeft bool, steps int) {
	steps %= len(p.value)

	if rotateLeft {
		p.value = append(p.value[steps:], p.value[0:steps]...)
	} else {
		p.value = append(p.value[len(p.value)-steps:], p.value[0:len(p.value)-steps]...)
	}
}

func (p *password) reverse(x, y int) {
	for i, j := x, y; i < j; i, j = i+1, j-1 {
		p.value[i], p.value[j] = p.value[j], p.value[i]
	}
}

func (p *password) move(x, y int) {
	l := p.value[x]

	if x < y {
		copy(p.value[x:y], p.value[x+1:y+1])
	} else {
		copy(p.value[y+1:x+1], p.value[y:x])
	}

	p.value[y] = l
}
