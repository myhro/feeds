#!/bin/bash

set -eu

RETRY=1

while true; do
    if [ $RETRY -gt 3 ]; then
        exit 1
    fi

    echo "Trying $RETRY time(s)..."

    if npx wrangler pages deploy --branch "$BRANCH" --project-name feeds dist/
    then
        break
    fi

    ((RETRY++))
    sleep 10
done
