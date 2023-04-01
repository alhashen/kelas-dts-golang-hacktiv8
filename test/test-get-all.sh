#!/bin/sh

curl -s -X GET http://localhost:2400/book/all \
  -H 'Content-Type: application/json' \
  | json_pp
