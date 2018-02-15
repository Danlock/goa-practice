# Steps:

1. Download go's official dep tool. https://github.com/golang/dep/releases
2. run make

Note that if main.go is deleted and recreated, make will fail because main.go is mostly there as a guide and doesnt generate the proper import paths.
Note that because dep ensure happens before the code is generated, all of the projects dependencies will only become part of the lockfile after make is run twice. This is only required for initial repositories.


# Environment:
1. Install docker and docker-compose
2. docker compose build requires a docker/.env file providing required      environment variables outlined in the compose file.
3. docker-compose up -d

To locally check out the API:
0. Run the environment in the background
1. run ` make && go run main.go `
2. open new terminal and run ` go run generated/tool/practice-cli/main.go ` to check out the cli client for interacting with the API
3. Check out the docs hosted at localhost:8080/docs to see what can be done

## Tips:
Easy way of logging into to mongo container via CLI

`docker exec -it docker_mongodb_1 bash -c  'mongo -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" admin'`

Easy way of looking at status of RabbitMQ:
` http://localhost:15672 `

Try hosting a pub sub session with the API by creating an asset, subscribing to an assets events in a different terminal window and publishing events to the first window.
This isn't really meant to work with any frontend, just mostly a way for me to learn about using Golang, MongoDB and RabbitMQ and the Goa framework.
