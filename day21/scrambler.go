package main

func scrambler(value string, ops []operation) string {
	pwd := newPassword(value)

	for _, op := range ops {
		op.perform(pwd)
	}

	return pwd.String()
}

func unscrambler(value string, ops []operation) string {
	pwd := newPassword(value)

	for i := len(ops) - 1; i >= 0; i-- {
		ops[i].unperform(pwd)
	}

	return pwd.String()
}
