FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN apk add --no-cache gcc musl-dev sqlite-dev
RUN go mod tidy
RUN go build -o main .
CMD ["./main"]