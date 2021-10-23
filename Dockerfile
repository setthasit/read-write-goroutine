FROM golang:1.16 AS build_base

WORKDIR /app/read-digit
# Copy everything from the current directory to docker image
COPY . .
# Download and install dependencies
RUN go mod tidy

FROM build_base AS builder
# Build the app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -o read-digit

# Run the executable
CMD ["read-digit"]

FROM alpine

ENV DIGITS_COUNT 2048

RUN apk add --no-cache ca-certificates apache2-utils tzdata
WORKDIR /app/cmd/app
COPY --from=builder /app/read-digit .

CMD ./read-digit ${DIGITS_COUNT}
