#!/bin/bash

set -eu

FILE="$1"
URL="$2"

curl \
    --etag-compare "$FILE.etag" \
    --etag-save "$FILE.etag" \
    --output "$FILE" \
    --user-agent "Wget/1.25.0" \
    "$URL"
