FROM golang:1.22.5-alpine as builder

ADD . /code
WORKDIR /code
RUN CGO_ENABLED=0 GOOS=linux go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o dtonetest ./cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o dtone_migrate ./cmd/automigration/automigration.go



FROM golang:1.22.5-alpine
WORKDIR /app
COPY --from=builder /code/dtonetest .
COPY --from=builder /code/dtone_migrate .
COPY --from=builder /code/.env .

EXPOSE 8080
ENTRYPOINT ["./dtonetest"]

# Set destination for COPY
WORKDIR /app
