# The Orc Shack

## Overview
This project is a web API for managing restaurant dish ratings. It uses Go with the Gin framework and interacts with a PostgreSQL database. It supports multitenancy.

## Built With

* [Gin](https://github.com/gin-gonic/gin) - The web framework used
* [GORM](https://gorm.io/) - ORM tool
* [Docker](https://www.docker.com/) - Containerization

## Prerequisites
- Go 1.15 or higher
- PostgreSQL 12.0
- JWT_SECRET_KEY set in your environment variables for authentication purposes
- Go
- Gin
-GORM

## Setting Up Your Environment
1. Clone the repository:
   ```bash
   git clone http://tfg-labs-meluld@git.codesubmit.io/tfg-labs/the-dancing-pony-v2-wqfxtm

2. go mod tidy

3. Make sure you have docker destop installed on your machine
  - Run `docker-compose up --build` to build your container image with all the needed project dependencies

4. The api will be served on localhost port 8080

3. Use any databse postgres database client of your choice to connect to the database and use sql to interact with database tables

5. use Postman or thunderclient or any service to test the api
