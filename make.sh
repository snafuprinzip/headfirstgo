#!/bin/sh
GOPATH=$HOME/go/src/github.com/snafuprinzip/headfirstgo/

for srcfile in src/*/*.go; do 
	/bin/echo -n "Building $srcfile ->"
	go build -o bin $srcfile
	echo " bin/"$(basename $srcfile | cut -d. -f1) 
done
