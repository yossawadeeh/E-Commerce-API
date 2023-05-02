FROM golang:1.20.3-alpine3.16
WORKDIR /app
COPY . /app
EXPOSE 8001
CMD ["go", "run", "main.go"]