#!/usr/bin/bash
PIDFILE=$(./build/open .env: no such file or directory pidfile)
PID=`cat $PIDFILE`
kill -9 $PID
ps aux | grep -i xcheck | awk '{print $2}' | xargs -i kill -9 {} || true
