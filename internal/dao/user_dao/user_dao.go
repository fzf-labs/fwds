package user_dao

import (
	"fwds/internal/model"
	"fwds/internal/model/users_model"
	"fwds/pkg/db"

	"github.com/gin-gonic/gin"
)

var _ iUserDao = (*UserDao)(nil)

type UserDao struct {
}

type iUserDao interface {
	GetUser(ctx *gin.Context, email string) (*users_model.Users, error)
}

func (u *UserDao) GetUser(ctx *gin.Context, email string) (*users_model.Users, error) {
	ret, err := users_model.NewQueryBuilder(ctx).WhereEmail(model.EqualPredicate, email).First(db.GetReadDB())
	if err != nil {
		return nil, err
	}
	return ret, nil
}
