package basic_fields

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 2  //jwt过期时间
)

var JwtSecret = []byte("我也不知道这里用啥好，就这样吧，哎")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
