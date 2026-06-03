# Backend Engineer Challenge REST

### Exercise: Task Management API

### Objective:

Build a RESTful API for a simple task management system. The API should allow users to create, read, update, and delete tasks. Each task should have a title, description, status, and due date.

⚠️ All the information you need to successfully complete the challenge is provided here. If you encounter any doubts or questions, we encourage you to make assumptions, use your best judgment, and proceed.

### Requirements:

1. **Setup**:
    - Use Go for the server.
    - Use a database of your choice (MongoDB, PostgreSQL, MySQL).
2. **Endpoints**:
    - **Create a task**: `POST /tasks`
    - **Get all tasks**: `GET /tasks`
    - **Get a single task**: `GET /tasks/:id`
    - **Update a task**: `PUT /tasks/:id`
    - **Delete a task**: `DELETE /tasks/:id`
3. **Task Model**:
    - **Title**: String (required)
    - **Description**: String (optional)
    - **Status**: String (e.g., "pending", "in-progress", "completed") (required)
    - **Due Date**: Date (optional)
4. **Validation**:
    - Ensure all required fields are provided and valid.
    - Return appropriate HTTP status codes (e.g., 400 for bad requests, 404 for not found).
5. **Error Handling**:
    - Handle errors gracefully and return meaningful error messages.
6. **Testing**:
    - Write unit tests for the API endpoints using a testing framework.

### Bonus Points:

- Implement pagination for the `GET /tasks` endpoint.
- Add filtering options to get tasks by status or due date.
- Secure the API using authentication (e.g., JWT).
- Document the API using Swagger.
- A Dockerfile is included.

✅ **All the Bonus Points have been implemented!**

## Deadline

Your challenge must be completed and submitted within seven days. We believe this timeframe will provide you with ample opportunity to familiarize yourself with any new concepts, brainstorm, and implement your solution. We look forward to seeing your innovative approach!

## Getting Started

### Prerequisites:

Make sure you have Docker Desktop and Docker Compose installed.

### Installation & Execution:

1. Clone this repository and navigate to the project root:
   ```bash
   git clone <repository-url>
   cd task-api
   ```

2. Build and spin up the environment with a single command:
    ```bash
    docker compose up --build
    ```

The API gateway will be up and running locally at `http://localhost:8080`.

### Interactive API Docs (Swagger):

The API is fully documented under the OpenAPI specification. You can access the interactive UI to read structural schemas, expected error responses, and test endpoints live.

**Swagger UI URL:** `http://localhost:8080/swagger/index.html`.

### How to Test the Protected Endpoints:

1. Open Swagger UI and find the POST /login endpoint.

2. Click "Try it out" and use the administrative default credentials:
    - **Username:** `admin`
    - **Password:** `admin`

3. Execute the request and copy the generated token string from the 200 OK JSON body.

4. Scroll to the top of the page, click the green "Authorize" button.

5. In the input field, write Bearer followed by a space and your token:
   ```text
   Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
   ```

6. Click Authorize and close the modal. Now you can test all the /tasks endpoints freely!

### Running Unit Tests:

To execute the suite of automated unit tests built for the endpoints, run the following command locally at the project root:
    ```bash
    go test ./internal/handlers -v
    ```
