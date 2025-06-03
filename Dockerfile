# 빌드 스테이지
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# 빌드에 필요한 패키지 설치
RUN apk add --no-cache git

# 소스 코드와 필요한 파일들 복사
COPY go.mod go.sum ./
COPY abi.json ./
COPY . .

# 의존성 다운로드 및 빌드
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bsh-api main.go

# 실행 스테이지
FROM alpine:latest

WORKDIR /app

# SSL 인증서 및 타임존 설정
RUN apk add --no-cache ca-certificates tzdata

# 빌드된 바이너리와 필요한 파일들 복사
COPY --from=builder /app/bsh-api .
COPY .env .
COPY abi.json .

# 실행 권한 설정
RUN chmod +x bsh-api

# 포트 노출
EXPOSE 1111

# 실행
CMD ["./bsh-api"] 