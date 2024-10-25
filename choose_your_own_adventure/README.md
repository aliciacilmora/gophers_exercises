[![Completed](https://img.shields.io/badge/project%20status-completed-green.svg?style=for-the-badge)](http://20.118.227.150:3000/story/) [![demo: ->](https://img.shields.io/badge/demo-%E2%86%92-blue.svg?style=for-the-badge)](http://20.118.227.150:3000/story/)

# Choose Your Own Adventure (CYOA)

Choose Your Own Adventure is (was?) a series of books intended for children where as you read you would occasionally be given options about how you want to proceed. For instance, you might read about a boy walking in a cave when he stumbles across a dark passage or a ladder leading to an upper level and the reader will be presented with two options like:

Turn to page 44 to go up the ladder.
Turn to page 87 to venture down the dark passage.
The goal of this exercise is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc).

Stories will be provided via a JSON file with the following format:

```json
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
```

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