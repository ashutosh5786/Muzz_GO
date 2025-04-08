Muzz_GO
Muzz_GO is a Go-based RESTful API for managing jobs, built using the Fiber web framework and MongoDB as the database. The application supports creating and retrieving jobs, with features like pagination and checkpoint-based filtering. It is containerized using Docker and can be deployed with Docker Compose.

Features
Create Jobs: Add new jobs with a unique jobId and a timestamp.
Retrieve Jobs: Fetch jobs with optional pagination and checkpoint-based filtering.
MongoDB Integration: Uses MongoDB for persistent storage.
Dockerized: Easily deployable with Docker and Docker Compose.
Installation
Prerequisites
Go (version 1.20 or later)
Docker
Docker Compose
Steps to Run Locally
Clone the Repository:

Set Up Environment Variables: Create a .env file in the root directory with the following content:

Run the Application with Docker Compose:

Access the Application:

API: http://localhost:8080
MongoDB: mongodb://localhost:27017
API Endpoints

1. Create a Job
   Endpoint: POST /job
   Description: Adds a new job to the database.
   Request:
   Response:
2. Retrieve Jobs
   Endpoint: GET /job
   Description: Fetches jobs with optional pagination and checkpoint filtering.
   Query Parameters:
   amount (optional): Number of jobs to return (default: 50).
   checkpoint (optional): Filters jobs created after the specified jobId.
   Request:
   Response:
   Project Structure
   Choices We Made and Why
   Fiber Framework:

Chosen for its lightweight and high-performance nature, making it ideal for building RESTful APIs.
MongoDB:

Selected for its flexibility and ability to handle unstructured data, which suits the job management use case.
Docker and Docker Compose:

Used to simplify deployment and ensure consistency across environments.
Pagination and Checkpoint Filtering:

Implemented to handle large datasets efficiently and allow incremental data retrieval.
Environment Variables:

Used to manage sensitive configurations like the MongoDB connection string.
Future Improvements
Authentication and Authorization:

Add user authentication (e.g., JWT) to secure the API.
Validation:

Implement stricter validation for input data to ensure data integrity.
Error Handling:

Improve error messages and add structured logging for better debugging.
Testing:

Add unit tests and integration tests to ensure code reliability.
Scalability:

Use a load balancer and horizontal scaling to handle increased traffic.
Monitoring:

Integrate monitoring tools like Prometheus and Grafana to track application performance.
