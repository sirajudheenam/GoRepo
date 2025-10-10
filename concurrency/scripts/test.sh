#!/bin/bash

# Test the channel manager server with Apache Bench (ab)
# Make sure the server is running before executing this script
# Usage: ./test.sh

# test if 127.0.0.1:8000 is reachable
if !nc -zv 127.0.0.1 8000 ; then
    echo "Error: Channel manager server is not reachable"
    echo "Make sure the server is running on 127.0.0.1:8000"
    exit 1
fi

# -n Number of requests to perform
# -c Number of multiple requests to make at a time (Concurrency level)
# -s Timeout for each request in seconds (optional)
# Example: 10 requests with a concurrency level of 10
echo "Testing the channel manager server with Apache Bench..."
echo
echo "Incrementing counter 'i' 100 times with concurrency level 10"
echo
ab -n 100 -c 10 -s 2 "127.0.0.1:8000/inc?name=i"
echo
echo "Getting counter 'j' 100 times with concurrency level 10"
echo
ab -n 100 -c 10 -s 2 "127.0.0.1:8000/get?name=j"
echo
echo "Setting counter 'j' to 25, 100 times with concurrency level 10"
echo
ab -n 100 -c 10 -s 2 "127.0.0.1:8000/set?name=j&val=25"
