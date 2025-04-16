FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY . .

RUN <<EOF
go mod tidy 
go build -o mcp-server-govbox cmd/app/main.go
EOF

FROM scratch AS runner
WORKDIR /app
COPY --from=builder /app/mcp-server-govbox .
ENTRYPOINT ["./mcp-server-govbox"]