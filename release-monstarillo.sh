#!/bin/bash

version=""
versionDescription=""
token=""
while getopts "h?v:d:t:" opt; do
    case "$opt" in
    h|\?)
        echo "Usage: $0 [-v version] [-d versionDescription] [-d token"
        exit 0
        ;;
    v)  version=$OPTARG
        ;;
    d)  versionDescription=$OPTARG
            ;;
    t)  token=$OPTARG
            ;;
    esac
done

if [ -z "$version" ] || [ -z "$versionDescription"] || [-z "$token"]; then
  echo "version and versionDescription are required"
  exit 1
fi

echo "releasing version '$version' '$versionDescription'"

export GITHUB_TOKEN=$token

git tag -a $version -m "$versionDescription"
git push origin $version

goreleaser release --clean