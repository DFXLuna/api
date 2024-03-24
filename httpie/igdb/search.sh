#!/bin/bash
set -e

set -o allexport
source ../.env
source .token
set +o allexport

http --ignore-stdin --raw 'search "Halo"; fields name,id;' -A bearer -a ${TOKEN} \
    POST https://api.igdb.com/v4/games/ \
    Client-ID:${IGDB_CLIENT_ID} \
    Accept:application/json