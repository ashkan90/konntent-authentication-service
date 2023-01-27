package authorize

import "konntent-authentication-service/internal/app/dto/aggregation"

type JWTClaim struct {
	User       JWTUserClaim        `json:"user"`
	Authority  JWTAuthorityClaim   `json:"authority"`
	Workspaces []JWTWorkspaceClaim `json:"workspaces"`
}

type JWTUserClaim struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type JWTWorkspaceClaim struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type JWTAuthorityClaim struct {
	Test string `json:"test"`
}

func NewClaimAggregation(ua *aggregation.User, wa []aggregation.Workspace) JWTClaim {
	return JWTClaim{
		User:       NewUserClaim(ua),
		Authority:  JWTAuthorityClaim{},
		Workspaces: NewWorkspaceClaim(wa),
	}
}

func NewUserClaim(ua *aggregation.User) JWTUserClaim {
	return JWTUserClaim{
		FirstName: ua.FirstName,
		LastName:  ua.LastName,
		Email:     ua.Email,
	}
}

func NewWorkspaceClaim(wa []aggregation.Workspace) []JWTWorkspaceClaim {
	var ws = make([]JWTWorkspaceClaim, 0, len(wa))

	for _, workspace := range wa {
		ws = append(ws, JWTWorkspaceClaim{
			ID:   workspace.ID,
			Name: workspace.Name,
			URL:  workspace.URL,
		})
	}

	return ws
}
