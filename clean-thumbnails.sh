#!/bin/bash

for d in $(find ./galleries/ -maxdepth 1 -type d)
do
    rm -rf $d/thumbs
done
