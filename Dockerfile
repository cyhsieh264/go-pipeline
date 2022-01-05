# Stage 1: pull official base image and build executable file

FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main

# Stage 2: copy and run executable file

FROM alpine

WORKDIR /app

COPY --from=builder /app/main ./

CMD ["./main"]
