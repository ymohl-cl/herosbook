package controller

import "github.com/ymohl-cl/herosbook/pkg/model"

// boy user genre
const (
	uBoyID      = "uB-0000"
	uBoyPrivate = "boy"
	uBoyInFR    = "Garçon"
	uBoyInEN    = "Boy"
)

// girl user genre
const (
	uGirlID      = "uG-0000"
	uGirlPrivate = "girl"
	uGirlInFR    = "Fille"
	uGirlInEN    = "Girl"
)

func (c controller) UserGenre(l Langage) []model.Classification {
	var list []model.Classification

	switch l {
	case frLangPrivate:
		list = append(list, model.Classification{
			Identifier:  uBoyID,
			PrivateName: uBoyPrivate,
			PublicName:  uBoyInFR,
		})
		list = append(list, model.Classification{
			Identifier:  uGirlID,
			PrivateName: uGirlPrivate,
			PublicName:  uGirlInFR,
		})
	case enLangPrivate:
		list = append(list, model.Classification{
			Identifier:  uBoyID,
			PrivateName: uBoyPrivate,
			PublicName:  uBoyInEN,
		})
		list = append(list, model.Classification{
			Identifier:  uGirlID,
			PrivateName: uGirlPrivate,
			PublicName:  uGirlInEN,
		})
	}
	return list
}
