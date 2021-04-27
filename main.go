package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/punyaikhwan/echo-rsa-jwt"
	mw "github.com/punyaikhwan/echo-rsa-jwt/middleware"
)

func tesJWT(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, echo_rsa_jwt.DecodeJWT(c.Get("user").(*jwt.Token)))
}

func genJWT(c echo.Context) (err error) {
	claims := map[string]interface{}{
		"name": c.FormValue("name"),
		"nim":  c.FormValue("nim"),
	}
	token, err := echo_rsa_jwt.GenerateJWT(echo_rsa_jwt.GenerateJWTInput{
		PrivateKeyPath: "private.pem",
		Claims:         claims,
		MinuteToExpire: 30,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
func main() {
	e := echo.New()
	e.GET("/tesjwt", tesJWT, mw.JWTEchoRSA("public.pem"))
	e.POST("/genjwt", genJWT)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":9000"))
}
