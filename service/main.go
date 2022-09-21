package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type Ad struct {
	url   string
	users []string
	price int
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

func getAllAds() ([]Ad, error) {
	var bsonAds []bson.M
	var ads []Ad
	var ad Ad
	res, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Panic(err)
	}
	if err = res.All(context.TODO(), &ads); err != nil {
		log.Panic(err)
	}
	if bsonAds == nil {
		return ads, errors.New("Ad not found")
	}
	for _, a := range bsonAds {
		bsonBytes, _ := bson.Marshal(a)
		bson.Unmarshal(bsonBytes, &ad)
		ads = append(ads, ad)
	}
	return ads, nil
}

func subscriber(user string, url string) {
	ad, err := checkInDB(url)
	if err != nil {
		ad.url = url
	}
	ad.users = append(ad.users, user)
	ad.price = priceGetter(url)
	// TODO: Implement adding a structure to a mongo
}

func checkInDB(url string) (Ad, error) {
	var ads []bson.M
	var ad Ad
	res, err := collection.Find(context.TODO(), bson.M{"url": url})
	if err != nil {
		log.Panic(err)
	}
	if err = res.All(context.TODO(), &ads); err != nil {
		log.Panic(err)
	}
	if ads == nil {
		return ad, errors.New("Ad not found")
	}
	bsonBytes, _ := bson.Marshal(ads[0])
	bson.Unmarshal(bsonBytes, &ad)
	return ad, nil
}

func updater() {
	ads, err := getAllAds()
	if err != nil {
		log.Println(err)
	}
	for _, ad := range ads {
		price := priceGetter(ad.url)
		if price != ad.price {
			ad.price = price
			// TODO: Update data in the database
			notifier(ad)
		}
	}
}

func notifier(ad Ad) {
	// TDDO: Implement sending a price change notification
}

func priceGetter(url string) int {
	var price int
	// TODO: Get the price by parseing the page from the URL
	return price
}

func main() {
	db()
	server()
}
