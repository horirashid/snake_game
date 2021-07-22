#! /bin/bash
clear
stty cbreak min 1
stty -echo
go run main.go inkey.go
stty echo