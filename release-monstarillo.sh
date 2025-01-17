#!/bin/bash

version=""
versionDescription=""
while getopts "h?v:d:" opt; do
    case "$opt" in
    h|\?)
        echo "Usage: $0 [-v version] [-d versionDescription]"
        exit 0
        ;;
    v)  version=$OPTARG
        ;;
    d)  versionDescription=$OPTARG
            ;;

    esac
done

if [ -z "$version" ] || [ -z "$versionDescription"]; then
  echo "version and versionDescription are required"
  exit 1
fi

echo "releasing version '$version' '$versionDescription'"
git tag -a $version -m "$versionDescription"
git push origin $version