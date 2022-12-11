# syntax=docker/dockerfile:1

## Build for golang
FROM golang:1.19-alpine AS build-go

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o home-server main.go


## Build for vue
FROM node:lts-alpine AS build-vue

WORKDIR /app

COPY /frontend/package*.json ./
RUN npm install

COPY ./frontend ./

RUN npm run build


## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build-go ./app/home-server ./
COPY --from=build-vue ./app/dist ./frontend/dist

EXPOSE 3000

CMD [ "./home-server" ]

