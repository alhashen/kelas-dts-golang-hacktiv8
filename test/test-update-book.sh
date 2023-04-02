#!/bin/sh

url=http://localhost:2400/book 
valid_id=18
invalid_id=100

echo '[LOG] book with id '$valid_id' before the update'
curl -s -X GET $url/$valid_id \
  -H 'Content-Type: application/json' \
  | json_pp

echo '[LOG] book with id '$valid_id' after the update'
curl -s -X PUT $url/$valid_id \
  -H 'Content-Type: application/json' \
  -d '{"description":"Humor, Satire, Science Fiction"}'\
  | json_pp

echo '[LOG] UPDATE book with invalid id'
curl -s -X PUT $url/$invalid_id \
  -H 'Content-Type: application/json' \
  -d '{"description":"Philosophical Fiction"}'\
  | json_pp


