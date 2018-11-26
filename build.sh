#!/bin/bash

function usage {
  local tool=$(basename $0)
  cat <<EOF

  USAGE:
    $ $tool [h|-h|--help] COMMAND

  EXAMPLES:
    $ $tool deps      Install dependencies
    $ $tool build     Build a binary
EOF
  exit 1;
}

function build {
  printf 'building the CLI for Linux, MacOS and Windows...\n'
  GOOS=windows GOARCH=386 go build -o build/countries-x86.exe countries.go
  GOOS=windows GOARCH=amd64 go build -o build/countries.exe countries.go
  GOOS=linux GOARCH=386 go build -o build/countries-x86 countries.go
  GOOS=linux GOARCH=amd64 go build -o build/countries countries.go
  GOOS=darwin GOARCH=386 go build -o build/countries-darwin-x86 countries.go
  GOOS=darwin GOARCH=amd64 go build -o build/countries-darwin countries.go
  printf '...countries CLI built\n';
}

function remove {
  rm -rf build;
}

# total arguments should be 1
if [ $# -ne 1 ]; then
  usage;
fi

# displays help menu
if { [ -z "$1" ] && [ -t 0 ] ; } || [ "$1" == 'h' ] || [ "$1" == '-h' ] || [ "$1" == '--help' ]
then
  usage;
fi

# show help for no arguments if stdin is a terminal
if [ "$1" == "deps" ]; then
  go get github.com/urfave/cli
elif [ "$1" == "build" ]; then
  remove
  build
else
  usage;
fi
