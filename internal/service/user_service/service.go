package admin_service

import (
	"github.com/gin-gonic/gin"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Create(ctx *gin.Context, adminData string) (id int32, err error)
	Update(ctx *gin.Context, searchData string) (listData []string, err error)
}

type service struct{}

func (s *service) i() {}

func (s *service) Create(ctx *gin.Context, adminData string) (id int32, err error) {
	panic("implement me")
}

func (s *service) Update(ctx *gin.Context, searchData string) (listData []string, err error) {
	panic("implement me")
}
