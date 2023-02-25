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
    - `make up` 將 docker image 跑起來
    - `make down` 將 docker image 關起來