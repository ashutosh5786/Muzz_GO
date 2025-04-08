package main

import (
	"context"
	"log"
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Job struct to represent a job
// TODO Needed to generate a unique jobId for each job. For now, we are using a simple int
// TODO Needed to change the jobId to string and use UUID for unique jobId generated automatically
type Job struct {
	JobId string `json:"jobId" bson:"jobId"`
	Job   string `json:"job" bson:"job"`
}

var jobCollection *mongo.Collection

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

	jobCollection := client.Database("jobDB").Collection("jobs")

	// Initialize a new Fiber app
	app := fiber.New()

	// app.Get("/job", func(c *fiber.Ctx) error {
	// 	amount, checkpoint := c.Query("amount"), c.Query("checkpoint")

	// 	// REMOVE IT AFTER TESTING
	// 	log.Printf("Amount: %s, Checkpoint: %s", amount, checkpoint)

	// 	// return c.JSON(jobList)
	// 	return c.SendString("Job Get API")
	// })
	// Set up a simple GET endpoint to return the job list TODO Needed to get the data from Any DB

	app.Post("/job", func(c *fiber.Ctx) error {

		// Parse the request body to get the job description
		jobDescription := string(c.Body())

		//Creating the struct for the job
		newJob := Job{
			JobId: uuid.New().String(),
			Job:  jobDescription,
		}

		// Insearting the new job into the MongoDB collection
		fmt.Printf("%+v\n", newJob)
		_, err := jobCollection.InsertOne(context.TODO(), newJob)
		if err != nil {
			log.Fatalf("Failed to insert job into MongoDB: %v", err)
			return c.Status(500).SendString("Failed to insert job into MongoDB")
		}
		log.Printf("Inserted job: %v", newJob)
		// Append the new job to the job list
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Job created successfully",
			"jobId":   newJob.JobId,
		})

	})

	// Start the server on port 80 make the docker compose to serve on port 8080 TODO Dont Forget about it in the end
	log.Fatal(app.Listen(":80"))
}
