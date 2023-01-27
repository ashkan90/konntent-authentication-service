package request

import (
	"konntent-authentication-service/internal/app/datamodel"
	workspaceModel "konntent-authentication-service/pkg/workspaceclient/model"
)

type Register struct {
	User      User      `json:"user"`
	Workspace Workspace `json:"workspace"`
}

type User struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	JobDescription string `json:"jobDescription"`
}

type Workspace struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	BM          string `json:"businessModel"`
	CompanyUnit int    `json:"companyUnit"`
}

func (u User) ToUserModel() *datamodel.User {
	return &datamodel.User{
		Email:    u.Email,
		Password: u.Password,
		Account:  u.ToUserAccountModel(),
	}
}

func (u User) ToUserAccountModel() *datamodel.UserAccount {
	return &datamodel.UserAccount{
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		JobDescription: u.JobDescription,
	}
}
func (w Workspace) ToModel(uid int) workspaceModel.AddWorkspaceRequest {
	return workspaceModel.AddWorkspaceRequest{
		Name:        w.Name,
		URL:         w.URL,
		BM:          w.BM,
		CompanyUnit: w.CompanyUnit,
		UserID:      uid,
	}
}
