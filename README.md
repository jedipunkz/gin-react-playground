# Gin and React Todo List

This repository contains a simple Todo List application built with Gin, a web framework for Go, and React.js, a JavaScript library for building user interfaces. It's intended for learning purposes and demonstrates how to set up a full-stack application with a Go back-end and a React front-end.

## Getting Started

### Prerequisites

Make sure you have the following installed on your system:

1. Go (Version 1.16 or later)
2. Node.js (Version 14.0.0 or later)

### Running the Application

1. Clone this repository:

       git clone https://github.com/jedipunkz/gin-react-playground.git

2. Navigate to the `backend` directory, and run the Go server:

       cd gin-react-todo-list/backend
       go run main.go

   The Go server will start running on `http://localhost:8080`.

3. In another terminal window, navigate to the `frontend` directory, and install the necessary npm packages:

       cd ../frontend
       npm install

4. Start the React application:

       npm start

   The React application will start running on `http://localhost:3000`. You can now add and delete items from your todo list.

## Author

This project was created by [jedipunkz](https://github.com/jedipunkz).

