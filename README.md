

# Go Product API
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
<!-- ![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white) -->

> API for product inventory management, with user and admin user permission system, also with authentication and authorization.

### Developed with:
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/index.html)
- [Zap Logger](https://github.com/uber-go/zap)
- [Redis](https://github.com/redis/go-redis)
<!-- - [NGnix](https://nginx.org/) -->

## API Features:

> Auth
- [x] OAuth Authentication Endpoint (*_grant_type=client_credentials_*)
- [x] Authorization from **JWT Bearer Token**
> **All Users**
- [x] Create User
- [x] Get User
- [x] Update User (*_username, password_*)
> **Admin Only**
- [x] List Users
- [x] Delete User (*_soft delete, hard delete_*)
- [x] Recovery User
- [x] Active/Deactive User
- [x] Delete Product (*_soft delete_*, *_hard delete_*)
- [x] Recovery Product
> **All Users** Products
- [x] Create Product
- [x] Get Product
- [x] List Products
- [x] Update Product
> **Application systems**
- [x] **JSON format log** system
- [ ] **Cache** system
- [x] **Rate limit** system
- [x] Auto **Migrations** system
<!-- - [ ] **NGnix** proxy System -->

## Run localy application:
Copy and export variables (*__runner in root project__*)
```bash
cp docs/samples/* .; for file in *.sample; do cp -r "$file" "${file%.sample}"; done && rm *.sample && source envs.sh
```

Docker *__with docker compose module__*
```bash
docker compose up -d # exposed in port 3000
```

Runner Local Server *__Gin__*
```bash
docker compose up -d go_product_api_db go_product_api_cache
go run cmd/main.go # exposed in port 5000
```

Of default as users is created with `active` and `is_admin` as `false`, execute the following command to active the admin user **after crated first user using the API**:
```bash
docker exec -ti $DB_CONTAINER_NAME psql -U $DB_USER \
-d $DB_NAME -c "UPDATE public.users SET active = true, is_admin = true WHERE id = 1;"
```
> **Note:** With admin, calling endpoint `/v1/admin/users/<user-id>/status?active=<true|false>` is possible **enable** or **disable** others users.

Pre commit (*__For development__*)
```bash
pip intall pre-commit # install pre-commit with python pip
go install golang.org/x/tools/cmd/goimports@latest # pre-commit hook
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4 # pre-commit hook
pre-commit install # install pre-commit hooks
```
