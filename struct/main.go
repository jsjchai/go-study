package main

import (
	"fmt"
)

type User struct {
	id       uint32
	username string
	age      uint8
	address  string
}

func toString(user User) string {
	return fmt.Sprintf("id=%d,username=%s,age=%d,address=%s", user.id, user.username, user.age, user.address)
}

func printf(user *User) {
	fmt.Println(fmt.Sprintf("id=%d,username=%s,age=%d,address=%s", user.id, user.username, user.age, user.address))
}

// 定义矩形
type Rect struct {
	width, heigth float32
}

//面积
func (r *Rect) area() float32 {
	return r.width * r.heigth
}

//周长
func (r *Rect) perimeter() float32 {
	return (r.width + r.heigth) * 2
}

func main() {
	var user User
	user = User{10000, "scott", 18, "HZ,China"}
	fmt.Println(toString(user))
	var user1 User = User{id: 10000, username: "Tom", age: 18, address: "BeiJing,China"}
	fmt.Println(toString(user1))
	user.id = 10001

	printf(&user)
	fmt.Println(user1)

	fmt.Println("-----------------------------------")
	var rect Rect = Rect{1.2, 3.2}
	fmt.Println("面积:", rect.area())
	fmt.Println("周长:", rect.perimeter())
}
