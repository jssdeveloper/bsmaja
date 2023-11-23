FROM golang:1.21.1
WORKDIR /app
EXPOSE 3000
COPY . .
RUN go mod tidy
RUN go build -o app .