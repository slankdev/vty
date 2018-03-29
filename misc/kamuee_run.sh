#!/bin/bash

KAMUEE_HOST="127.0.0.1"
KAMUEE_PORT="9077"

if [ "$#" -lt 1 ]; then
        echo "Usage: kamuee_run.sh -c <command> -c <command>"
        exit 1
fi

RUNCOMMAND="$*"
IFS=''
(expect - << EOF
set timeout 20
log_user 0
spawn telnet $KAMUEE_HOST $KAMUEE_PORT
log_user 1
expect "kamuee-vty*>"
send "terminal length 0\r"
expect "kamuee-vty*>"
send "$RUNCOMMAND\r"
expect "kamuee-vty*>"
send "quit\r"
EOF
) | while read line; do
        if [[ ! "$line" =~ "kamuee-vty" ]]; then
                echo $line
        fi
done
