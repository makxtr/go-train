package main

func main() {
	println(handle())
}

type err string

func (e err) Error() string {
	return string(e)
}

func handle() error {
	var er err = "error"
	return er
}
