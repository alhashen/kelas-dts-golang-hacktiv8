#!/bin/sh

url=http://localhost:2400/book 
valid_id=18
invalid_id=100
bad_params=vy

echo 'GET valid book id'
curl -s -X GET $url/$valid_id \
  -H 'Content-Type: application/json' \
  | json_pp

echo 'GET invalid book id'
curl -s -X GET $url/$invalid_id \
  -H 'Content-Type: application/json' \
  | json_pp

echo 'GET invalid book id with bad params'
curl -s -X GET $url/$bad_params \
  -H 'Content-Type: application/json' \
  | json_pp
