#!/bin/sh

# Run cleanup
source cleanup.sh

# build binary
go build -o out/

# make bin dir in home and copy binary into it
mkdir -p ~/bin && cp out/grep-in-go ~/bin/grep-in-go

# create grip function and add to shell rc
echo 'function Ggrep() { ~/bin/grep-in-go "$@"; }' >> ~/.$@rc

rm -rf out/