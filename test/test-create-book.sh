#!/bin/sh

echo 'CREATE 3 books'
url=http://localhost:2400/book 

curl -s -X POST $url \
  -H 'Content-Type: application/json' \
  -d '{"book_title":"The Brothers Karamazov", "book_author":"Fyodor Dostoevsky", "description":"Fiction"}' \
  | json_pp

curl -s -X POST $url \
  -H 'Content-Type: application/json' \
  -d '{"book_title":"Master and Margarita", "book_author":"Mikhail Bulgakov", "description":"Fiction"}' \
  | json_pp

curl -s -X POST $url \
  -H 'Content-Type: application/json' \
  -d '{"book_title":"The Cyberiad", "book_author":"Stanislaw Lem", "description":"Science Fiction"}'\
  | json_pp

curl -s -X POST $url \
  -H 'Content-Type: application/json' \
  -d '{"book_title": "asijsi"}' \
  | json_pp
