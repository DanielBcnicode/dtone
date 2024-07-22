# DTOnetest
## Presentation
## Services

This microservice expose an API in 8080 port with the next characteristics:

    - Languaje used : Golang.
    - JWT security.
    - Gin gonic for HTTP API.
    - Gorm for db repositories.
    - PostgreSql as db manager.
    - Open Telemetry as standard telemetry system.
    - Dockerised solution.
    - Swagger service.
    - Github Actions as CI system
    - .env file for configuration data
    - SigNoz for visualize API Service and Traces (External installation)

## Installation
### Requirements
    - Git
    - Docker / Docker-compose
### Clone the DTOnetest repository
```
$ git clone https://github.com/DanielBcnicode/dtone.git
```
### Setup
#### First time

The first time you run this API is mandatory to create the database schema and tables

```
$ docker compose up -d
```
Update db structure
```
docker compose exec test ./dtone_migrate
```
Re-start the service
```
$ docker compose down
```
and finally ...

#### Next time
```
$ docker compose up -d
```

#### Installing SigNoz (optional)
In some other folder execute:
```
git clone -b main https://github.com/SigNoz/signoz.git
```
and then:
``` 
docker compose -f signoz/deploy/docker/clickhouse-setup/docker-compose.yaml up -d
```
you must wait until the application loads then you can go to
`http://localhost:3301` and open the UI to see the traces.
You must register previously (only in local)

## Execution
The Swagger page is located in `http://localhost:8080/api/v1/swagger/index.html`
Steps:
- Register an user: `/api/v1/register`
- Login the user: `/api/v1/login`

Response:
```
{
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InR3b0B0b3cuY29tIiwiZXhwIjoxNzIxNjYxNzEyLCJ1c2VyX2lkIjoiODE0N2QyNmEtZWQ5MS00ZjZkLWFhZmMtOWNkMTI2Mjk1NjY1In0.T46gsB4lZndfK6IUDgKQ2IdHSaxRZNbwK8HTptmJU88"
}
```

  - Use the jwt token in the next calls:
    - Get the jwt value and put it using the Authorize button
    - You must put : bearer xxxxxxxxxxxxxxxxxxxxxxxxx in the text field and press Authorize
  - TopUp the balance of the user `/api/v1/users/{user_id}/topup`
    - user_id is in uuid4 format, and you can get it in the register response
    - amount is in text format with 2 decimals. Ex:  "1000.87"
  - Create products
  - Upload files to products
  - Buy products from other users
  - Gift products to other users
  - Get user transactions
## Improvements 

