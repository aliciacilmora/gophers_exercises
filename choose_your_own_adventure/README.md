# Choose Your Own Adventure (CYOA)

A simple, web-based "Choose Your Own Adventure" story application built with Go. This project leverages Docker for easy deployment and is designed to work with a multi-stage build Dockerfile for optimal image size. The application runs on a Go server, serving a story in JSON format, and can be deployed to Azure or any cloud service supporting Docker containers.

## Features
* Interactive Story Application: Reads story data from a JSON file and serves it as an interactive "Choose Your Own Adventure" experience.
* Dockerized Deployment: Multi-stage Dockerfile ensures a small final image size.
* Automated CI/CD: GitHub Actions workflow automates building, testing, and pushing the Docker image to Docker Hub on every change to the main branch.

## Installation 

```
git clone --no-checkout https://github.com/aliciacilmora/gophers_exercises.git
cd gophers_exercises/choose_your_own_adventure
git sparse-checkout init --cone
git sparse-checkout set choose_your_own_adventure
git checkout main
```