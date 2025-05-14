# BSH (Blockchain Server Hosting)

블록체인 기반의 서버 호스팅 서비스를 위한 CLI 도구입니다.

## 설치 방법

### Homebrew를 통한 설치 (macOS)

```bash
brew tap kimtuna/bsh
brew install bsh
```

### 수동 설치

1. [릴리즈 페이지](https://github.com/kimtuna/bsh/releases)에서 운영체제에 맞는 바이너리를 다운로드
2. 다운로드한 파일을 압축 해제
3. `bsh` 실행 파일을 PATH에 추가

## 사용 방법

### 회사 등록

```bash
bsh register
```

등록 과정에서 다음 정보를 입력해야 합니다:
- 회사 지갑 주소
- 회사 이름
- 대표자 이름
- 이메일
- 구독 유형 (1: 1개월, 2: 3개월, 3: 1년)
- 서버 IP
- 서버 이름
- SSH 포트

### 서버 로그인

```bash
bsh login
```

로그인 시 회사 지갑 주소만 입력하면 됩니다. 로그인 성공 시 자동으로 SSH 접속이 시작됩니다.

## 구독 가격

- 1개월: 0.000001 ETH
- 3개월: 0.0000025 ETH
- 1년: 0.000008 ETH

## 라이선스

MIT License 