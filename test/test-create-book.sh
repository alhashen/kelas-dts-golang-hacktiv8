#!/bin/sh

curl -s -X POST http://localhost:2400/book \
  -H 'Content-Type: application/json' \
  -d '{"book_title":"The Brothers Karamazov", "book_author":"Fyodor Dostoevsky", "description":"Fiction"}' \
  | json_pp
