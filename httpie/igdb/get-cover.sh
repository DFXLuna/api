#!/bin/bash
set -e

set -o allexport
source ../.env
source .token
set +o allexport

poetry run http --ignore-stdin --raw 'fields *; where game = 1942;' -A bearer -a ${TOKEN} \
    POST https://api.igdb.com/v4/covers/ \
    Client-ID:${IGDB_CLIENT_ID} \
    Accept:application/json
