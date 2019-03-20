#!/bin/bash

today=$(date "+%Y%m%d")
git add .
git commit -m "git_${today}"
git push -u origin master
