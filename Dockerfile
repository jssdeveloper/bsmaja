FROM golang:1.21.4-alpine3.18
WORKDIR /app
EXPOSE 80
COPY . .
RUN go mod tidy
RUN go build -o app
CMD ["./app"]