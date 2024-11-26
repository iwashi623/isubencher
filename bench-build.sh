#!/bin/bash

ISUCON_NAME=${ISUCON_NAME}

if [ "$ISUCON_NAME" = "kayac-listen80" ]; then
    make init
elif [ "$ISUCON_NAME" = "kayac-listen90" ]; then
    TARGET_STAGE="kayac-listen90"
else
    echo "Error: Unknown ISUCON_NAME: $ISUCON_NAME"
    exit 1
fi
