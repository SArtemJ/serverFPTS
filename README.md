### serverFPTS

> Run with docker:
- `docker-compose up db` to build clean db container
- `docker-compose up init_db` to init db schema
- `docker-compose up data_db` to fill db test data
- `docker-compose up fpts` to run webserver
- `docker-compose up drop_db` to remove data and schema

> `Available commands`
````
./fpts db init  (init schema)
./fpts db data  (fill db test data)
./fpts db drop  (cleanup db data and schema)
````

> `Avaliable flags`
````
      --cleaner              use background cleaner transactions - (default true)
  -c, --config string        Config file path
      --db.host string       Database server host
      --db.login string      Database login
      --db.name string       Database name
      --db.pass string       Database pass
      --db.port int          Database server port (default 5432)
  -h, --help                 help for fpts
      --http.host string     API host
      --http.port int        API port
      --limit.clean int      limit last transactions -(default 3)
      --log.fmt string       log format json or text
      --log.level string     log level (default "DEBUG")
      --logger.file string   stdout or file (default "STDOUT")
      --timer.period int     timer period - (default 5)
      --timer.type string    s-seconds/m-minutes/h-hours - (default "m")

````

> `Send post Request on webserver`
- `userGUID` - required param for identifier user wallet
- url by default `http://localhost:8099/transaction`
````
curl --location --request POST 'http://localhost:8099/transaction' \
  --header 'Source-Type: server' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  	"state": "win", 
  	"amount": "10.15", 
  	"transactionId": "144de7a7-9222-43d0-90d8-ac48630af3ad",
  	"userGUID": "aa4d02cb-2ca9-4e34-8b39-5e49559c1136"
  }'
````

> `For local build`
- `go generate`
- `go build -o fpts .`