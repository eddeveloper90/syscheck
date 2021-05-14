all:
#	[[ -d ./build/ ]] || mkdir ./build/
#	go build -o ./build/`go run . version` .
	./build.sh

install:
	[[ -d ./build/ ]] || mkdir build
	[[ -d ~/xcheck ]] || mkdir -p ~/xcheck/ && cp app.conf ~/xcheck
	[[ -d ~/xcheck/build/ ]] || mkdir -p ~/xcheck/build/
	cp -rf ./build/* ~/xcheck/build/
	cp -rf ./run.sh ~/xcheck/
	cp -rf ./stop.sh ~/xcheck/
	chmod a+x ~/xcheck/run.sh
	chmod a+x ~/xcheck/stop.sh

.PHONY: clean
clean:
	rm -rf ./build/*