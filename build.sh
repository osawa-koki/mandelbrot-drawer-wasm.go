#!/bin/bash

### Go側
# JSでWebAssemblyを実行するためのファイルをコピー。
cat $(go env GOROOT)/misc/wasm/wasm_exec.js > ./public/wasm_exec.js
# WebAssemblyをビルドする。
GOOS=js GOARCH=wasm go build -o ./public/main.wasm  .

cp ./index.html ./public/index.html
