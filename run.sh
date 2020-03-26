#!/bin/bash
#setup config
echo "Setup config"
source ./conf/setup.env

#Run server
echo "start server:1234 by reflex"
reflex -r '\.(html|go)' -s go run main.go