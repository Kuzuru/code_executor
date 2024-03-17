FROM golang:1.22-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /docker-dbworker ./main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Yekaterinburg /usr/share/zoneinfo/Asia/Yekaterinburg

ENV TZ Asia/Yekaterinburg

WORKDIR /app

COPY --from=builder /docker-dbworker /app/docker-dbworker

EXPOSE 3031

# Run
CMD ["./docker-dbworker"]