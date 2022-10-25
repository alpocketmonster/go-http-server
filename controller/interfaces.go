package controller

type ValidatorInt interface {
	IsValid(val string) bool
	GetRule() string
}
