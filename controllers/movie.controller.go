package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/debapriya36/mongo-go-mux-crud/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI string = "mongodb+srv://go:go@cluster0.ixx7ibe.mongodb.net/"
const dbName string = "goDB"
const collectionName string = "watchlist"

var ctx = context.TODO()

var collection *mongo.Collection

func CheckNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	clinetOption := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clinetOption)
	CheckNilError(err)
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created!")
}

func createMovie(movie models.Movie) *mongo.InsertOneResult {
	insertedMovie, err := collection.InsertOne(ctx, movie)
	CheckNilError(err)
	// fmt.Println("Inserted Movie: ", insertedMovie)
	// fmt.Printf("Inserted Movie : %T", insertedMovie)
	return insertedMovie
}

func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	CheckNilError(err)
	//fmt.Println("Movie: ", movie)
	insertedMovie := createMovie(movie)
	response := map[string]interface{}{"status": "success", "message": "Movie created successfully!", "data": insertedMovie}
	json.NewEncoder(w).Encode(response)
}

func getAllMovies() []bson.M {
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	CheckNilError(err)
	var allMovies []bson.M
	for cursor.Next(ctx) {
		var movie bson.M
		err := cursor.Decode(&movie)
		CheckNilError(err)
		allMovies = append(allMovies, movie)
	}

	defer cursor.Close(ctx)
	return allMovies
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	allMovies := getAllMovies()
	response := map[string]interface{}{
		"status":  "success",
		"message": "All movies fetched successfully!",
		"data":    allMovies,
	}
	json.NewEncoder(w).Encode(response)
}

func updateWatch(movieId string) *mongo.SingleResult {
	id, err := primitive.ObjectIDFromHex(movieId)
	CheckNilError(err)
	filter := bson.M{"_id": id}
	updatedResult := collection.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": bson.M{
			"watched": true,
		},
	})
	//fmt.Println("Updated Result: ", updatedResult)
	//fmt.Printf("Updated Result: %T", updatedResult)
	return updatedResult
}

func UpdateWatchById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	movieId := vars["id"]
	response := map[string]interface{}{
		"status":  "success",
		"message": "Movie updated successfully!",
		"data":    updateWatch(movieId),
	}

	json.NewEncoder(w).Encode(response)
}

func deleteALlMovie() int64 {
	filter := bson.M{}
	deletedResult, err := collection.DeleteMany(ctx, filter)
	CheckNilError(err)
   //fmt.Println("Deleted Result: ", deletedResult.DeletedCount)
	return deletedResult.DeletedCount
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := map[string]interface{}{
		"status":  "success",
		"message": "All movies deleted successfully!",
		"data":    deleteALlMovie(),
	}
	json.NewEncoder(w).Encode(response)
}

func getMovies(name string) []bson.M {
	allMovies := []bson.M{}
	filter := bson.M{
		"name": name,
	}
	cursor, err := collection.Find(ctx, filter)
	CheckNilError(err)
	for cursor.Next(ctx) {
		var movie bson.M
		err := cursor.Decode(&movie)
		CheckNilError(err)
		allMovies = append(allMovies, movie)
	}
	defer cursor.Close(ctx)
	return allMovies
}

func GetMovieByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	name := vars["name"]
	moviesList := getMovies(name)
	response := map[string]interface{}{
		"status":  "success",
		"message": "All movies fetched successfully!",
		"data":    moviesList,
	}
	json.NewEncoder(w).Encode(response)
}
