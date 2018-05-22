package model

type Model interface {
	Validate() (err error)
}
