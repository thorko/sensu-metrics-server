#!/bin/sh

url=$1

out=$(./curl -s $1)

d=$(date +%s)

IFS=$'\n'
for line in $out; do
	a=$(echo $line | ./sed 's/{/,/g' | ./sed 's/[}"]//g')
	echo $a | ./awk -v date="$d" '{print $1" value="$2" "date}'
done
