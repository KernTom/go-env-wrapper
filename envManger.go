package env

import (
	"log"

	"github.com/joho/godotenv"
)

//LoadEnvironment ... loads environmentfile where envfile is path to file and has default value ".env"
var LoadEnvironment = func(envfile string) (state bool) {
	if envfile == "" {
		envfile = ".env"
	}
	state = true
	if err := godotenv.Load(envfile); err != nil {
		state = false
		log.Println("File .env not found, reading configuration from ENV")
	}
	return
}
