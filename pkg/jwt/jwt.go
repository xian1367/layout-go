// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt/v5"
	"github.com/golang-module/carbon/v2"
	"github.com/xian137/layout-go/config"
	"github.com/xian137/layout-go/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
)

// JWT 定义一个jwt对象
type JWT struct {

	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

// CustomClaims 自定义载荷
type CustomClaims struct {
	UserID       string `json:"user_id"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtPkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.Get().Jwt.Key),
		MaxRefresh: time.Duration(config.Get().Jwt.MaxRefreshTime) * time.Minute,
	}
}

// ParserToken 解析 Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*CustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 2. 解析出错
	if err != nil {
		if err == jwtPkg.ErrTokenMalformed {
			return nil, ErrTokenMalformed
		}
		if err == jwtPkg.ErrTokenExpired {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 CustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	carbon.SetTimezone(carbon.PRC)
	// 1. 从 Header 里获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		if err == jwtPkg.ErrTokenExpired {
			return "", ErrTokenExpired
		}
	}

	// 4. 解析 CustomClaims 的数据
	claims := token.Claims.(*CustomClaims)

	// 5. 检查是否过了『最大允许刷新的时间』
	x := time.Now().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Unix() > x {
		// 修改过期时间
		claims.RegisteredClaims.ExpiresAt = jwtPkg.NewNumericDate(time.Unix(jwt.expireAtTime(), 0))
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 生成  Token，在登录成功时调用
func (jwt *JWT) IssueToken(userID string) map[string]string {
	// 1. 构造用户 claims 信息(负荷)
	expireAtTime := jwt.expireAtTime()
	claims := CustomClaims{
		userID,
		expireAtTime,
		jwtPkg.RegisteredClaims{
			NotBefore: jwtPkg.NewNumericDate(time.Unix(time.Now().Unix(), 0)), // 签名生效时间
			IssuedAt:  jwtPkg.NewNumericDate(time.Unix(time.Now().Unix(), 0)), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtPkg.NewNumericDate(time.Unix(expireAtTime, 0)),      // 签名过期时间
			Issuer:    config.Get().App.Name,                                  // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.ErrorIf(err)
		return map[string]string{}
	}

	return map[string]string{
		"token":  token,
		"expire": carbon.CreateFromTimestamp(expireAtTime).ToDateTimeString(),
	}
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims CustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
	var expireTime int
	if config.Get().App.Debug {
		expireTime = config.Get().Jwt.DebugExpireTime
	} else {
		expireTime = config.Get().Jwt.ExpireTime
	}

	expire := time.Duration(expireTime) * time.Minute
	return time.Now().Add(expire).Unix()
}

// parseTokenString 使用 jwtPkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtPkg.Token, error) {
	return jwtPkg.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// getTokenFromHeader 使用 jwtPkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
