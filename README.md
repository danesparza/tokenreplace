# tokenreplace [![Circle CI](https://circleci.com/gh/danesparza/tokenreplace.svg?style=svg)](https://circleci.com/gh/danesparza/tokenreplace)
A simple token replacement command-line tool for files, written in Go.  

Recently, I needed a nice solution for updating a build number in a deployment document.  I decided to use a token in the deployment document.  The token gets updated in a batch file by calling this tool.

#### Installing
Installing is as simple as grabbing the [latest release](https://github.com/danesparza/tokenreplace/releases) and copying to the directory you need to use the tool in.

#### Command line flags
```sh
  -file="filename.html": File to use for token replacement
  -token="{build}": Token to replace
  -replacement="Build number 1": Replace the token with this string
```
