package code

type Code int64

//
const (
	OK Code = iota + 200
)

// Errors

const (
	Internal Code = iota + 600
	Invalid
	InvalidRevenueRate
	InvalidRelation
)
