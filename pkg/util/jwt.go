package util

import (
	"gin-pro/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(usernane, password string) (string, error) {
	nowTime := time.Now()
	//过期时间
	expireTime := nowTime.Add(2 * time.Hour)
	claims := Claims{
		usernane,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	//声明令牌
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//转成签名字符串
	token, err := tokenClaims.SigningString()

	return token, err
}

//解析令牌
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

/*
NewWithClaims(method SigningMethod, claims Claims)，method对应着SigningMethodHMAC struct{}，其包含SigningMethodHS256、SigningMethodHS384、SigningMethodHS512三种crypto.Hash方案
func (t *Token) SignedString(key interface{}) 该方法内部生成签名字符串，再用于获取完整、已签名的token
func (p *Parser) ParseWithClaims 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
func (m MapClaims) Valid() 验证基于时间的声明exp, iat, nbf，注意如果没有任何声明在令牌中，仍然会被认为是有效的。并且对于时区偏差没有计算方法
*/
