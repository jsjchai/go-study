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

func main() {
	var user User
	user = User{10000, "scott", 18, "HZ,China"}
	fmt.Println(toString(user))
	var user1 User = User{id: 10000, username: "Tom", age: 18, address: "beijing,China"}
	fmt.Println(toString(user1))
	user.id = 10001

	printf(&user)
}
