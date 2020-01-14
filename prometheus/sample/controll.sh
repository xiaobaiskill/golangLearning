export binaryName=promethues

function build(){
  GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $binaryName
}

function start(){
  echo "start success"
}

function stop(){
  echo "stop success"
}

function restart() {
    stop
    sleep 1
    start
}

function help(){
  echo "$0 build"
}

if [ "$1" == "build" ];then
  build
elif [ "$1" == "start" ];then
  start
elif [ "$1" == "stop" ];then
  stop
elif [ "$1" == "restart" ];then
  restart
else
  help
fi
