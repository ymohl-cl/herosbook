package users

// Parameters users
const (
	PseudoSizeMin = 4
	PseudoSizeMax = 56
	AgeMin        = 10
	AgeMax        = 142
)

// User : _
type User struct {
	ID    PublicID     `json:"public_id" valid:"-"`
	Pass  Passwords    `json:"passwords" valid:"required"`
	Infos Informations `json:"informations" valid:"required"`
}

// PublicID : _
type PublicID struct {
	Value string `json:"id" valid:"required"`
}

// Passwords : _
type Passwords struct {
	One string `json:"password_1" valid:"printableascii"`
	Two string `json:"password_2" valid:"printableascii"`
	Old string `json:"oldPassword" valid:"-"`
}

// Informations : _
type Informations struct {
	Pseudo string `json:"pseudo" valid:"alphanum"`
	Age    uint8  `json:"age" valid:"required"`
	Sex    string `json:"sex" valid:"alpha, in(male|female)"`
	Email  string `json:"email" valid:"email"`
}

// ClearPasswords reset the Passwords structure
func (u *User) ClearPasswords() {
	u.Pass = Passwords{}
}
