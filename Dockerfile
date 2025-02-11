FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o teacher-site .

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/teacher-site .

# Copy template files into the final image
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/data ./data
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["/app/teacher-site"]