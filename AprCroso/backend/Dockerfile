FROM golang:alpine as build_container
WORKDIR /app
COPY ./AprCroso/backend/go.mod ./AprCroso/backend/go.sum ./
RUN go mod download
COPY ./AprCroso/backend ./
RUN go build -o server

FROM alpine
WORKDIR /usr/bin
COPY --from=build_container /app/server .
EXPOSE 8004
ENTRYPOINT ["server"]
