# MockSrv – A Simple Mock REST API Server in Go made for practice

MockSrv is a lightweight Go application that turns a JSON file into a full CRUD API with minimal configuration. It is ideal for prototyping, testing front-end applications, or simulating backends in development.

## ✨ Features

- Automatically generates CRUD endpoints from a `db.json` file
- Supports filtering via query parameters
- No external dependencies
- Easy to extend

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-username/mocksrv.git
cd mocksrv
```

### 2. Create your `db.json`

Create a `db.json` file in the root directory with one or more collections. For example:

```json
{
  "users": [
    { "id": 1, "name": "Alice", "email": "alice@example.com" },
    { "id": 2, "name": "Bob", "email": "bob@example.com" }
  ],
  "posts": [
    { "id": 1, "title": "Hello World", "content": "My first post" }
  ]
}
```

### 3. Run the server

```bash
go run main.go
```

Server will start on `http://localhost:8000`.

---

## 🛠 API Endpoints

Assuming a collection `users`, the following routes will be available:

| Method | Endpoint           | Description              |
|--------|--------------------|--------------------------|
| GET    | `/users`           | Get all users (supports filters) |
| GET    | `/users/{id}`      | Get user by ID           |
| POST   | `/users`           | Create a new user        |
| PUT    | `/users/{id}`      | Update a user            |
| DELETE | `/users/{id}`      | Delete a user            |

### Example: Filter users by name

```http
GET /users?name=Alice
```

---

## 🧩 Project Structure

```
mocksrv/
│
├── main.go              # Server bootstrap
├── db/
│   └── database.go      # File-backed data storage and manipulation
├── handlers/
│   └── handler.go       # HTTP handlers for CRUD operations
└── db.json              # Your mock data file (not included by default)
```

---

## 📦 Dependencies

Only uses Go standard library (`net/http`, `encoding/json`, etc.)

---

## 🔒 Notes

- Each object **must** have a numeric `id` field.
- The server reads and writes directly to `db.json` for persistence.
- Filtering is basic: it only supports matching one query parameter value per field.

---

## 📜 License

MIT — feel free to use and adapt.

---

## 👨‍💻 Author

Created by Daniel Zietala