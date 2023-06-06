package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func main() {
	user := User{
		ID:        1,
		FirstName: "Cecep",
		LastName:  "Suracep",
		Email:     "cecep@mail.com",
		IsActive:  true,
	}
	user2 := User{
		ID:        2,
		FirstName: "Ujang",
		LastName:  "rembo",
		Email:     "ujang@mail.com",
		IsActive:  true,
	}

	users := []User{user, user2}
	group := Group{"Padepokan Sambat", user, users, true}

	// fmt.Printf("Group Name : %s\nKetua Group : %s\nNama Anggota : %s\n", group.Name, group.Admin.FirstName,group.Users[1].FirstName)
	// fmt.Println(group.Users[1].FirstName)
	displayGroup(group)
	result := displayUser(user2)
	fmt.Println(result)
}

func displayGroup(group Group){
	fmt.Printf("Name : %s\n", group.Name)
	fmt.Printf("Member count : %d\n", len(group.Users))
}

func displayUser(user User) string{
	return fmt.Sprintf("ID : %d FirstName %s %s. Email : %s. Status : %t", user.ID, user.FirstName, user.LastName, user.Email, user.IsActive)
}
