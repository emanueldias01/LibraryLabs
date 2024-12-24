FROM golang:1.23.2

EXPOSE 8000

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

CMD [ "go" , "run", "main.go" ]