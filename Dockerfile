# Base image: specify the environment
FROM golang:1.19-alpine AS build

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to working directory
COPY go.mod go.sum ./

# Run go mod download
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the application
RUN go build -o ./out/dist ./cmd/api

# Use a smaller base image for the final image
FROM alpine:3.14 AS prod

# Copy html files
COPY template ./template

# Copy only the necessary files from the build image
COPY --from=build /app/out/dist /app/dist

# Set the working directory
WORKDIR /app

# Set the entry point
CMD ["/app/dist"]
