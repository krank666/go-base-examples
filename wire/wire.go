//+build wireinject

package main

import "github.com/google/wire"

func InitPlayer(personName PersonParam, PhoneName PhoneParam) (Player,error) {
	wire.Build(NewPlayer, NewPhone, NewPerson)
	return Player{},nil
}