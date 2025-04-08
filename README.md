# URL Shortener in Go

A simple and efficient URL shortener service built using Go. This project allows users to shorten long URLs and retrieve the original URLs using the generated short codes.

## Features
- Shorten long URLs into compact, shareable links.
- Retrieve original URLs using short codes.
- Lightweight and fast implementation using Go.

## Prerequisites
- Go (Golang) installed on your system.
- Basic understanding of Go and REST APIs.

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/rajsingh112005-go.git
   cd url-shortener-go
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. Access the service at `http://localhost:8080`.

## Usage
- **Shorten a URL**: Send a POST request to `/shorten` with the long URL in the request body.
- **Retrieve original URL**: Send a GET request to `/<short-code>`.

Example using `curl`:
```bash
# Shorten a URL
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/shorten

# Retrieve original URL
curl http://localhost:8080/<short-code>
```



