package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// Job struct to represent a job
// TODO Needed to generate a unique jobId for each job. For now, we are using a simple int
// TODO Needed to change the jobId to string and use UUID for unique jobId generated automatically
type Job struct {
	jobId int    `json:"jobId"`
	job   string `json:"job"`
}

// JobList struct to represent a list of jobs. Making a Slice as it can grow in size depending upon the struct job which is declate above
type JobList struct {
	Jobs []Job `json:"jobs"`
}

var jobList = JobList{
	Jobs: []Job{},
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Set up a simple GET endpoint to return the job list TODO Needed to get the data from Any DB
	app.Get("/job", func(c fiber.Ctx) error {
		amount, checkpoint := c.Query("amount"), c.Query("checkpoint")

		// REMOVE IT AFTER TESTING
		log.Printf("Amount: %s, Checkpoint: %s", amount, checkpoint)

		return c.JSON(jobList)
	})

	app.Post("/job", func(c fiber.Ctx) error {
		return c.SendString("Job Post API")
	})

	// Start the server on port 80 make the docker compose to serve on port 8080 TODO Dont Forget about it in the end
	log.Fatal(app.Listen(":80"))
}
