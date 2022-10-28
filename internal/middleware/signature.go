package middleware

import (
	"bytes"
	"encoding/json"
	"fwds/internal/enum/enum_status"
	"fwds/internal/errno"
	"fwds/internal/logic"
	"fwds/internal/response"
	"fwds/pkg/debug"
	"fwds/pkg/signature"
	"fwds/pkg/urltable"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
	"time"
)

const (
	ttl = time.Minute * 5 // 签名超时时间 5 分钟
)

func Signature() gin.HandlerFunc {
	return func(c *gin.Context) {
		signatureStr := c.GetHeader("signature")
		if signatureStr == "" {
			response.Json(c, errno.ErrSignatureParam, nil)
			c.Abort()
			return
		}
		signatures := strings.Split(signatureStr, " ")
		if len(signatures) != 3 {
			response.Json(c, errno.ErrSignatureParam, nil)
			c.Abort()
			return
		}
		appKey := signatures[0]
		date := signatures[1]
		sign := signatures[2]
		//参数校验
		if appKey == "" || sign == "" || date == "" {
			response.Json(c, errno.ErrSignatureParam, nil)
			c.Abort()
			return
		}
		//校验参数
		params, err := CheckAndGetRequest(c)
		if err != nil {
			response.Json(c, errno.ErrNotJsonRequest, nil)
			c.Abort()
			return
		}
		//校验key
		accountDetail, err := logic.NewAccount().GetAccountDetailByKey(c, appKey)
		if err != nil {
			response.Json(c, errno.ErrAuthorization, nil)
			c.Abort()
			return
		}
		//校验状态
		if accountDetail.Status != int32(enum_status.Normal) {
			response.Json(c, errno.ErrAuthorization.WithCustomMsg(appKey+"已被禁止调用"), nil)
			c.Abort()
			return
		}
		//接口授权校验
		if len(accountDetail.Apis) < 1 {
			response.Json(c, errno.ErrAuthorization.WithCustomMsg(appKey+"未进行接口授权"), nil)
			c.Abort()
			return
		}
		//接口校验
		table := urltable.NewTable()
		for _, v := range accountDetail.Apis {
			_ = table.Append(v.Method + v.Api)
		}

		if pattern, _ := table.Mapping(c.Request.Method + c.Request.URL.Path); pattern == "" {
			response.Json(c, errno.ErrAuthorization.WithCustomMsg(appKey+"无接口权限"), nil)
			c.Abort()
			return
		}
		err = signature.New(accountDetail.Key, accountDetail.Secret, ttl).Verify(c.Request.URL.Path, c.Request.Method, params, sign, date)
		if err != nil {
			debug.PrintErr(c, err)
			response.Json(c, errno.ErrAuthorization.WithCustomMsg("签名校验错误").WithErr(err), nil)
			c.Abort()
			return
		}
	}
}

func CheckAndGetRequest(c *gin.Context) (json.RawMessage, error) {
	if c.ContentType() != binding.MIMEJSON {
		return nil, errors.New("Content-Type must be application/json")
	}
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	if data != nil {
		if !json.Valid(data) {
			return nil, errors.New("request param must be json")
		}
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 重新设置
	return data, nil
}
