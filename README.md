# HTTP Job Queue (Go + MongoDB)

This project is a simple, FIFO-compliant HTTP-based job queue implemented in **Go** with **MongoDB** as the persistent backend. It is containerized using Docker and supports checkpoint-based job fetching.

---

## Features

- Add jobs via `POST /job`
- Retrieve jobs FIFO-style via `GET /job`
- Supports checkpointing
- Jobs are persisted in MongoDB
- Lightweight and containerized (Docker-ready)

---

## Requirements

- Go 1.20+
- Docker & Docker Compose
- MongoDB 5.0+
---

## Installation & Setup

1. **Clone the repo**:
   ```bash
   git clone https://github.com/ashutosh5786/Muzz_GO.git
   cd Muzz_GO
   ```
2. **Create .env**:

   ```bash
   MONGO_URI=mongodb://mongo:27017
   ```

   **Note**: It's already present in the Zip

3. **Run with Docker Compose**:
   ```bash
   docker-compose up --build
   ```
4. App will be available at: http://localhost:8080
---
## API Endpoints

### POST /job

Add a new job to the queue.

**Request Body (raw text or JSON):**<br>

```bash
curl -d "test" localhost:8080/job
```

**Response**<br>

```bash
{
  "message": "Job created successfully",
  "jobId": "d3a76b6f-4b1b-4b1e-bc4b-2c4733b4a317"
}
```

### GET /job

Fetch job in FIFO order

**Query Param**

- `amount` (optional): max number of jobs (default: 50)
- `checkpoint` (optional): return jobs **after** this jobId

**Example**

```bash
curl "http://localhost:8080/job?amount=3"
```

**Response**

```bash
{
  "jobs": [
    {
      "jobId": "abc123",
      "job": "send_email"
    },
    {
      "jobId": "abc124",
      "job": "generate_invoice"
    }
  ]
}
```
---
## Design Choices

- **Go**: Efficient for concurrent HTTP servers and aligns with Muzz's preferred language.

- **MongoDB**: Provides persistent storage, native timestamp support for ordering, and simple querying.

- **UUIDs**: Ensures each job has a globally unique identifier.

- **Fiber Framework**: Lightweight and fast web framework for Go, inspired by Express.js.
---

## Future Development Ideas
- Job deletion after processing
- Job retries or dead-letter queues
- Authentication & rate-limiting
- Kafka or Redis stream adapter (swap MongoDB)