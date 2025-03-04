# About Project

<i style="color: gray;">Updated 2025-03-05</i>

이 프로젝트는 `cobra` 기반의 CLI 애플리케이션입니다.
프로젝트 초기 구축과 CI/CD yaml 자동화 생성을 돕는 프로그램입니다.

---

추가적인 커스텀 기능 넣을 예정입니다.

1. custom input
2. ci/cd path
3. ci/cd customized jobs

---

<br>

# Architecture

```bash
dev-quickstart/
├── cmd/
│   ├── root.go       ✅ (CLI 명령어 관련 코드만 유지)
│   ├── frontend.go   # FE Project Setting
│   ├── backend.go    # BE ..
│   ├── ci_cd.go      # CI/CD Pipeline
├── utils/
│   ├── utils.go
├── main.go           ✅ (루트로 이동)
├── go.mod

```

# How to run

프로젝트를 설정하려면 다음 명령어를 실행하세요:

```bash
go run main.go setup
```
