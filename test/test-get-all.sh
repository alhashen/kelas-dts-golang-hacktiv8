#!/bin/sh

url=http://localhost:2400/book/all

curl -s -X GET $url \
  -H 'Content-Type: application/json' \
  | json_pp
