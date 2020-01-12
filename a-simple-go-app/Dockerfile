FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o a-simple-go-app .
RUN make build

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/a-simple-go-app .
EXPOSE 1487

# Command to run the executable
CMD ["./a-simple-go-app"] 