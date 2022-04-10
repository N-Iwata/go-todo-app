## postgresコンテナ起動

```
docker-compose up -d
```

## postgresコンテナの中に入る

```
docker exec -it postgres bash
or
docker-compose exec db bash
```

## postgresのtodoappの中に入る

```
psql todoapp -U admin 
```

## postgresから出る

```
\q
```

## コンテナから出る

```
exit
```

## コンテナ停止

```
docker-compose down
```


