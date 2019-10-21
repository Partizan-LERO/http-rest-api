RUN migrations:

`docker exec -ti deployments_app_1  migrate -path migrations -database "postgres://postgres-dev:s3cr3tp4ssw0rd@172.17.0.1:5435/dev?sslmode=disable" up`

CREATE a migration:
`docker exec -ti deployments_app_1  migrate create -ext sql -dir migrations/ [migration_name] 
`