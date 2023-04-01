#!/bin/sh

echo 'GET valid book id'
curl -s -X GET http://localhost:2400/book/6 \
  -H 'Content-Type: application/json' \
  | json_pp

echo 'GET invalid book id'
curl -s -X GET http://localhost:2400/book/1 \
  -H 'Content-Type: application/json' \
  | json_pp
