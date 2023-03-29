package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	fmt.Println(os.Getenv("SPOTIFY_CLIENT_SECRET"))
}
