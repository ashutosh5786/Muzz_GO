package main

import (
	"context"
	"log"
	"os"
	

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Job struct to represent a job
// TODO Needed to generate a unique jobId for each job. For now, we are using a simple int
// TODO Needed to change the jobId to string and use UUID for unique jobId generated automatically
type Job struct {
	jobId int    `json:"jobId" bson:"jobId"`
	job   string `json:"job" bson:"job"`
}

// JobList struct to represent a list of jobs. Making a Slice as it can grow in size depending upon the struct job which is declate above
// type JobList struct {
// 	Jobs []Job `json:"jobs"`
// }

// var jobList = JobList{
// 	Jobs: []Job{},
// }

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Getting MonggDB connection string from the .env file
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	// Ping the MongoDB server to check the connection
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")
	


	// Initialize a new Fiber app
	app := fiber.New()

	// Set up a simple GET endpoint to return the job list TODO Needed to get the data from Any DB
	app.Get("/job", func(c fiber.Ctx) error {
		amount, checkpoint := c.Query("amount"), c.Query("checkpoint")

		// REMOVE IT AFTER TESTING
		log.Printf("Amount: %s, Checkpoint: %s", amount, checkpoint)

		// return c.JSON(jobList)
		return c.SendString("Job Get API")
	})

	app.Post("/job", func(c fiber.Ctx) error {
		return c.SendString("Job Post API")
	})

	// Start the server on port 80 make the docker compose to serve on port 8080 TODO Dont Forget about it in the end
	log.Fatal(app.Listen(":80"))
}
