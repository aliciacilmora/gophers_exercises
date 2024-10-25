# Choose Your Own Adventure (CYOA)

A simple, web-based "Choose Your Own Adventure" story application built with Go. This project leverages Docker for easy deployment and is designed to work with a multi-stage build Dockerfile for optimal image size. The application runs on a Go server, serving a story in JSON format, and can be deployed to Azure or any cloud service supporting Docker containers.

## Features
* Interactive Story Application: Reads story data from a JSON file and serves it as an interactive "Choose Your Own Adventure" experience.
* Dockerized Deployment: Multi-stage Dockerfile ensures a small final image size.
* Automated CI/CD: GitHub Actions workflow automates building, testing, and pushing the Docker image to Docker Hub on every change to the main branch.

## Installation 

### Install Docker
#### Create an Ubuntu EC2 Instance on AWS and run the below commands to install docker.
```bash
sudo apt update && sudo apt install -y docker.io
```

#### Start Docker and Grant Access
```bash
sudo usermod -aG docker <username_for_instance>
sudo systemctl start docker
sudo systemctl status docker
```

### Clone the repo on cloud instance
```bash
git clone --no-checkout https://github.com/aliciacilmora/gophers_exercises.git
cd gophers_exercises/choose_your_own_adventure
git sparse-checkout init --cone
git sparse-checkout set choose_your_own_adventure
git checkout main
```

### Build the Docker image
```bash
docker build -t <dockerhub_username/name:tag> .
```

### Run the Docker container
```bash
docker run -d -p 3000:3000 <container_name:tag>
```

### Check the logs using Container id
```bash
docker ps # container_id
docker logs <container_id>
docker stop <container_id> # to stop the container
```

## Sample Dockerfile (Multi-Stage)
```dockerfile
# Stage 1: Build the Go application
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /app/bin/cyoaweb ./cmd/cyoaweb

# Stage 2: Minimal runtime image
FROM scratch
COPY --from=build /app/bin/cyoaweb /bin/cyoaweb
COPY --from=build /app/gophers.json /app/gophers.json
WORKDIR /app
CMD ["/bin/cyoaweb"]
```