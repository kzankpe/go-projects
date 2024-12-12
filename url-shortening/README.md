# URL Shortening Service

Project Url : https://roadmap.sh/projects/url-shortening-service

Golang Rest API solution for the url shortening [challenge](https://roadmap.sh/projects/url-shortening-service) from [roadmap.sh](https://roadmap.sh/).

This project provides a  RESTful API that allows users to shorten long URLs. The API provides endpoints to create, retrieve, update, and delete short URLs. It should also provide statistics on the number of times a short URL has been accessed. 

## Features

- Create a new short URL
- Retrieve an original URL from a short URL
- Update an existing short URL
- Delete an existing short URL
- Get statistics on the short URL (e.g., number of times accessed)

## Installation

To run this application, follow these steps:
Create an .env file in the roor dir (next to the main.go) using the example ``.env.example``. Make sure that the postgresql server and data are available and accessible by the application.


### Run application locally 

```bash
git clone https://github.com/kzankpe/go-projects.git
```
Run the following command to build and run the project:

```bash
cd url-shortening
go build -o url-shortening
```

Run the server:
```bash
./ulr-shortening # Run the application locally on port 8090
```

### Run application in docker

```bash
git clone https://github.com/kzankpe/go-projects.git
```
Run the following command to build and run the project:

Build the Docker application

```bash
cd url-shortening
docker build -t url-shortening .
```

Run the Docker application

```bash
docker run -p 8090:8090 url-shortening
```

The application should be available on ``http://localhost:8090``