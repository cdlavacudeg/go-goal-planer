# Goal planner REST API

## How to use

```
git clone https://github.com/cdlavacudeg/go-goal-planner.git
```

```
make
```

```
./bin/goal-planner
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
aws dynamodb delete-table --table-name GoalPlanner --endpoint-url http://localhost:8000
aws dynamodb describe-table --table-name GoalPlanner --endpoint-url http://localhost:8000
aws dynamodb get-item \
  --table-name GoalPlanner \
  --key '{"PK": {"S": "User#100"}, "SK": {"S": "User#100-t100"}}' \
  --endpoint-url http://localhost:8000

```

## DynamoDB one table design

| Entity    | PK               | SK                      |
| --------- | ---------------- | ----------------------- |
| User      | USER#<Username>  | USER#<Username>         |
| Vision    | UV#<Username>    | VISION#<VisionId>       |
| Objective | VO#<VisionId>    | OBJECTIVE#<ObjectiveId> |
| Tasks     | OT#<ObjectiveId> | TASK#<TaskId>           |
