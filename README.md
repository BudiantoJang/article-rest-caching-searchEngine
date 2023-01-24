# Jang's Article Back-End Application
This is a simple back-end system for article application that allows users to create article and search for article. Implementing Redis and Redisearch makes the application faster, easier, and more efficient to use

### Prerequisites
- Go 1.18
- Docker
- Direnv 

## Steps
### Creating Database
1. pull and run docker image using this command:

```
docker run -e POSTGRES_PASSWORD=mysecret -p 5431:5432 -d postgres
```

the reason why I used port 5431 is to prevent conflict while having postgres running in local machine

2. Create database table using this image as a reference:

![Database Table Schema](https://drive.google.com/uc?export=view&id=1brNUJ4-IldS-wPRlIyjT1MBfeO-jGR-o)

3. Populate the data using the data you desires


### Creating Redis and Redisearch
1. pull and run docker image using this command:

```
docker run -p 6378:6379 -d redis/redis-stack:latest
```
the reason why I used port 6378 is to prevent conflict while having redis/redisearch/redis-stack-server running in local machine


### Adding Credentials
1. Create a `.envrc` file in the root directory of the project and fill it with the following environment variables:

```
export REDIS_HOST="localhost:6378"
export REDIS_PASSWORD=""
export REDIS_DB="0"
export REDISEARCH_HOST="localhost:6378"
export REDISEARCH_INDEX="articleIndex"
export POSTGRES_HOST="localhost"
export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="mysecret"
export POSTGRES_DB_NAME="postgres"
export POSTGRES_PORT="5431"
```

### Running the Application
1. Run `direnv allow` to load the environment variables
2. Make sure having all the prequisites installed, run 'go mod tidy' to install all the dependencies
3. Run `go run main.go` or `make run` to run the application


## API Documentation

The API documentation for this project can be accessed from : [POSTMAN COLLECTION](https://documenter.getpostman.com/view/20605497/2s8ZDbWg1N)

## Future Update

- [ ] Add integration test
- [ ] Add more features
- [ ] integrate all the container using docker compose