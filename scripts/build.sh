#!/bin/bash

VERSION="0.1.0"
PLATFORMS=("darwin/amd64" "linux/amd64")
OUTPUT_DIR="dist"

# 출력 디렉토리 생성
mkdir -p $OUTPUT_DIR

# 각 플랫폼별로 빌드
for PLATFORM in "${PLATFORMS[@]}"
do
    # 플랫폼 정보 파싱
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    
    # 빌드
    echo "Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_DIR/bsh-$GOOS-$GOARCH cmd/bsh/main.go
    
    # tar.gz 생성
    cd $OUTPUT_DIR
    tar -czf bsh-$GOOS-$GOARCH.tar.gz bsh-$GOOS-$GOARCH
    cd ..
    
    # SHA256 계산
    SHA256=$(shasum -a 256 $OUTPUT_DIR/bsh-$GOOS-$GOARCH.tar.gz | awk '{print $1}')
    echo "SHA256 for $GOOS/$GOARCH: $SHA256"
done

echo "Build complete! Files are in the $OUTPUT_DIR directory"
echo "Don't forget to update the SHA256 values in Formula/bsh.rb" 