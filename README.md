<img width="1949" height="1227" alt="image" src="https://github.com/user-attachments/assets/5f13839c-5d10-4959-a1e6-d1bc65ae70b1" />KosKuy – Personal Expense Management REST API
📌 Overview

KosKuy is a RESTful API built using Go (Golang) and Gin Framework to help boarding house residents (students or remote workers) manage their personal living expenses digitally and in a structured way.

Many students and workers living in boarding houses struggle with:

Tracking rent, electricity, water, and daily expenses

Forgetting to record transactions

Difficulty evaluating monthly spending

Lack of structured financial records

KosKuy provides a backend system that allows users to:

Register and login securely

Categorize expenses

Record financial transactions

Generate automatic financial reports

This project was developed as part of a REST API Bootcamp using Go and Gin.

🚀 Features
🔐 Authentication & Security

User registration and login

JWT-based authentication

Protected routes using middleware

Token validation for authorized access

👤 User Management

Get all users

Get user by ID

Update user

Delete user

📂 Category Management

Create expense categories

Retrieve category list

💰 Transaction Management

Record expense transactions

Retrieve transaction history

📊 Financial Reporting

Automatic calculation of total expenses

Financial report endpoint with total expense summary

🛠 Tech Stack

Language: Go

Framework: Gin

Database: PostgreSQL

Authentication: JWT

Testing Tool: Postman

📂 Project Structure
KOSKUY/
├── config/          
│   └── database.go
├── controllers/     
├── middlewares/     
├── models/          
├── routes/          
├── utils/           
├── koskuy_db.sql    
├── koskuy.postman_collection.json
└── main.go

This project follows modular backend architecture separating routing, controller, model, middleware, and utility layers.

🔑 API Endpoints
Public Routes
Method	Endpoint	Description
POST	/register	Register new user
POST	/login	Login user
Protected Routes (JWT Required)
User
Method	Endpoint
GET	/users
GET	/users/:id
PUT	/users/:id
DELETE	/users/:id
Category
Method	Endpoint
GET	/kategori
POST	/kategori
Transaction
Method	Endpoint
GET	/transaksi
POST	/transaksi
Report
Method	Endpoint
GET	/laporan
⚙ Installation & Setup
1️⃣ Clone Repository
git clone ...
cd koskuy
2️⃣ Setup Database

Create PostgreSQL database

Import koskuy_db.sql

3️⃣ Configure Database

Update database credentials inside:

config/database.go
4️⃣ Install Dependencies
go mod tidy
5️⃣ Run Application
go run main.go

Server runs on:

http://localhost:8080

The Postman collection is included:

koskuy.postman_collection.json

You can import it directly into Postman to test all endpoints.

📈 Future Improvements

Expense filtering by date

Monthly financial summary

Docker support

Swagger API documentation

Pagination for transactions

Unit testing implementation
