# IoT-parking

Run the following line to start database:

```code
docker run --rm   --name pg-parking -e POSTGRES_PASSWORD=admin -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
```
