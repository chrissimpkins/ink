#!/bin/sh

#INK_VERSION="v0.6.0"
#
#env GOOS=linux GOARCH=amd64 go build .
#mv ink archives/linux-amd64/ink
#
#
#env GOOS=darwin GOARCH=amd64 go build .
#mv ink archives/macOS-amd64/ink
#
#
#env GOOS=windows GOARCH=386 go build .
#mv ink.exe archives/windows-386/ink.exe
#
#
#env GOOS=windows GOARCH=amd64 go build .
#mv ink.exe archives/windows-amd64/ink.exe
#
#cd archives || exit
#
#zip -r "ink-${INK_VERSION}-linux-amd64.zip" linux-amd64 -x "*.DS_Store"
#zip -r "ink-${INK_VERSION}-macOS-amd64.zip" macOS-amd64 -x "*.DS_Store"
#zip -r "ink-${INK_VERSION}-windows-386.zip" windows-386 -x "*.DS_Store"
#zip -r "ink-${INK_VERSION}-windows-amd64.zip" windows-amd64 -x "*.DS_Store"
#
#hsh sha256 "ink-${INK_VERSION}-linux-amd64.zip"
#echo " "
#hsh sha256 "ink-${INK_VERSION}-macOS-amd64.zip"
#echo " "
#hsh sha256 "ink-${INK_VERSION}-windows-386.zip"
#echo " "
#hsh sha256 "ink-${INK_VERSION}-windows-amd64.zip"

export GITHUB_TOKEN="$GORELEASE_GITHUB_TOKEN"
goreleaser --rm-dist