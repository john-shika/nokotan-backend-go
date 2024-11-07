#!/usr/bin/env bash

find . -type f | grep -Ei '\.[a-zA-Z]+\~$' | xargs echo
