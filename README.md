# Goal planer REST API

## How to use

```
git clone https://github.com/goalplaner/goalplaner
```

```
make
```

```
./bin/goal-planer
```

## Dynamodb Local

### [Setting up aws cli](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

- [Install](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [Configure](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)

### [Docker image](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html#docker)

### Test

```
docker compose up
aws dynamodb list-tables --endpoint-url http://localhost:8000
```
