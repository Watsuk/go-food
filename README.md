
# Go-Food

## Introduction

Go-Food is a project aimed at creating a platform for a food court. The site allows users to order food from various food trucks.

## Getting Started

To launch the site on your local machine, follow these steps:

### Prerequisites

- Docker installed on your system
- Basic knowledge of Docker commands and operations
- Node.js and npm installed for frontend setup

### Installation

1. Clone the repository to your local machine:
   ```sh
   git clone https://github.com/Watsuk/go-food.git
   ```

2. To set up the backend services, navigate to the `backend` directory:
   ```sh
   cd go-food/backend
   docker compose up -d --build
   ```
   This command sets up the backend service and any associated databases.

3. For the frontend application, open a new terminal, navigate to the `client` directory:
   ```sh
   cd go-food/client
   npm install
   npm run dev
   ```
   This will install all the required npm packages and start the React application.

### Environment Variables

Before starting the backend, you will need to create a `.env` file in the `backend` directory with the following content:
```env
DB_PASS='root'
```
If you have a custom configuration for the database (e.g., different username or password), make sure to update the `main.go` file in the `backend/src/main` directory accordingly.

### Usage

After starting the services with Docker Compose, you can access:

- **Frontend Application**: Open your web browser and visit `http://localhost:5173` for the React frontend application.
- **Backend API**: The backend API can be accessed via `http://localhost:3000`. This is where the API calls are made from the frontend.
- **Adminer**: To manage the database directly, navigate to `http://localhost:8081` in your web browser. Ensure your `.env` file reflects any custom database configurations.

