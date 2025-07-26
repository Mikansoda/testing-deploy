# About the project: API Deployment to VPS
This project is an API backend system, deployed to VPS using Go, Gin, Docker, and GitHub Actions.

# Tech stack
1. Golang 1.24.2
2. Gin Web Framework
3. GitHub Actions (CI/CD)
4. VPS (ubuntu)
5. Postman

# How to use
1. Clone the repositoruy
2. Use cd student7 to change into the project directory
3. Run cp .env.example .env to create a new .env file from the example
4. Run the server using go run main.go on the terminal

# Documentation:
The APIs within this project is built using Go and Gin Web Framework, deployed to VPS. Including routes like ping, greeting,  GET and POST users data (dummy, stored within memory). Endpoints are testable locally or after deployment to VPS.

# API Endpoints
- **GET** `/v1/ping`  
  Returns "message": "Pong!" as message

- **GET** `/v1/greeting`  
  Returns "message": "Hello, Zahra!" as message

- **GET** `/v1/about`  
  Get basic information about of the project as message

- **GET** `/v1/users`  
  Retrieve all user data (dummy/saved in-memory)

- **POST** `/v1/users`  
  Add a new user to the dummy (dummy/saved in-memory within slice)


# Example responses
**GET** `/v1/users`
[
  {
    "id": 1,
    "name": "Clara"
  },
  {
    "id": 2,
    "name": "Miki"
  },
  {
    "id": 3,
    "name": "Chiki"
  },
  {
    "id": 4,
    "name": "Suneo"
  },
  {
    "id": 5,
    "name": "Jojo"
  }
]