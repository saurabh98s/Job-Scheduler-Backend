FROM golang
# as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

Copy . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# FROM scratch

# COPY --from=builder /app/go-docker /app/
EXPOSE 8080
ENTRYPOINT ["/app/go-docker"]