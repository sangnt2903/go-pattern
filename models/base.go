package models

type IModel interface {
	Setter(newKey uint)
	Getter() uint
	TableName() string
}
