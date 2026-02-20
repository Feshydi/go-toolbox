# ---------------------- Base ----------------------
FROM golang:1.26-bookworm AS base

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .


# ---------------------- Develop ----------------------
FROM base AS develop

CMD ["go", "run", "/src/cmd/app/main.go"]


# ---------------------- Builder ----------------------
FROM base AS builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/app/main.go


# -------------------- Production  --------------------
FROM debian:bookworm-slim AS production

WORKDIR /bin

COPY --from=builder /bin/app ./app

CMD ["/bin/app"]
