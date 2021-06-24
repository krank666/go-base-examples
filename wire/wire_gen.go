// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitPlayer(personName PersonParam, PhoneName PhoneParam) (Player, error) {
	person, err := NewPerson(personName)
	if err != nil {
		return Player{}, err
	}
	phone := NewPhone(PhoneName)
	player := NewPlayer(person, phone)
	return player, nil
}
