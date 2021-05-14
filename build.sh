#/!bin/bash
RFILE="run.sh"
SFILE="stop.sh"
FILENAME=$(go run . version)
[[ -d ./build/ ]] || mkdir build
go build -o ./build/$FILENAME .

echo '#!/usr/bin/bash' > $RFILE
echo 'PIDFILE=$(./build/'$FILENAME' pidfile)' >> $RFILE
echo 'touch $PIDFILE' >> $RFILE
echo "./build/$FILENAME >> log 2>> log &" >> $RFILE
echo 'PID=$!' >> $RFILE
echo 'echo $PID > $PIDFILE' >> $RFILE

echo '#!/usr/bin/bash' > $SFILE
echo 'PIDFILE=$(./build/'$FILENAME' pidfile)' >> $SFILE
echo 'PID=`cat $PIDFILE`' >> $SFILE
echo 'kill -9 $PID' >> $SFILE
echo 'ps aux | grep -i xcheck | awk '"'{print \$2}'"' | xargs -i kill -9 {} || true' >> $SFILE

chmod +x run.sh
chmod +x stop.sh
