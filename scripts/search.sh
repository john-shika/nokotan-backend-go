#!/usr/bin/env bash

CurrentWorkDir=$(pwd)
ScriptDir=$(dirname $0)
cd "$ScriptDir"
cd ..

while IFS= read -r line; do echo "$line"; cat "$line" | grep -Ei "$1"; done <<< $(find "app" -type f | grep -Ei '\.go$')

cd "$CurrentWorkDir"
