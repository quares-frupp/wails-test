package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "World!"
}

func GetUsers() []User {
	p1 := User{"Nombre", "email@email.com", 3}
	p2 := User{"Nombre2", "email@email.com", 5}
	p3 := User{"Nombre3", "email@email.com", 67}

	return []User{p1, p2, p3}
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Wails Test",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(GetUsers)
	app.Run()
}
