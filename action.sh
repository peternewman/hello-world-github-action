#!/bin/bash

set -e
set -x

if [[ "${GITHUB_EVENT_NAME}" != "pull_request_review" ]]; then
  echo "unsupported event: ${GITHUB_EVENT_NAME}"
  exit 1
fi

USER=$(jq -r .review.user.login ${GITHUB_EVENT_PATH})
echo "reviewer is ${USER}"
exit 0
