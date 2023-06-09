# IoT-parking

There are 3 different ways to run the project. 

1.First way is to run to code locally and run the database and the redis using docker.
So, for staring the database container, we need to run the following command:

```code
docker run --rm --name parking-db -e POSTGRES_PASSWORD=admin -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
```
After that we need to start the redis container with the following command:
```code
docker run --rm --name parking-redis -d -p 6379:6379 redis
```

Then these configs must be set in .env file to run locally:
```code
PARKING_REDIS_HOST=localhost
PARKING_REDIS_PORT='6379'
PARKING_DB_HOST=localhost
PARKING_DB_PASSWORD=admin
PARKING_DB_PORT='5432'
```

```code
docker run --rm --name test -e DB_DBNAME=parking -e DB_DRIVER=postgres -e DB_HOST=172.17.0.2 -e DB_PASSWORD=admin -e DB_PORT='5432' -e DB_TZ=Asia/Tehran -e DB_USER=postgres -e ENVIRONMENT=DEVELOPMENT -e LOG_CONSOLE=DEBUG -e LOG_EXTRA_APPNAME='true' -e LOG_EXTRA_BRANCH='true' -e LOG_EXTRA_COMMIT='true' -e LOG_EXTRA_DATA='true' -e LOG_LEVEL=DEBUG -e MONITORING='false' -e REST='true' -e REST_LOG='true' -e REST_MIDDLEWARE_BODY_LIMIT='' -e REST_MIDDLEWARE_CORE='true' -e REST_MIDDLEWARE_GZIP='false' -e REST_MIDDLEWARE_RECOVER='true' -e REST_MIDDLEWARE_REMOVE_TRAILING_SLASH='true' -e REST_PORT='8080' -e REST_VALIDATOR='true' -e DEBUG='true' -e API_KEY='user' -e ADMIN_API_KEY='admin' -p 8080:8080 farazff/parking:1.1
```
