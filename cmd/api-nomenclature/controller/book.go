package controller

import "github.com/ymohl-cl/herosbook/pkg/model"

// // Theme list to books
// t (theme) - SF (science-fiction) - infos Science-fiction
const (
	tSFID      = "tsf-0000"
	tSFPrivate = "fiction"
	tSFInFR    = "Science-fiction"
	tSFInEN    = "Fiction"
)

// t (theme) - FT (Fantasy) - infos Fantasy
const (
	tFTID      = "tft-0000"
	tFTPrivate = "fantasy"
	tFTInFR    = "Fantaisie"
	tFTInEN    = "Fatasy"
)

// // Category list to book
// person info
const (
	cPersonID      = "cPerson-0000"
	cPersonPrivate = "person"
	cPersonInFR    = "Personnage"
	cPersonInEN    = "Person"
)

// location info
const (
	cLocationID      = "cLocation-0000"
	cLocationPrivate = "location"
	cLocationInFR    = "Lieux"
	cLocationInEN    = "Location"
)

// tags ingo
const (
	cTagID      = "cTag-0000"
	cTagPrivate = "tag"
	cTagInFR    = "Étiquette"
	cTagInEN    = "Tag"
)

func (c controller) BookTheme(l Langage) []model.Classification {
	var list []model.Classification

	switch l {
	case frLangPrivate:
		list = append(list, model.Classification{
			Identifier:  tSFID,
			PrivateName: tSFPrivate,
			PublicName:  tSFInFR,
		})
		list = append(list, model.Classification{
			Identifier:  tFTID,
			PrivateName: tFTPrivate,
			PublicName:  tFTInFR,
		})
	case enLangPrivate:
		list = append(list, model.Classification{
			Identifier:  tSFID,
			PrivateName: tSFPrivate,
			PublicName:  tSFInEN,
		})
		list = append(list, model.Classification{
			Identifier:  tFTID,
			PrivateName: tFTPrivate,
			PublicName:  tFTInEN,
		})
	}
	return list
}

func (c controller) BookCategory(l Langage) []model.Classification {
	var list []model.Classification

	switch l {
	case frLangPrivate:
		list = append(list, model.Classification{
			Identifier:  cPersonID,
			PrivateName: cPersonPrivate,
			PublicName:  cPersonInFR,
		})
		list = append(list, model.Classification{
			Identifier:  cLocationID,
			PrivateName: cLocationPrivate,
			PublicName:  cLocationInFR,
		})
		list = append(list, model.Classification{
			Identifier:  cTagID,
			PrivateName: cTagPrivate,
			PublicName:  cTagInFR,
		})
	case enLangPrivate:
		list = append(list, model.Classification{
			Identifier:  cPersonID,
			PrivateName: cPersonPrivate,
			PublicName:  cPersonInEN,
		})
		list = append(list, model.Classification{
			Identifier:  cLocationID,
			PrivateName: cLocationPrivate,
			PublicName:  cLocationInEN,
		})
		list = append(list, model.Classification{
			Identifier:  cTagID,
			PrivateName: cTagPrivate,
			PublicName:  cTagInEN,
		})
	}
	return list
}
