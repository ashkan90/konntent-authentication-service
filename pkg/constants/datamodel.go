package constants

const (
	// GenericIDCol is used to name id column on [Any] model w/out any relation call
	GenericIDCol = "id"
)

const (
	// UserEmailCol is used to name email column on User model w/out any relation call
	UserEmailCol = "email"
)

const (
	// RelAccountFNameCol is used to name first_name column on Account model with relation call
	RelAccountFNameCol = "Account.first_name"
	// RelAccountLNameCol is used to name first_name column on Account model with relation call
	RelAccountLNameCol = "Account.last_name"
)
