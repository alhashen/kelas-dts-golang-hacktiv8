#!/bin/sh

echo '[LOG] book with id 8 before the update'
curl -s -X GET http://localhost:2400/book/9 \
  -H 'Content-Type: application/json' \
  | json_pp

echo '[LOG] UPDATE book with id 8'
curl -s -X PUT http://localhost:2400/book/9 \
  -H 'Content-Type: application/json' \
  -d '{"description":"Philosophical Fiction"}' \
  | json_pp



