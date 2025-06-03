#!/bin/bash

# 필수 파라미터 확인
if [ "$#" -ne 4 ]; then
    echo "Usage: $0 <ip> <port> <username> <password>"
    exit 1
fi

IP=$1
PORT=$2
USERNAME=$3
PASSWORD=$4

# SSH 연결 테스트
sshpass -p "$PASSWORD" ssh -o StrictHostKeyChecking=no -o ConnectTimeout=5 -p "$PORT" "$USERNAME@$IP" "echo 'SSH connection successful'"

if [ $? -eq 0 ]; then
    echo "SSH connection test passed"
    exit 0
else
    echo "SSH connection test failed"
    exit 1
fi 