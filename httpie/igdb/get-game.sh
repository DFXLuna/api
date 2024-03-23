#!/bin/bash
set -e

set -o allexport
source ../.env
source .token
set +o allexport

poetry run http --ignore-stdin --raw 'fields *; where id = 1942;' -A bearer -a ${TOKEN} \
    POST https://api.igdb.com/v4/games/ \
    Client-ID:${IGDB_CLIENT_ID} \
    Accept:application/json
