package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	fmt.Println(os.Getenv("SPOTIFY_CLIENT_SECRET"))
	_, err := os.Stat("public/build/manifest.json")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exists")
	}
}
