package dto

type DTO interface {
	ToEntity()
	TableName() string
}
