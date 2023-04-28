package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aparnabhure/golangmongodbapis/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aparnabhure:aparna1234@cluster0.vcbmlab.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

// init run only one time at very first time when app started
func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))

	//client option
	//clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo
	//client, err := mongo.Connect(context.TODO(), clientOption)

	checkError(err)

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	fmt.Println("Mongo DB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance is ready")
}

//Helper functions

// Insert Movie
func insertMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	checkError(err)

	fmt.Println("Inserted Move in DB with ID ", inserted.InsertedID)
}

// Update Movie
func updateMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkError(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkError(err)

	fmt.Println("Number of records updated ", result.ModifiedCount)
}

// Delete Movie
func deleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	checkError(err)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	checkError(err)

	fmt.Println("Number of records deleted ", result.DeletedCount)
}

// Delete all movies
func deleteAllMovies() {
	deletedCounts, err := collection.DeleteMany(context.Background(), bson.D{{}})
	checkError(err)
	fmt.Println("Number of records deleted ", deletedCounts.DeletedCount)
}

// Get all movies
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkError(err)

	defer cursor.Close(context.Background())

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie primitive.M
		err := cursor.Decode(&movie)
		checkError(err)

		movies = append(movies, movie)
	}

	return movies
}

// Controller APIs
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	checkError(err)

	insertMovie(movie)
	json.NewEncoder(w).Encode("Movie inserted")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateMovie(params["id"])
	json.NewEncoder(w).Encode("Movie Updated ")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteMovie(params["id"])
	json.NewEncoder(w).Encode("Movie Deleted")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()
	json.NewEncoder(w).Encode("All Movies Deleted")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
