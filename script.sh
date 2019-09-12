#!/bin/bash
set -e
set -x

env

echo "event:"
cat $GITHUB_EVENT_PATH
echo "workspace:"
cat $GITHUB_WORKSPACE
