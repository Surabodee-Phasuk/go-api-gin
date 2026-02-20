# REST API Enhancement with Gin & SQLite
## 🚀 How to Run

### 1. Prerequisites
* **Go**: Version 1.20 or later installed -> [Go Install](https://go.dev/doc/install)
* **Git**: For repository management.

### 2. Setup & Installation
1.  **Clone the repository**:
    ```bash
    git clone <repository-link>
    cd <......>
    ```
2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```
    *This command will download the Gin Gonic framework and the SQLite driver.*

### 3. Running the Server
1.  Execute the application:
    ```bash
    go run main.go
    ```
2.  The server will be live at: `http://localhost:8080`
3.  You will see registered routes (GET, POST, PUT, DELETE) in the terminal logs.

### 4. Testing the API
You can test the endpoints using **Postman** or **Curl**

## Available API endpoints

| Method | Endpoint | Description | Expected Status |
| :--- | :--- | :--- | :--- |
| **GET** | `/students` | List all students | 200 OK |
| **GET** | `/students/:id` | Get student by ID | 200 OK / 404 Not Found |
| **POST** | `/students` | Create student | 201 Created / 400 Bad Request |
| **PUT** | `/students/:id` | Update student | 200 OK / 404 Not Found |
| **DELETE** | `/students/:id` | Delete student | **204 No Content** / 404 Not Found |
