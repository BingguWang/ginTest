package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
	token

	json web token 就是JWT，是种身份验证的方式

	JSON Web令牌（JWT）作为令牌系统而不是在每次请求时都发送用户名和密码，因此比其他方法（如基本身份验证）具有固有的优势。

	主要有两个部分：
		提供用户名和密码以获取令牌；
		并根据请求检查该令牌。

	需要将APP_KEY常量更改为机密（理想情况下，该常量将存储在代码库外部），
	并改进用户名/密码检查中的内容，TokenHandler以检查不仅仅是myusername/ mypassword组合。

*/
var jwtkey = []byte("wangbing") //自定义一个字符串
var str string                  //保存token

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func main() {
	r := gin.Default()

	r.GET("/set", setToken)
	r.GET("/get", getToken)

	r.Run(":8088")
}

func setToken(c *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour) //token过期时间
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  //签名颁发机构
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //需要传入Claims
	fmt.Printf("token:\n%v\n", token)
	tokenString, err := token.SignedString(jwtkey) //生成token
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tokenString:\n", tokenString)
	str = tokenString
	c.JSON(200, gin.H{"token": str})
}

func getToken(c *gin.Context) { //获取请求携带的token，也就是生成token的时候产生的tokenString，然后解析出Claims和token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "没有权限"})
		c.Abort()
		return
	}
	token, Claims, err := ParseToken(tokenString)
	fmt.Printf("token:\n%v\n", token)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return
	}
	fmt.Println("Claims:\n", Claims)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) { //解析token
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	}) //ParseWithClaims解析出token中的Claims和token
	return token, Claims, err
}
