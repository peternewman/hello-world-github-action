#!/bin/bash

set -e
set -x

if [[ "${GITHUB_EVENT_NAME}" != "pull_request_review" ]]; then
  echo "unsupported event: ${GITHUB_EVENT_NAME}"
  exit 1
fi

user=$(jq -r .review.user.login ${GITHUB_EVENT_PATH})
cmd=$(jq -r .review.body ${GITHUB_EVENT_PATH})
echo "reviewer is ${user}, command is ${cmd}"

if [[ "${cmd}" == "merge" ]]; then
  head=$(jq -r .pull_request.head.ref ${GITHUB_EVENT_PATH})
  git config user.email me@rultor.com
  git config user.name rultor
  git remote set-url --push origin \
    https://oauth2:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git
  git checkout -B __rultor origin/${head}
  git checkout -B master origin/master
  git merge --no-ff __rultor
  git push origin master
fi
exit 0
