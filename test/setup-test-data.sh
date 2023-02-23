#!/bin/bash
set -ex
cd $(dirname $0)
export REPO_ROOT=$(cd .. && pwd)
export GIT_AUTHOR_NAME="Christof Lemke"
export GIT_AUTHOR_EMAIL="doesnotexist@gmail.com"
env
(
  set -e
  REF_REPO_PATH=reference-repo
  rm -rf "$REF_REPO_PATH"
  git init $REF_REPO_PATH
  git config --local user.name $GIT_AUTHOR_NAME
  git config --local user.email $GIT_AUTHOR_EMAIL
  cd $REF_REPO_PATH
  echo 'hello' > hello.txt
  echo 'world' > world.txt
  git add .
  git commit -m 'test message'
)

(
  set -e
  MY_REPO_PATH=my-repo
  rm -rf "$MY_REPO_PATH"
  $GOROOT/bin/go run $REPO_ROOT/main.go init $MY_REPO_PATH
  cd $MY_REPO_PATH
  echo 'hello' > hello.txt
  echo 'world' > world.txt
  echo 'test message' | $GOROOT/bin/go run $REPO_ROOT/main.go commit
)