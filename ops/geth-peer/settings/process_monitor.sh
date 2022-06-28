#!/bin/bash
set -e
# waits for l2geth to be up
curl --fail \
    --show-error \
    --silent \
    --retry-connrefused \
    --retry $RETRIES \
    --retry-delay 1 \
    --output /dev/null \
    $L2_NODE_WEB3_URL

while IFS== read -r key value; do
    if [[ -n "$key" ]]; then
        export $key=$value
    fi
done < /app/env.sh

while true
do
    geth_pid=`ps -ef | grep geth | grep verbosity | grep -v grep | awk '{print $2}'`
    if [[ -z "$geth_pid" ]]; then
        echo "monitor_geth...">>/app/log/t_supervisord.log
        /app/geth.sh
    fi
    sleep 30
done