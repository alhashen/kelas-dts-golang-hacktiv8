#!/bin/sh

echo 'DELETE valid book id'
curl -s -X DELETE http://localhost:2400/book/6 \
  -H 'Content-Type: application/json' \
  | json_pp

echo 'DELETE invalid book id'
curl -s -X DELETE http://localhost:2400/book/2 \
  -H 'Content-Type: application/json' \
  | json_pp
