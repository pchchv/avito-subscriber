package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type Ad struct {
	link  string
	users []string
}

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func getAllAds(user string) {}

func subscriber(user string, link string) {
	var ad Ad
	ad.link = link
	ad.users = append(ad.users, user)
	// TODO: Implement adding a structure to a mongo
}

func main() {
	db()
	server()
}
