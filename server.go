package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Job struct to represent a job
type Job struct {
	JobId     string    `json:"jobId" bson:"jobId"`
	Job       string    `json:"job" bson:"job"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

var jobCollection *mongo.Collection

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

	jobCollection = client.Database("jobDB").Collection("jobs")

	// Initialize a new Fiber app
	app := fiber.New()

	// Get /Job Endpoint to retrive the jobs
	app.Get("/job", func(c *fiber.Ctx) error {
		amount, checkpoint := c.Query("amount"), c.Query("checkpoint")

		jobCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
			Keys: bson.D{{Key: "createdAt", Value: 1}},
		})

		// REMOVE IT AFTER TESTING
		log.Printf("Amount: %s, Checkpoint: %s", amount, checkpoint)

		// return c.JSON(jobList)
		return c.SendString("Job Get API")
	})

	app.Post("/job", func(c *fiber.Ctx) error {

		// Parse the request body to get the job description
		jobDescription := string(c.Body())

		//Creating the struct for the job
		newJob := Job{
			JobId:     uuid.New().String(),
			Job:       jobDescription,
			CreatedAt: time.Now(),
		}

		// Insearting the new job into the MongoDB collection
		_, err := jobCollection.InsertOne(context.TODO(), newJob)
		if err != nil {
			log.Printf("Failed to insert job into MongoDB: %v", err)
			return c.Status(500).SendString("Failed to insert job into MongoDB")
		}
		log.Printf("Inserted job: %v", newJob)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Job created successfully",
			"jobId":   newJob.JobId,
		})

	})

	// Start the server on port 80 make the docker compose to serve on port 8080 TODO Dont Forget about it in the end
	log.Fatal(app.Listen(":80"))
}
