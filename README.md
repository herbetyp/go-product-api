![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
<!-- ![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white) -->
<!-- ![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white) -->

# Go Product API

> API for product inventory management.

### Developed with:
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Gin Web Framwork](https://gin-gonic.com/)
- [GORM](https://gorm.io/index.html)
<!-- - [Redis](https://redis.io/) -->
<!-- - [NGnix](https://nginx.org/) -->

### API Features:

> Auth
- [x] OAuth Authentication Endpoint (*_grant_type=client_credentials_*)
- [x] Authorization from **JWT Bearer Token**
> User
- [x] Create User
- [x] Get User
- [x] List Users
- [x] Delete User (*_soft delete_*, *_hard delete_*)
- [x] Update User (*_username, password_*)
- [x] Recovery User
> Products
- [x] Create Product
- [x] Get Product
- [x] List Products
- [x] Update Product
- [x] Delete Product (*_soft delete_*, *_hard delete_*)
- [x] Recovery Product
> Application systems
- [ ] **JSON format log** system
- [ ] **Cache** system
- [x] Auto **Migrations** system
<!-- - [ ] **NGnix** proxy System -->
---

### Run localy application:
Pre commit (For development)
```bash
pip intall pre-commit # install pre-commit with python pip
go install golang.org/x/tools/cmd/goimports@latest # pre-commit hook
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4 # pre-commit hook
pre-commit install # install pre-commit hooks
```

Docker *__with docker compose module__*
```bash
docker compose up -d # exposed in port 3000
```
Runner Local Server *__Gin__*
```bash
export GIN_MODE="test" # set gin mode
docker compose up -d product_api_go_db product_api_go_cache # run database/cache container
go run cmd/main.go # exposed in port 5000
```

<!-- ### Architecture Diagram

![Architecture](./docs/img/architecture_diagram.png) -->
