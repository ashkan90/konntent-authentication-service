package aggregation

import (
	"konntent-authentication-service/internal/app/datamodel"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
}

func NewUserAggregation(dm *datamodel.User) *User {
	return &User{
		FirstName: dm.Account.FirstName,
		LastName:  dm.Account.LastName,
		Email:     dm.Email,
	}
}
