#!/usr/bin/bash
PIDFILE=$(./build/open .env: no such file or directory pidfile)
touch $PIDFILE
./build/open .env: no such file or directory >> log 2>> log &
PID=$!
echo $PID > $PIDFILE
