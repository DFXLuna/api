#!/bin/bash
set -e

set -o allexport
source ../.env
set +o allexport

json=$(poetry run http --ignore-stdin POST https://id.twitch.tv/oauth2/token \
    client_id==${IGDB_CLIENT_ID} \
    client_secret==${IGDB_CLIENT_SECRET} \
    grant_type==client_credentials)


token=$(echo ${json} | jq '.access_token')
echo TOKEN=${token} > .token
