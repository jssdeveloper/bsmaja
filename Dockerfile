FROM golang:1.21.4-alpine3.18
RUN apk add git
WORKDIR /
RUN git clone https://github.com/jssdeveloper/bsmaja.git
RUN mkdir app
RUN mv bsmaja/* app/
WORKDIR /app
EXPOSE 80
# COPY . .
RUN go mod tidy
RUN go build -o app
CMD ["./app"]