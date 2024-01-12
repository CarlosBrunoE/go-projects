# 3 Rest API with Gin and Postgresql

Reference: https://www.twilio.com/blog/build-restful-api-using-golang-and-gin

REST API using Gin that connects to Postgresql database. Authentication is based on JWT.

## Prerrequisites
- Docker
- Docker compose
- Golang 
- curl
- .env file

``` 
### .env.local file:

# Database credentials
DB_HOST="localhost"
DB_USER="postgres"
DB_PASSWORD="superpassword123"
DB_NAME="gin-rest-api"
DB_PORT="5432"

# Authentication credentials
TOKEN_TTL="2000"
JWT_PRIVATE_KEY=""

## Steps:

1. First, we need our database. To set it up run the following command at root of the project:

```bash
docker-compose up -d
```

2. Run project with the following command:

```
go run .
```

To test API run:

3. Create our user:

```bash
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"<<USERNAME>>", "password":"<<PASSWORD>>"}' \
    http://localhost:8000/auth/register
```

4. Login:

```bash
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"<<USERNAME>>", "password":"<<PASSWORD>>"}' \
    http://localhost:8000/auth/login
```
The login will retrieve our JWT token. It will be used to authenticate us to our private endpoints.

5. Create new entry:

```bash
curl -d '{"content":"Set here the content you want to register."}' \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer <<JWT>>" \
    -X POST http://localhost:8000/api/entry
```


6. Retrieve list of entries:

This just going to show us the entries that belongs to our current user.

```bash
curl -H "Content-Type: application/json" -H "Authorization: Bearer <<JWT>>" -X GET http://localhost:8000/api/entry
```