FROM golang:1.23.2-alpine as build

WORKDIR /app

COPY ./models /app/models/
COPY ./database /app/database/
COPY ./migrations /app/migrations/
COPY ./repository /app/repository/
COPY ./service /app/service
COPY ./controllers /app/controllers/
COPY ./routes /app/routes
COPY ./main.go /app/
COPY ./go.mod /app/
COPY ./go.sum /app/

RUN go build main.go

FROM alpine:latest as prod

EXPOSE 8000

WORKDIR /app

COPY --from=build /app/main /app/

CMD [ "./main" ]