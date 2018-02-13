# Steps:

1. Download go's official dep tool. https://github.com/golang/dep/releases
2. run make

Note that if main.go is deleted and recreated, make will fail because main.go is mostly there as a guide and doesnt generate the proper import paths.
Note that because dep ensure happens before the code is generated, all of the projects dependencies will only become part of the lockfile after make is run twice. This is only required for initial repositories.

# Environment:
1. Install docker and docker-compose
2. docker compose build requires a docker/.env file providing required      environment variables outlined in the compose file.
3. docker-compose up -d

## Tips:

Easy way of logging into to mongo container via CLI

`docker exec -it docker_mongodb_1 bash -c  'mongo -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" admin'`
