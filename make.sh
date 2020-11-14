#!/bin/sh

for srcfile in src/*.go; do 
	/bin/echo -n "Building $srcfile ->"
	go build -o bin $srcfile
	echo " bin/"$(basename $srcfile | cut -d. -f1) 
done

