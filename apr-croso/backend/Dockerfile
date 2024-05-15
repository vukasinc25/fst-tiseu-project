FROM golang:alpine as build_container
WORKDIR /app
COPY ./apr-croso/backend/go.mod ./apr-croso/backend/go.sum ./
RUN go mod download
COPY ./apr-croso/backend ./
RUN go build -o server

FROM alpine
WORKDIR /usr/bin
COPY --from=build_container /app/server .
EXPOSE 8003
ENTRYPOINT ["server"]