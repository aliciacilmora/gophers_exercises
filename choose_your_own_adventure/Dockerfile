#######################################
# BASE IMAGE FOR BUILDING THE BINARY
#######################################

FROM golang:1.21-alpine AS build

# Working directory for the application
WORKDIR /app

# Copy go.mod to leverage Docker caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o /app/bin/cyoaweb ./cmd/cyoaweb


###########################################
# MULTISTAGE BUILD FOR THE GO APPLICATION
###########################################

FROM scratch

# Copy the compiled binary
COPY --from=build /app/bin/cyoaweb /bin/cyoaweb

# Copy the story JSON file
COPY --from=build /app/gophers.json /app/gophers.json

# Set the working directory and entry point
WORKDIR /app
CMD ["/bin/cyoaweb"]
