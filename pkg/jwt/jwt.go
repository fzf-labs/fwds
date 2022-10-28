package jwt

import (
	"context"
	"errors"
	"fwds/internal/conf"
	"strconv"
	"time"

	"fwds/internal/constants/cachekey"
	"fwds/pkg/redis"

	j "github.com/dgrijalva/jwt-go"
)

//获取jwt的secret 在配置文件中
type JwtConfig struct {
	JwtSecret   []byte
	JwtDuration time.Duration
}

func NewJC() *JwtConfig {
	return &JwtConfig{
		JwtSecret:   []byte(conf.Conf.Jwt.JwtSecret),
		JwtDuration: conf.Conf.Jwt.JwtDuration,
	}
}

var (
	TokenExpired       = errors.New("Token is expired")
	TokenNotValidYet   = errors.New("Token not active yet")
	TokenMalformed     = errors.New("That's not even a token")
	TokenInvalid       = errors.New("Couldn't handle this token:")
	TokenAddBlackErr   = errors.New("Add token black fail")
	TokenGetBlackErr   = errors.New("get token black fail")
	TokenCheckBlackErr = errors.New("check token black fail")
)

type Context struct {
	UUID     string
	NickName string
}

// CustomClaims
// @Description:
//
type CustomClaims struct {
	Context
	j.StandardClaims
}

// GenerateToken 生成token
func (jc *JwtConfig) GenerateToken(c Context) (tokenStr string, expiresAt int64, err error) {
	expiresAt = time.Now().Add(time.Hour * jc.JwtDuration).Unix() // 过期时间
	customClaims := CustomClaims{
		Context: Context{
			UUID:     c.UUID,
			NickName: c.NickName,
		},
		StandardClaims: j.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: expiresAt,                // 过期时间 一周
			Issuer:    conf.Conf.App.Name,
		},
	}
	token := j.NewWithClaims(j.SigningMethodHS256, customClaims)
	tokenStr, err = token.SignedString(jc.JwtSecret)
	return
}

// ParseToken 解析token
func (jc *JwtConfig) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := j.ParseWithClaims(tokenString, &CustomClaims{}, func(token *j.Token) (interface{}, error) {
		return jc.JwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*j.ValidationError); ok {
			if ve.Errors&j.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&j.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&j.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if customClaims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return customClaims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// RefreshToken 刷新token
func (jc *JwtConfig) RefreshToken(tokenString string) (string, error) {
	j.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := j.ParseWithClaims(tokenString, &CustomClaims{}, func(token *j.Token) (interface{}, error) {
		return jc.JwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if customClaims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		j.TimeFunc = time.Now
		customClaims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		claims := j.NewWithClaims(j.SigningMethodES256, customClaims)
		return claims.SignedString(jc.JwtSecret)
	}
	return "", TokenInvalid
}

// AddBlack token加入黑名单
func (jc *JwtConfig) AddBlack(tokenString string) error {
	c, err := jc.ParseToken(tokenString)
	if err != nil {
		return err
	}
	//获取redis客户端
	r := redis.Client
	//获取key
	key := cachekey.JwtBlack.BuildCacheKey(c.UUID)
	//token的过期时间-当前的时间差
	expiresAt := time.Now().Sub(time.Unix(c.ExpiresAt, 0))
	err = r.Set(context.Background(), key, c.NotBefore, expiresAt).Err()
	if err != nil {
		return TokenAddBlackErr
	}
	return nil
}

// CheckBlack 校验token是否是黑名单中的token
func (jc *JwtConfig) CheckBlack(c *CustomClaims) error {
	//获取redis客户端
	r := redis.Client
	//获取key
	key := cachekey.JwtBlack.BuildCacheKey(c.UUID)
	//获取时间
	result, err := r.Get(context.Background(), key).Result()
	if err != nil {
		if err != redis.ErrRedisNil {
			return TokenGetBlackErr
		}
	}
	resultInt64, _ := strconv.ParseInt(result, 10, 64)
	//小于签发日期的key都是失效的
	if c.NotBefore < resultInt64 {
		return TokenCheckBlackErr
	}
	return nil
}
