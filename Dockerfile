FROM golang:1.13.3-stretch as builder
WORKDIR /app
COPY go.* /
RUN go mod download
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o webapp .
 
EXPOSE 8000
ENTRYPOINT ["./webapp"]
