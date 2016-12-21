package main

var input *model

func init() {
	input = newModel()

	input.AddDevice("PrG", 1)
	input.AddDevice("PrM", 1)

	input.AddDevice("CoG", 2)
	input.AddDevice("CoM", 3)

	input.AddDevice("CuG", 2)
	input.AddDevice("CuM", 3)

	input.AddDevice("RuG", 2)
	input.AddDevice("RuM", 3)

	input.AddDevice("PlG", 2)
	input.AddDevice("PlM", 3)

	input.Elevator = 1
}
