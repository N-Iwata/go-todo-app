## コンテナ起動

```
docker-compose up -d
```

## コンテナの中に入る

```
docker exec -it postgres bash
or
docker-compose exec db bash
```

## postgresのtestdbの中に入る

```
psql todoapp -U admin 
```

## postgresから出る

```
\q
```

## 

```
\d
```

## コンテナから出る

```
exit
```

## コンテナ停止

```
docker-compose down
```


