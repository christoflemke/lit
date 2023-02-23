#!/bin/bash
set -ex
cd $(dirname $0)
export REPO_ROOT=$(cd .. && pwd)
env
(
  set -e
  REF_REPO_PATH=reference-repo
  rm -rf "$REF_REPO_PATH"
  git init $REF_REPO_PATH
  cd $REF_REPO_PATH
  echo 'hello' > hello.txt
  echo 'world' > world.txt
  git add .
  git commit -m test
)

(
  set -e
  MY_REPO_PATH=my-repo
  rm -rf "MY_REPO_PATH"
  $GOROOT/bin/go run $REPO_ROOT/main.go init $MY_REPO_PATH
  cd $MY_REPO_PATH
  echo 'hello' > hello.txt
  echo 'world' > world.txt
  $GOROOT/bin/go run $REPO_ROOT/main.go commit
)