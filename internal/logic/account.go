package logic

import (
	"encoding/json"
	"fwds/internal/constants/cachekey"
	"fwds/internal/model"
	"fwds/internal/model/pay_account_api_model"
	"fwds/internal/model/pay_account_model"
	"fwds/pkg/db"
	"fwds/pkg/redis"
	"github.com/gin-gonic/gin"
)

type Account struct {
}

func NewAccount() *Account {
	return &Account{}
}

// AccountDetail 缓存结构
type AccountDetail struct {
	Key    string                 `json:"key"`    // 调用方 key
	Secret string                 `json:"secret"` // 调用方 secret
	Status int32                  `json:"status"` // 调用方启用状态 1=启用 -1=禁用
	Apis   []AccountDetailApiData `json:"apis"`   // 调用方授权的 Apis
}
type AccountDetailApiData struct {
	Method string `json:"method"` // 请求方式
	Api    string `json:"api"`    // 请求地址
}

func (s *Account) GetAccountDetailByKey(c *gin.Context, key string) (*AccountDetail, error) {
	data := new(AccountDetail)
	cacheKey := cachekey.SignatureDetail.BuildCacheKey(key)
	result, err := redis.Client.WithContext(c).Get(c, cacheKey).Result()
	switch err {
	case redis.ErrRedisNil:
		//查询数据生成
		payAccount, err := pay_account_model.NewQueryBuilder().WhereStatus(model.EqualPredicate, 1).WhereAppKey(model.EqualPredicate, key).First(db.GetDB().WithContext(c))
		if err != nil {
			return nil, err
		}
		payAccountApi, err := pay_account_api_model.NewQueryBuilder().WherePayAccountId(model.EqualPredicate, payAccount.Id).WhereStatus(model.EqualPredicate, 1).OrderById(false).QueryAll(db.GetDB().WithContext(c))
		if err != nil {
			return nil, err
		}
		data.Key = key
		data.Secret = payAccount.AppSecret
		data.Status = payAccount.Status
		data.Apis = make([]AccountDetailApiData, len(payAccountApi))
		for k, v := range payAccountApi {
			apiData := AccountDetailApiData{
				Method: v.Method,
				Api:    v.Api,
			}
			data.Apis[k] = apiData
		}
		marshal, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		redis.Client.WithContext(c).Set(c, cacheKey, string(marshal), cachekey.SignatureDetail.TTL())
		return data, nil
	case nil:
		err := json.Unmarshal([]byte(result), data)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, err
	}

}

func (s *Account) SetAccountDetailByKey(c *gin.Context, key string) (err error) {
	panic("implement me")
}
