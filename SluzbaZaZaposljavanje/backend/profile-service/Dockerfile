FROM golang:alpine as build_container
WORKDIR /app
COPY ./SluzbaZaZaposljavanje/backend/profile-service/go.mod ./SluzbaZaZaposljavanje/backend/profile-service/go.sum ./
RUN go mod download
COPY ./SluzbaZaZaposljavanje/backend/profile-service ./
RUN go build -o server

FROM alpine
WORKDIR /usr/bin
COPY --from=build_container /app/server .
EXPOSE 8000
ENTRYPOINT ["server"]
