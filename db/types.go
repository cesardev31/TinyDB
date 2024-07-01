package db

type DataType int

const (
	String DataType = iota
	Integer
	Float
	Boolean
)

type Column struct {
	Name string
	Type DataType
}
