#!/bin/bash
while read -r line; do
    a=( $line )
    l=${#a[*]}
    n=${a[$(($l-1))]}
    if [ $n -lt $l ]; then
        echo ${a[$(($l-1-$n))]}
    fi
done < $1
