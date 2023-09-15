# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd
RUN GOOS=linux go build -o main

EXPOSE 8080

CMD ["./main"]

# docker build -f ".\build\Dockerfile" --progress=plain -t gavriladockeril/task18:latest .
# docker run --rm -p 8080:8080 gavriladockeril/task18:latest