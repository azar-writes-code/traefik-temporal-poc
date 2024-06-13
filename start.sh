#!/bin/sh
if [ "$1" = "server" ]; then
    echo "Starting server..."
    exec ./server
elif [ "$1" = "worker" ]; then
    echo "Starting worker..."
    exec ./worker
else
    echo "Usage: start.sh {server|worker}"
    exit 1
fi
