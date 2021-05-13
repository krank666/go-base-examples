package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type Info struct {
	Name string
	Age int
}

type Args []string

type User struct {
	ID uint
	Args Args
}

func (a *Args)Scan(value interface{})error{
	str,ok :=  value.([]byte)
	if !ok {
		return errors.New("无法解析的数据类型")
	}
	arr := strings.Split(string(str),",")
	*a = arr
	return nil
}

func (a Args)Value()(driver.Value, error){
	if len(a)>0 {
		var str string = a[0]
		for _,v:=range a[1:] {
			str+= ","+v
		}
		return str,nil
	}
	return "",nil
}

func (i *Info)Scan(value interface{})error  {
	switch v:= value.(type) {
	case string:
		fmt.Println(v)
	case []uint8:
		json.Unmarshal(v,i)
		fmt.Println(i)
	default:
		fmt.Println(i)
		return errors.New("不能识别的类型")
	}
	return nil
}

func (i Info)Value() (driver.Value, error)  {
	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return "",err
	}
	return string(b), nil
}

func main() {

	dsn := "root:Aa@6447985@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	var u User
	err := db.First(&u).Error
	fmt.Println(err,u)

	te([]uint8{1,2,3,4,5})
}

func te(i interface{}){
	_,ok :=  i.([]byte)
	fmt.Println(ok)
}