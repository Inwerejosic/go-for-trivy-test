
````md
# Gin Greeting API

A lightweight Go API built with **Gin** that demonstrates:
- Basic REST endpoints
- Docker containerization
- GitHub Actions CI
- Container security scanning with Trivy

---

## 🚀 Features

- `GET /ping` – Health check endpoint
- `POST /greet` – Returns a time-based greeting (Morning, Afternoon, Evening, or Night)
- Dockerized application
- CI pipeline with build, test, and security scan

---

## 📦 API Endpoints

### Health Check

**GET `/ping`**

**Response**
```json
{
  "message": "I am doing this to test trivy installation!"
}
````

---

### Greeting Endpoint

**POST `/greet`**

**Request Body**

```json
{
  "name": "GitHub Actions"
}
```

**Response**

```json
{
  "message": "Good Morning, GitHub Actions"
}
```

> The greeting changes automatically based on the server time.

---

## 🐳 Run Locally with Docker

### Build the image

```bash
docker build -t gin-ping-app .
```

### Run the container

```bash
docker run -p 8080:8080 gin-ping-app
```

### Test the API

```bash
curl http://localhost:8080/ping
```

```bash
curl -X POST http://localhost:8080/greet \
  -H "Content-Type: application/json" \
  -d '{"name":"World"}'
```

---

## 🤖 CI/CD (GitHub Actions)

The GitHub Actions pipeline automatically:

1. Builds the Docker image
2. Runs the container
3. Tests API endpoints
4. Scans the image for vulnerabilities using **Trivy**
5. Fails the build on HIGH or CRITICAL issues

This ensures code quality, correctness, and security on every push or pull request to `main`.

---

## 🔐 Security Scanning

Trivy scans the Docker image for known vulnerabilities.

```bash
trivy image gin-ping-app:latest
```

The CI pipeline fails if **HIGH** or **CRITICAL** vulnerabilities are detected.

---

## 🛠 Tech Stack

* Go
* Gin Web Framework
* Docker
* GitHub Actions
* Trivy
