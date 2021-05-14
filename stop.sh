#!/usr/bin/bash
PIDFILE=$(./build/xcheck-0.0.5 pidfile)
PID=`cat $PIDFILE`
kill -9 $PID
ps aux | grep -i xcheck | awk '{print $2}' | xargs -i kill -9 {} || true
