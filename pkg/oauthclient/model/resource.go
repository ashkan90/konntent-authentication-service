package model

type Generic struct {
	Token         string `json:"-"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
}

//func (g Generic) ToGoogle() GoogleResource {
//	return GoogleResource{
//		ID:            g.ID,
//		Email:         g.Email,
//		VerifiedEmail: g.VerifiedEmail,
//	}
//}
