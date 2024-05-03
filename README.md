# Emergecy-notification-system
---


## Start docker compose

### Start for tests

```
docker-compose -f docker-compose-test.yml up -d
```

### Start for development or production

```
docker-compose -f docker-compose.yml up -d
```

### Start both

```
docker-compose -f docker-compose.yml -f docker-compose-test.yml up -d
```