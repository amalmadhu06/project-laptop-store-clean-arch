# Base image : specify the enviornment
FROM golang:1.19-alpine AS build
# working directory
WORKDIR /app
# copy go.mod and go.sum to working directory
COPY go.mod go.sum ./
# run go mod tidy
RUN go mod download
# Copy the rest of the files
COPY . .
# Build the application
RUN go build -o ./out/dist ./cmd/api

# Use a smaller base image for the final image
FROM alpine:3.14 AS prod
# copy html files
COPY template ./template
# Copy only the necessary files from the build image
COPY --from=build /app/out/dist /app/dist
# Set the working directory
WORKDIR /app
# Set the entry point
CMD ["/app/dist"]