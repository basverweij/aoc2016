package main

import "fmt"

type operation interface {
	perform(pwd *password)
}

type swapPosition struct {
	from, to int
}

func (sp *swapPosition) String() string {
	return fmt.Sprintf(
		"swap position %d with position %d",
		sp.from,
		sp.to)
}

func (sp *swapPosition) perform(pwd *password) {
	pwd.swap(sp.from, sp.to)
}

type swapLetter struct {
	fromLetter, toLetter byte
}

func (sl *swapLetter) String() string {
	return fmt.Sprintf(
		"swap letter %s with letter %s",
		string(sl.fromLetter),
		string(sl.toLetter))
}

func (sl *swapLetter) perform(pwd *password) {
	var from, to int

	for i, l := range pwd.value {
		if l == sl.fromLetter {
			from = i
		} else if l == sl.toLetter {
			to = i
		}
	}

	pwd.swap(from, to)
}

type rotate struct {
	rotateLeft bool
	steps      int
}

func (r *rotate) String() string {
	var dir string
	if r.rotateLeft {
		dir = "left"
	} else {
		dir = "right"
	}

	suffix := ""
	if r.steps != 1 {
		suffix = "s"
	}

	return fmt.Sprintf(
		"rotate %s %d step%s",
		dir,
		r.steps,
		suffix)
}

func (r *rotate) perform(pwd *password) {
	pwd.rotate(r.rotateLeft, r.steps)
}

type rotateLetter struct {
	letter byte
}

func (rl *rotateLetter) String() string {
	return fmt.Sprintf(
		"rotate based on position of letter %s",
		string(rl.letter))
}

func (rl *rotateLetter) perform(pwd *password) {
	var steps int

	for i, l := range pwd.value {
		if l == rl.letter {
			steps = i + 1
			break
		}
	}

	if steps >= 5 {
		steps++
	}

	pwd.rotate(false, steps)
}

type reverse struct {
	from, to int
}

func (r *reverse) String() string {
	return fmt.Sprintf(
		"reverse positions %d through %d",
		r.from,
		r.to)
}

func (r *reverse) perform(pwd *password) {
	pwd.reverse(r.from, r.to)
}

type move struct {
	from, to int
}

func (m *move) String() string {
	return fmt.Sprintf(
		"move position %d to position %d",
		m.from,
		m.to)
}

func (m *move) perform(pwd *password) {
	pwd.move(m.from, m.to)
}
