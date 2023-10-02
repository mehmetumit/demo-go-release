FROM golang:alpine as builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Seperate copying for caching
# COPY go.mod go.sum ./
COPY go.mod ./

# RUN go mod download

COPY ./ ./

RUN go build -o app cmd/main.go

FROM scratch as final

COPY --from=builder /app/app ./

ENTRYPOINT [ "./app" ]
