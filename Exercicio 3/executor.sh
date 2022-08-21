#!/bin/bash
CONTADOR=0
while [ $CONTADOR -lt $1 ]; do
	if [ $CONTADOR -eq 0 ]; then
		go run cliente.go $2 > teste$1$2.txt & 
	else
		go run cliente.go $2 & 
	fi
    
    let CONTADOR=CONTADOR+1; 
done