package main

import (
	"bytes"
	"fmt"
	"io"
)

type User struct{}

func (u *User) Create() {}
func (u *User) Get()    {}
func (u *User) List()   {}
func (u *User) Delete() {}

type Reader interface {
	Get()
	List()
}

type Writer interface {
	Create()
	Delete()
}

func main() {
	var userReader Reader = &User{}
	userWriter := userReader.(Writer)
	//userWriter.Get()
	_ = userWriter

	var buf *bytes.Buffer
	var buf1 io.Writer

	fmt.Printf("T% %v\n", buf, buf)
	fmt.Printf("T% %v\n", buf1, buf1)
}
