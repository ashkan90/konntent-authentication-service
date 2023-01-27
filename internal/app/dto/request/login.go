package request

import (
	"konntent-authentication-service/pkg/workspaceclient/model"
	"strconv"
)

type LoginInternalRequest struct {
	UID int
}

func NewLoginInternalRequest(uid string) LoginInternalRequest {
	val, _ := strconv.Atoi(uid)
	return LoginInternalRequest{UID: val}
}

func (r LoginInternalRequest) ToModel() model.GetWorkspacesByUIDRequest {
	return model.GetWorkspacesByUIDRequest{
		UserID: r.UID,
	}
}
