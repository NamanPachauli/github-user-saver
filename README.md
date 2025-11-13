
GITHUB USER SAVER - GOLANG| DOCKER | KUBERNETES | GIT | GITHUB.

## 📌 Overview
GitHub Saver is a **Go-based microservice** that fetches and stores GitHub user information using the **GitHub REST API**.  
It is fully **containerized with Docker** and deployable on **Kubernetes**, making it a great mix of **Backend Development** and **DevOps** skills.

## ⚙️ Tech Stack
| Layer | Technology |
|:------|:------------|
| Language | Go (Golang) |
| Containerization | Docker |
| Orchestration | Kubernetes |
| Version Control | Git & GitHub |
| API | GitHub REST API |
| Optional (Future) | MongoDB / PostgreSQL for persistent storage |

## 🚀 Features
- Fetch GitHub user details (name, followers, repos, etc.)
- Save user data (in-memory or future DB)
- REST API built with Go’s `net/http` package
- Dockerized for consistent deployments
- Kubernetes manifests for scaling and service exposure

---

## 🧩 Project Structure
├── githubsaver.go # main service file
├── Dockerfile # container image definition
├── deployment.yaml # Kubernetes deployment config
├── service.yaml # Kubernetes service config
└── README.md # project documentation

---

## 🧠 How It Works
1. User sends a request → `/user/{username}`
2. Service fetches user info from GitHub API
3. Returns structured JSON output
4. (Optional) Data can be saved to a database

---

## 🐳 Docker Commands
```bash
# Build Docker image
docker build -t github-saver .

# Run container
docker run -p 9095:9095 github-saver

# Deploy application
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

# Check pods and services
kubectl get pods
kubectl get svc

GET http://localhost:9095/user/namanpachauli

Response:

{
  "login": "namanpachauli",
  "name": "Naman Pachauli",
  "followers": 10,
  "public_repos": 5
}



