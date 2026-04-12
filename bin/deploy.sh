#!/bin/bash

set -eu

RETRY_MARKER="dist/.deploy-retry"

OLD_MD5=""
if ls dist/*.xml 1>/dev/null 2>&1; then
    OLD_MD5=$(md5sum dist/*.xml | md5sum | cut -d ' ' -f 1)
fi

make build

NEW_MD5=$(md5sum dist/*.xml | md5sum | cut -d ' ' -f 1)

if [ "$OLD_MD5" = "$NEW_MD5" ] && [ ! -f "$RETRY_MARKER" ]; then
    echo "No changes detected, skipping deploy."
    exit 0
fi

rm -f "$RETRY_MARKER"

RETRY=1

while true; do
    if [ $RETRY -gt 3 ]; then
        touch "$RETRY_MARKER"
        exit 1
    fi

    echo "Trying $RETRY time(s)..."

    if npx wrangler deploy; then
        break
    fi

    ((RETRY++))
    sleep 10
done
