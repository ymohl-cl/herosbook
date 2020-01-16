package controller

import "github.com/ymohl-cl/herosbook/pkg/model"

// Controller interface to implement private manager resources
type Controller interface {
	BookCategory(l Langage) []model.Classification
	BookTheme(l Langage) []model.Classification
	UserGenre(l Langage) []model.Classification
	Langage(l Langage) []model.Classification
	NewLangage(l string) Langage
}

type controller struct{}

// NewController instance
func New(appName string) Controller {
	return &controller{}
}
