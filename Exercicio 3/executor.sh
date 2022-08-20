#!/bin/bash
for i in $1; do
	if (i==0) then
		go run cliente.go $2 > teste$1.txt & 
	else
		go run cliente.go $2 & 
	fi
done
