FROM golang:alpine as build_container
WORKDIR /app
COPY ./auth-service/go.mod ./auth-service/go.sum ./
RUN go mod download
COPY ./auth-service/ .
RUN go build -o server

FROM alpine
WORKDIR /usr/bin
COPY --from=build_container /app/server .
EXPOSE 8000
ENTRYPOINT ["server"]
