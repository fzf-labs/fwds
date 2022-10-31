package user_repo

import (
	"fwds/internal/dal/model"
	"fwds/internal/dal/query"
	"fwds/pkg/db"

	"github.com/gin-gonic/gin"
)

var _ iUserDao = (*UserRepo)(nil)

type UserRepo struct {
}

type iUserDao interface {
	GetUser(ctx *gin.Context, email string) (*model.User, error)
}

func (u *UserRepo) GetUser(ctx *gin.Context, email string) (*model.User, error) {
	query := query.Use(db.GetWriteDB())
	ret, err := query.User.WithContext(ctx).Where(query.User.Name.Eq(email)).First()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
