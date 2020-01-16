package model

// Classification json model to describe the nomenclature like a dicotionnary multilanguage
type Classification struct {
	Identifier  string `json:"identifier"`
	PrivateName string `json:"privateName"`
	PublicName  string `json:"publicName"`
}

// Nomenclature json model to describe a section of nomenclature
type Nomenclature struct {
	Number int              `json:"number"`
	List   []Classification `json:"list"`
}

// NomenclatureOutput json model to output on all nomenclature
type NomenclatureOutput struct {
	User         Nomenclature `json:"user"`
	BookCategory Nomenclature `json:"bookCategory"`
	BookTheme    Nomenclature `json:"bookTheme"`
	Langage      Nomenclature `json:"langage"`
}

// NomenclatureBookOutput json model to output on books
type NomenclatureBookOutput struct {
	BookCategory Nomenclature `json:"bookCategory"`
	BookTheme    Nomenclature `json:"bookTheme"`
}
