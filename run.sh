#! /bin/bash
echo GAME START
echo GOGOGO!!
clear
reset
clear
stty cbreak min 1
stty -echo
go run main.go inkey.go snake.go queue.go map.go game.go node.go saver.go score.go
stty echo