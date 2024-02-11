#!/bin/bash

# build executable in home dir
go build -o ~/bin/ggrep ./cmd/cli

# modify PATH env to let OS know where binary live
echo export PATH=$PATH:~/bin >> $1