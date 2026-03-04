# --- Build Frontend ---
FROM node:20 AS frontend-builder
WORKDIR /app/frontend
# 安裝package
COPY frontend/package*.json ./
RUN npm install
# 拷貝與編譯
COPY frontend .
RUN npm run build

# --- Build Stage ---
FROM golang:1.25-alpine AS builder
WORKDIR /app
# 安裝必要的系統工具
RUN apk add --no-cache git
# 複製並下載依賴
COPY go.mod go.sum ./
RUN go mod download
# 複製原始碼
COPY . .
# 編譯二進制檔案
RUN go build -o /app/bin/server ./cmd/server/main.go
RUN go build -o /app/bin/scheduler ./cmd/scheduler/main.go

# --- Final Stage: Scheduler ---
FROM alpine:latest AS scheduler
WORKDIR /root/
COPY --from=builder /app/bin/scheduler .
COPY --from=builder /app/config ./config
CMD ["./scheduler"]

# --- Final Stage: API Server ---
FROM alpine:latest AS server
WORKDIR /root/
COPY --from=frontend-builder /app/frontend/dist ./static
COPY --from=builder /app/bin/server .
COPY --from=builder /app/config ./config
# 如果有 .env (但在 Docker 中建議用環境變數)
EXPOSE 8080
CMD ["./server"]