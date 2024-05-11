FROM --platform=linux/amd64 golang:1.23
WORKDIR /app
COPY . /app
RUN go mod download  
EXPOSE 8080
CMD ["go", "run", "main.go"]