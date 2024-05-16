FROM golang:alpine as build_container
WORKDIR /app
COPY ./Fakultet/backend/go.mod ./Fakultet/backend/go.sum ./
RUN go mod download
COPY ./Fakultet/backend ./
RUN go build -o server

FROM alpine
WORKDIR /usr/bin
COPY --from=build_container /app/server .
EXPOSE 8001
ENTRYPOINT ["server"]
