#!/bin/bash
set -e
set -x

env

echo "event:"
cat $GITHUB_EVENT_PATH
echo "workspace:"
ls -la $GITHUB_WORKSPACE
