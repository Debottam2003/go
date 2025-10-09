package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int8
}

type Profile struct {
	UserData User // Name Field (Normal),
	// User          // Embading (Like the Inheritance property, now Profile hase the Name and Age fields directly but have to keep in mind while assigning data we have assign the embaded struct data as a hole but properties can be directly accessed.)
	Img string
}

type NewProfile struct {
	User
	ImgURL string
}

func Embade() {
	user1 := User{"deb", 22}
	fmt.Println(user1)
	profile1 := Profile{UserData: User{"deb", 22}, Img: "http://localhost:3333/images/deb.png"}
	fmt.Println(profile1)
	fmt.Println(profile1.UserData.Name) // correct
	// fmt.Println(profile1.Name)// wrong

	var profile2 NewProfile = NewProfile{User: user1, ImgURL: "http://localhost:3333/images/deb.png"}
	fmt.Println(profile2)
	fmt.Println(profile2.Name)

}
