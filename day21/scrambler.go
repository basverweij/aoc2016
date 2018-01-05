package main

func scrambler(value string, ops []operation) string {
	pwd := newPassword(value)

	for _, op := range ops {
		op.perform(pwd)
	}

	return pwd.String()
}
