#!/bin/sh

url=http://localhost:2400/book 
valid_id=16
invalid_id=100


echo 'DELETE book with valid id'
curl -s -X DELETE $url/$valid_id \
  -H 'Content-Type: application/json' \
  | json_pp

echo 'DELETE book with invalid id'
curl -s -X DELETE $url/$invalid_id \
  -H 'Content-Type: application/json' \
  | json_pp
