#!/bin/bash

dirs=($(ls pkg))

for i in ${dirs[@]};
do
    if [[ ${i} != "utils" ]];then
        go test ./pkg/${i} $1
    fi
done
