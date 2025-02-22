FROM golang:1.23.0-alpine3.20 as build

ENV TZ=Europe/Moscow

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/gojobtasks ./cmd/job/job_create_new_tasks.go

FROM alpine:3.20

ENV TZ=Europe/Moscow

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
COPY --from=build /app/bin/gojobtasks /app/gojobtasks

WORKDIR /app

RUN chmod +x /app/gojobtasks

CMD ["./gojobtasks"]
