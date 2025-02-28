# ✅ To-Do Manager (Golang + Gin)

A simple yet powerful **To-Do Manager** built using **Golang**, featuring **REST API** with **Gin Framework**, **JWT Authentication**, **Zerolog for logging**, and **PostgreSQL** for storage.

## 🔥 Features

- ✅ **CRUD Operations** – Create, Read, Update, and Delete tasks
- ✅ **JWT Authentication** – Secure user authentication
- ✅ **Role-Based Access** – Manage user permissions
- ✅ **Logging** – Structured logs using Zerolog
- ✅ **Dependency Injection** – Efficient code structuring

## 🛠️ Tech Stack

- **Golang** – Backend development
- **Gin Framework** – Lightweight and high-performance web framework
- **PostgreSQL** – Relational database for task storage
- **Zerolog** – Fast and structured logging
- **JWT Authentication** – Secure user sessions

## ⚙️ Installation & Setup

### Clone the Repository
```sh
git clone https://github.com/Purvig648/To-Do-Manager.git
cd To-Do-Manager
```

### Install Dependencies
```sh
go mod tidy
```

### Set Up PostgreSQL Database
1. Install and start PostgreSQL.
2. Update the `.env` file with your database credentials:

### Run the Application
```sh
go run .\cmd\to-do-manager\main.go
```

The server should now be running at `http://localhost:8081`. You can test API endpoints using **Postman** or **cURL**.

## 📡 API Endpoints
### Tasks

| Method | Endpoint | Description |
|--------|---------|-------------|
| `GET` | `/tasks` | Fetch all tasks |
| `POST` | `/tasks` | Create a new task |
| `PUT` | `/tasks/:id` | Update a task |
| `DELETE` | `/tasks/:id` | Delete a task |

## 📌 Future Improvements

- Implement task prioritization and categorization
- Add email notifications for due tasks
- Introduce a mobile-friendly frontend

## 🤝 Contributing

Pull requests are welcome! Feel free to fork the repo and submit improvements.

## 📜 License

This project is licensed under the **MIT License**.

