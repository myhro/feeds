#!/bin/bash

set -eu

FILE="dist/autossegredos.xml"

OLD_MD5=""
if [ -f "$FILE" ]; then
    OLD_MD5=$(md5sum "$FILE" | cut -d ' ' -f 1)
fi

make build

NEW_MD5=$(md5sum "$FILE" | cut -d ' ' -f 1)

if [ "$OLD_MD5" = "$NEW_MD5" ]; then
    echo "No changes detected, skipping deploy."
    exit 0
fi

RETRY=1

while true; do
    if [ $RETRY -gt 3 ]; then
        make clean
        exit 1
    fi

    echo "Trying $RETRY time(s)..."

    if npx wrangler deploy
    then
        break
    fi

    ((RETRY++))
    sleep 10
done
