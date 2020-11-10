#!/bin/bash

echo -n "Checking vm.max_map_count settings... ";

max_map_count=`sysctl -n vm.max_map_count`

if [ "$max_map_count" -lt "262000" ];then
    echo "fail";
    echo "Elasticsearch uses a mmapfs directory by default to store its indices. Limits on mmap counts is likely to be too low, which may result in out of memory exceptions.";
    echo "See https://www.elastic.co/guide/en/elasticsearch/reference/current/vm-max-map-count.html.";
    exit 1
else
    echo "ok";
fi
