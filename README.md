<div align="center">

# Go Product API :globe_with_meridians:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
<!-- ![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white) -->

API for product management, with user system, authentication and authorization.

</div>

## :wrench: Run localy application:

1) #### Copy and export variables (*__Run in root project__*)
```bash
cp -r docs/samples/*.sample $PWD; for i in *.sample ; do mv "$i" "$(basename "$i" .sample)" ; done && source envs.sh
```
> [!NOTE]
> When you restart the terminal session, you will need to run the `source envs.sh` command again.

2) Genereta RSA private key for assign JWT token
```bash
openssl genpkey -out rsa_pvkey.pem -algorithm RSA -pkeyopt rsa_keygen_bits:2048
```

3) #### Run application in Docker container (with Docker Composer module) or directly from terminal
```bash
# using docker container, running on http://localhost:3000
docker compose up -d
```

```bash
# using direct in terminal, running on http://localhost:5000
docker compose up -d go_product_api_db go_product_api_cache
go run cmd/main.go
```

4) #### Create user:
```bash
curl --request POST \
  --url https://go-product-api.onrender.com/v1/users \
  --header 'Content-Type: application/json' \
  --data '{
	"username": "Anonymous1",
	"email": "anonymous1@email.com",
	"password": "mypassword123"
}'
```

> [!IMPORTANT]
> By default, users are created with `active` set to `false`, run the following command to make the `admin` user active and the system unlock:

```bash
docker exec -ti $DB_CONTAINER_NAME psql -U $DB_USER \
-d $DB_NAME -c "UPDATE public.users SET active = true, is_admin = true WHERE id = 1;"
```

---
- Pre commit (*__For development only__*)
```bash
pip intall pre-commit # install pre-commit with python pip
go install golang.org/x/tools/cmd/goimports@latest # pre-commit hook
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4 # pre-commit hook
pre-commit install # install pre-commit hooks
```

## :dart: API Features:

| `Feature` |  `User` | `Admin` |
| --- | :---: | :---: |
| **OAuth Basic** Authentication Endpoint | :white_check_mark: | :white_check_mark: |
| Authorization from **JWT Bearer Token** | :white_check_mark: | :white_check_mark: |
| **Create User** | :white_check_mark: | :white_check_mark: |
| **Get User** | :white_check_mark: | :white_check_mark: |
| **Update User** (*_username, password_*) | :white_check_mark: | :white_check_mark: |
| **List Users** | :x: | :white_check_mark: |
| **Delete User** (*_soft delete, hard delete_*) | :x: | :white_check_mark: |
| **Recovery User** | :x: | :white_check_mark: |
| **Active/Deactive User** | :x: | :white_check_mark: |
| **Create Product** | :white_check_mark: | :white_check_mark: |
| **Get Product** | :white_check_mark: | :white_check_mark: |
| **List Products** | :white_check_mark: | :white_check_mark: |
| **Update Product** | :white_check_mark: | :white_check_mark: |
| **Delete Product** (*_soft delete, hard delete_*) | :x: | :white_check_mark: |

### ⚙️ Application systems
- [x] **JSON format output log**
- [x] **Memory in cache (Redis)**
- [x] **Rate limiting by IP**
- [x] **Database auto migrations**
<!-- - [ ] **NGnix** proxy System -->

### :space_invader: Developed with:
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/index.html)
- [Zap Logger](https://github.com/uber-go/zap)
- [Redis](https://github.com/redis/go-redis)
<!-- - [NGnix](https://nginx.org/) -->
