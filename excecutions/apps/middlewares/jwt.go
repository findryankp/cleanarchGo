package middlewares

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func jwtCreate() {
	file, err := generates.FilesCreate("./apps/middlewares/jwt.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, jwtContent())
		fmt.Println("JWT File Created")
	}
}

func jwtContent() string {
	var text = `package middlewares

import (
	"` + generates.ModuleName + `/apps/configs"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Id   int    ` + "`json:" + `"id"` + "`" + `
	Role string ` + "`json:" + `"role"` + "`" + `
	jwt.RegisteredClaims
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(configs.InitConfig().JWTKEY),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId int, role string) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 2))
	claims := &JwtCustomClaims{
		Id:   userId,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.InitConfig().JWTKEY))
	if err != nil {
		return "", err
	}

	return t, nil
}

func TrimPrefixHeaderToken(reqToken string) string {
	prefix := "Bearer "
	return strings.TrimPrefix(reqToken, prefix)
}

func ValidateToken(c echo.Context) error {
	reqToken := c.Request().Header.Get("Authorization")
	tokenString := TrimPrefixHeaderToken(reqToken)

	if !strings.HasPrefix(reqToken, "Bearer") {
		return errors.New("request does not contain a valid token")
	}

	if tokenString == "" {
		return errors.New("request does not contain a valid token")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.InitConfig().JWTKEY), nil
		},
	)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*JwtCustomClaims); !ok {
		return errors.New("couldn't parse claims")
	}

	return nil
}

func ClaimsToken(c echo.Context) JwtCustomClaims {
	reqToken := c.Request().Header.Get("Authorization")
	tokenString := TrimPrefixHeaderToken(reqToken)
	claims := &JwtCustomClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})

	return *claims
}
	
	`

	return text
}
