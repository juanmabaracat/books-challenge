FROM golang:1.21
RUN mkdir /app
WORKDIR /app
# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o /books-api ./cmd

EXPOSE 8080

# Run
CMD ["/books-api"]