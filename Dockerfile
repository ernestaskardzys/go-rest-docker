FROM golang:alpine3.10 as builder

LABEL maintainer="Ernestas Kardzys <ernestas.kardzys@gmail.com>"

WORKDIR /app
COPY /src /app

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

RUN go build -o main .

# Minimize the image
FROM alpine:3.9.5

WORKDIR /app

# Copy executable from previous stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
