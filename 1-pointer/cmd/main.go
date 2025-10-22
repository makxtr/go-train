package main

import "fmt"

type User struct {
	Name string
}

func main() {
	user := User{Name: "Олег"}
	fmt.Println("Имя до обновления:", user.Name) // Олег

	updateUser(user)
	fmt.Println("Имя после обновления:", user.Name) // Олег

}

func updateUser(u User) {
	u.Name = "Таненбаум"
	fmt.Println("Имя внутри функции [updateUser]:", u.Name) // Таненбаум

	resetUser(&u)
	fmt.Println("Имя после вызова функции [resetUser] внутри updateUser", u.Name)
}

func resetUser(u *User) {
	u = &User{Name: "Безымянный"}
	fmt.Println("Имя внутри функции [resetUser]:", u.Name) // Безымянный
}
