#!/bin/bash

LOCALURL="http://localhost:8888/notify"

if [ "$1" = "pr" ]; then
    curl -d '{ "project": "test project", "target": "master", "changeId": "00", "author": "someone", "changeUrl": "http://imgur.com" }' "$LOCALURL/pr"
elif [ "$1" = "s" ]; then
    curl -d '{ "project": "test project", "result": "-", "phase": "started", "build_url": "http://reddit.com"}' "$LOCALURL/build"
elif [ "$1" = "o" ]; then
    curl -d '{ "project": "test project", "result": "SUCCESS", "phase": "finished", "build_url": "http://reddit.com"}' "$LOCALURL/build"
elif [ "$1" = "f" ]; then
    curl -d '{ "project": "test project", "result": "FAILURE", "phase": "finished", "build_url": "http://reddit.com"}' "$LOCALURL/build"
elif [ "$1" = "a" ]; then
    curl -d '{ "project": "test project", "result": "ABORTED", "phase": "finished", "build_url": "http://reddit.com"}' "$LOCALURL/build"
elif [ "$1" = "i" ]; then
    curl -d '{ "project": "test project", "result": "-", "phase": "waiting", "build_url": "http://reddit.com"}' "$LOCALURL/build"
elif [ "$1" = "u" ]; then
    curl -d '{ "project": "test project", "result": "-", "phase": "invalid", "build_url": "http://reddit.com"}' "$LOCALURL/build"
else
    curl -d '{ "project": "test project", "result": "-", "phase": "invalid", "build_url": "http://reddit.com"}' "$LOCALURL/build"
fi
