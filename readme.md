# Blog Platform

A simple blog API built with Go, featuring user authentication, role‑based access control, and full CRUD functionality for blog posts.

## 📁 Project Setup

This project follows a clean, layered architecture to keep the codebase organised and maintainable.

### **Architecture & Initialization**
- Structured the project into folders for handlers, services, repositories, middleware, config, utils, and database.
- Initialised a Go module to manage internal packages.
- Set up the database inside the `db` folder and loaded environment variables from `.env`.

### **Authentication & Authorization**
- Implemented middleware to generate and verify JWTs.
- Added helper functions to extract user ID and role from tokens.
- Ensured:
  - Only authenticated users can access protected routes.
  - Users can update/delete **only their own posts**.
  - Role-based access is handled via a switch-based function in the `config` package.

### **Utilities**
- Password hashing and comparison functions for secure credential handling.

### **Domain Models**
- `User` and `Post` structs created with all required fields.

### **Features Implemented**
Built using the handler → service → repository pattern:

- Register a user  
- Login  
- Create a post  
- Get all posts  
- Get a specific post  
- Update own post  
- Delete own post  

Routes are defined in the `router` folder and initialised in `main.go`.

## 🛠 Technologies Used

- **Golang** — backend logic and API implementation  
- **GitHub** — version control and repository hosting  
- **Render** — deployment platform  

## 🚀 How to Use the API

**Register**

POST /register

Body:
{
  "name": "Name",
  "email": "email@example.com",
  "password": "password"
}

**Login (get JWT)**

POST /login

**Create Post**

POST /posts

Body:
{
  "title": "Title",
  "content": "Content"
}

**Get All Posts**

GET /posts

**Get A Single Post**

GET /posts/:id

**Update Post (owner only)**

PUT /posts/:id

**Delete Post (owner only)**

DELETE /posts/:id

