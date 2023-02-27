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

  echo 'hi!' > hi.txt
  git add hi.txt
  git commit -m 'hi message'
)

(
  set -e
  lit() {
    $GOROOT/bin/go run $REPO_ROOT/main.go "${@}"
  }
  MY_REPO_PATH=my-repo
  rm -rf "$MY_REPO_PATH"
  lit init $MY_REPO_PATH
  cd $MY_REPO_PATH
  echo 'hello' > hello.txt
  echo 'world' > world.txt
  echo 'test message' | lit commit

  echo 'hi!' > hi.txt
  echo 'hi commit message' | lit commit
)