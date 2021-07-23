#! /bin/bash
reset
stty cbreak min 1
stty -echo
go run main.go inkey.go snake.go queue.go
stty echo