FROM golang:1.24-alpine

WORKDIR /app

# 필요한 패키지 설치
RUN apk add --no-cache gcc musl-dev

# 소스 코드 복사
COPY . .

# 의존성 다운로드 및 빌드
RUN go mod download
RUN go build -o bsh-api main.go

# 2222 포트 노출
EXPOSE 2222

# 실행
CMD ["./bsh-api"] 