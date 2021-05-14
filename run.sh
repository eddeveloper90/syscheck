#!/usr/bin/bash
PIDFILE=$(./build/xcheck-0.0.5 pidfile)
touch $PIDFILE
./build/xcheck-0.0.5 >> log 2>> log &
PID=$!
echo $PID > $PIDFILE
