package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var client *mongo.Client

func connectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://root:example@mongo1:27017/?authSource=admin&replicaSet=rs0"))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("✅ Connected to MongoDB Replica Set!")
}

type Data struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getData(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("testdb").Collection("data")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var results []Data
	if err := cursor.All(ctx, &results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func postData(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("testdb").Collection("data")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func main() {
	connectMongo()
	http.HandleFunc("/get", getData)
	http.HandleFunc("/post", postData)
	log.Println("🚀 Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
