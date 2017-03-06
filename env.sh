if [ -z "$GOPATH" ]; then
	export $GOPATH=$(pwd)
else
	export GOPATH=$GOPATH:$(pwd)
fi


BASE=$(pwd)

export GOPATH=$GOPATH:$BASE/third
