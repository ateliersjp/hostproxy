#!/bin/sh

if [ "$1" != 'hostproxy' ]; then
    set -- 'hostproxy' "$@"
fi

exec "$@"
