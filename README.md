# Project

## 介紹
從 shortURL 開始，之後會陸續加入一些我想加入的功能，可能包含前後端，to-do list，學習 vue，react，kotlin，rust 等等...
同步整理部分內容在 https://hackmd.io/0p7KCTOiQ_6QwX_4UAYLIQ

## 如何執行
- 直接開發
    - 前端 `cd ./frontend && npm run serve`
    - 後端 `cd ./backend && go run main.go`
- Docker-Compose 開發
    - `make` 編譯前端和後端成 docker image
    - `make up` 將 docker images 跑起來
    - `make up db` 只將 db images 跑起來
    - `make down` 將 docker images 關起來

## Redis
- docker exec -it projecturlshorten-redis-1 redis-cli
- docker exec -it projecturlshorten-redis-1 redis-benchmark -n 1000000 -t set,get -q

## Mysql
- mysql --port=3306 -h0.0.0.0 -umyuser -pmypassword
- sysbench /usr/share/sysbench/oltp_read_write.lua --mysql-host=0.0.0.0 --mysql-port=3306 --mysql-user=myuser --mysql-password=mypassword --mysql-db=mydatabase --tables=10 --table-size=1000000 --events=100000000 --report-interval=10 --threads=4 --time=300 prepare
- sysbench /usr/share/sysbench/oltp_read_write.lua --mysql-host=0.0.0.0 --mysql-port=3306 --mysql-user=myuser --mysql-password=mypassword --mysql-db=mydatabase --tables=10 --table-size=1000000 --events=100000000 --report-interval=10 --threads=4 --time=300 run