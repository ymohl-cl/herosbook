package controller

import "github.com/ymohl-cl/herosbook/pkg/model"

// FR infos langue
const (
	// langage french - 0000
	frLangID      = "lfr-0000"
	frLangPrivate = "french"
	frLangInFR    = "Français"
	frLangInEN    = "French"
)

// EN infos langue
const (
	// langage english - 0000
	enLangID      = "len-0000"
	enLangPrivate = "english"
	enLangInFR    = "Anglais"
	enLangInEN    = "English"
)

const defaultLangage Langage = enLangPrivate

// Langage type supported by worker
type Langage string

func (c controller) NewLangage(l string) Langage {
	switch l {
	case frLangPrivate:
		return frLangPrivate
	case enLangPrivate:
		return enLangPrivate
	default:
		return defaultLangage
	}
}

func (c controller) Langage(l Langage) []model.Classification {
	var list []model.Classification

	switch l {
	case frLangPrivate:
		list = append(list, model.Classification{
			Identifier:  frLangID,
			PrivateName: frLangPrivate,
			PublicName:  frLangInFR,
		})
		list = append(list, model.Classification{
			Identifier:  enLangID,
			PrivateName: enLangPrivate,
			PublicName:  enLangInFR,
		})
	case enLangPrivate:
		list = append(list, model.Classification{
			Identifier:  frLangID,
			PrivateName: frLangPrivate,
			PublicName:  frLangInEN,
		})
		list = append(list, model.Classification{
			Identifier:  enLangID,
			PrivateName: enLangPrivate,
			PublicName:  enLangInEN,
		})
	}
	return list
}
