#!/bin/sh

action=$1

case $action in
"install")
    go install github.com/swaggo/swag/cmd/swag@latest
    ;;
"init")
    swag init -g swagger.go
    ;;
"format")
    swag fmt
    ;;
*)
    exit 0
    ;;
esac