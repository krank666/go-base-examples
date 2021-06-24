package main

import (
	"errors"
	"fmt"
	"time"
)

type PersonParam string
type PhoneParam string

type Person struct {
	Name PersonParam
}

func NewPerson(name PersonParam)(Person,error){
	if time.Now().Unix()%2 == 0 {
		return Person{
			Name: name,
		},errors.New("player dead")
	}
	return Person{
		Name: name,
	},nil
}

type Phone struct {
	Name PhoneParam
}

func NewPhone(name PhoneParam)Phone{
	return Phone{
	Name: name,
	}
}

type Player struct {
	Person Person
	Phone Phone
}

func NewPlayer(person Person,phone Phone)Player{
	return Player{Person:person,Phone:phone}
}

func (player Player)play()string{
	return fmt.Sprintf("%s在玩%s",player.Person.Name,player.Phone.Name)
}

func main()  {
	 InitPlayer("wo","iphone")
}
