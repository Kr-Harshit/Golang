package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"githhub.com/Kr-Harshit/golang-react-todo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// loadTheEnv()
	createDBInstance()
}

// func loadTheEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading the .env file; ", err)
// 	}
// }

// database setup
var collection *mongo.Collection

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")

	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")
}

//middlerware
func getAllTasks() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var results []primitive.M
	for cursor.Next(context.Background()) {
		var result bson.M
		e := cursor.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func insertOneTask(task models.Todo) interface{} {
	result, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Recorded added", result.InsertedID)
	return result.InsertedID
}

func updateTask(todoId int, todo models.Todo) {
	filter := bson.M{"id": todoId}
	result, err := collection.ReplaceOne(context.Background(), filter, todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneTask(taskID int) {
	filter := bson.M{"id": taskID}
	fmt.Println(filter)
	res, err := collection.DeleteOne(context.Background(), filter)
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted document count: ", res.DeletedCount)
}

func deleteAllTasks() int64 {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted document count: ", res.DeletedCount)
	return res.DeletedCount
}

// handler Func

func GetAllTasks(c *gin.Context) {
	c.Header("Context-Type", "application/json")
	payload := getAllTasks()
	fmt.Printf("%+v\n", payload)
	c.IndentedJSON(http.StatusOK, payload)
}

func CreateTask(c *gin.Context) {
	c.Header("Context-Type", "application/json")
	var task models.Todo
	if err := c.BindJSON(&task); err != nil {
		log.Fatal(err)
	}
	insertedId := insertOneTask(task)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task added!", "_id": insertedId})
}

func UpdateTask(c *gin.Context) {
	c.Header("Context-Type", "application/json")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "deletion failed!"})
	}
	var payload models.Todo

	if err := c.BindJSON(&payload); err != nil {
		log.Fatal(err)
	}
	updateTask(id, payload)
	c.IndentedJSON(http.StatusOK, gin.H{"_id": id})
}

func DeleteAllTasks(c *gin.Context) {
	c.Header("Context-Type", "application/json")
	res := deleteAllTasks()
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d tasks deleted", res)})
}

func DeleteTask(c *gin.Context) {
	c.Header("Context-Type", "application/json")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "deletion failed!"})
	}
	deleteOneTask(id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted", "_id": id})
}
