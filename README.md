# IoT-parking

Run the following line to start database:

```code
docker run --rm   --name pg-parking -e POSTGRES_PASSWORD=admin -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
```

```code
docker run --rm --name test -e DB_DBNAME=parking -e DB_DRIVER=postgres -e DB_HOST=172.17.0.2 -e DB_PASSWORD=admin -e DB_PORT='5432' -e DB_TZ=Asia/Tehran -e DB_USER=postgres -e ENVIRONMENT=DEVELOPMENT -e LOG_CONSOLE=DEBUG -e LOG_EXTRA_APPNAME='true' -e LOG_EXTRA_BRANCH='true' -e LOG_EXTRA_COMMIT='true' -e LOG_EXTRA_DATA='true' -e LOG_LEVEL=DEBUG -e MONITORING='false' -e REST='true' -e REST_LOG='true' -e REST_MIDDLEWARE_BODY_LIMIT='' -e REST_MIDDLEWARE_CORE='true' -e REST_MIDDLEWARE_GZIP='false' -e REST_MIDDLEWARE_RECOVER='true' -e REST_MIDDLEWARE_REMOVE_TRAILING_SLASH='true' -e REST_PORT='7676' -e REST_VALIDATOR='true' -e DEBUG='true' -e API_KEY='user' -e ADMIN_API_KEY='admin' farazff/parking:1.1
```
