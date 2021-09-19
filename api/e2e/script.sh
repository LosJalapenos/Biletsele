#!/bin/bash

# This script can be used to automatically test the API's functionalities
# A more suitable alternative: https://www.postman.com/automated-testing/

ENDPOINT=localhost:8080
COOKIES_FILE=cookies.txt

call () {
  curl "${ENDPOINT}/$1" "${@:2}" -"$(eval [[ "$1" == "connect" ]] && echo "c" || echo "b")" ${COOKIES_FILE}
}

connect() {
  call "connect" -H "Content-Type: application/json" -d "{\"name\":\"$1\"}"
}

create_room() {
  call "create_room"
}

connect LosJalapenos
create_room

rm ${COOKIES_FILE}
