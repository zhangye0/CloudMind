package model

import (
	"gorm.io/gorm"
)

var (
	UserAuthTypeQq    = "qq"
	UserAuthTypeWx    = "wx"
	UserAuthTypeEmail = "email"
)
var _ UserAuthModel = (*customUserAuthModel)(nil)

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}
)

// NewUserAuthModel returns a model for the database table.
func NewUserAuthModel(conn *gorm.DB) UserAuthModel {
	return &customUserAuthModel{
		defaultUserAuthModel: newUserAuthModel(conn),
	}
}
