# Building the binary of the App
FROM golang:1.18 AS build



# `boilerplate` should be replaced with your project name
WORKDIR /src

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

WORKDIR /src/cmd/web

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM ubuntu:latest

WORKDIR /app

# `boilerplate` should be replaced here as well
COPY --from=build /src/cmd/web/ .

# Exposes port 8090 because our program listens on that port
EXPOSE 8090

CMD ["./app"]
