package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	MyType string `json:"myType"`
	jwt.StandardClaims `json:"standardClaims"`
}

func main ()  {
	mySigningKey := []byte("isKey")

	// Create the Claims

	//claims := MyClaims{
	//	MyType: "qm",
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt:time.Now().Unix() +  60 * 60 * 24 * 7,
	//		NotBefore: time.Now().Unix() - 1000,
	//		Issuer:    "qmPlus",
	//	},
	//}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"myType": "myType",
		"nbf": time.Now().Unix() - 1000,
		"exp": time.Now().Unix() + 1000,
	})
	ss,err := token.SignedString(mySigningKey)
	fmt.Println(ss)
	fmt.Println(err)
	time.Sleep(time.Second * 15)
  p,erro := jwt.ParseWithClaims(ss,&jwt.MapClaims{},func(t *jwt.Token) (i interface{},err error){
		return mySigningKey,err
	})
  fmt.Println(p.Claims,erro)
}