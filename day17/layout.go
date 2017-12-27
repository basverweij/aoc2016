package main

// up, down, left, right
var layout = [][][]bool{
	// y == 0
	{
		// x==0
		{false, true, false, true},
		// x==1
		{false, true, true, true},
		// x==2
		{false, true, true, true},
		// x==3
		{false, true, true, false},
	},
	// y == 1
	{
		// x==0
		{true, true, false, true},
		// x==1
		{true, true, true, true},
		// x==2
		{true, true, true, true},
		// x==3
		{true, true, true, false},
	},
	// y == 2
	{
		// x==0
		{true, true, false, true},
		// x==1
		{true, true, true, true},
		// x==2
		{true, true, true, true},
		// x==3
		{true, true, true, false},
	},
	// y == 3
	{
		// x==0
		{true, false, false, true},
		// x==1
		{true, false, true, true},
		// x==2
		{true, false, true, true},
		// x==3
		{true, false, true, false},
	},
}
