projectName = sono
gopath = .
buildPath = .
projectPath = ${gopath}

all:build
 	
build:
	go build -o ${buildPath}/${buildPath} -ldflags -w ${projectPath}

run:build	
	nohup ./${buildPath}  >> ./output.log 2>&1 &