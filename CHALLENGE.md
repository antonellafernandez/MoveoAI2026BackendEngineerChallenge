### Exercise: Task Management API

### Objective:

Build a RESTful API for a simple task management system. The API should allow users to create, read, update, and delete tasks. Each task should have a title, description, status, and due date.

<aside>
⚠️ All the information you need to successfully complete the challenge is provided here. If you encounter any doubts or questions, we encourage you to make assumptions, use your best judgment, and proceed.

</aside>

<aside>
☎️ After you've completed the challenge, please email your Github repository to `german@moveo.ai` Use the subject line "Backend Engineer Challenge — **Moveo.AI**"

</aside>

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

## Deadline

Your challenge must be completed and submitted within seven days. We believe this timeframe will provide you with ample opportunity to familiarize yourself with any new concepts, brainstorm, and implement your solution. We look forward to seeing your innovative approach!