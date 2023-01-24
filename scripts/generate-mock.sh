#!/bin/bash


if [[ $1 = "port" ]]; then
  path=./internal/port
  destination="./internal/port/mock"

  for f in $path/*.go; do 
    # ext=${f##*_test.go}
    # fileName="${ext##*/}"
    fileName="${f##*/}"

    if [[ $fileName != *_test.go ]]; then
      name=${fileName%.*}
      fileGen="${name}_mock.go"
      
      printf "generating mock for $fileName ..."
      mockgen -source="$path/$fileName" -destination="$destination/$fileGen" -package=mock
      printf "finished\n"
    fi 
  done
  printf "mock successfully generated ...\n"

else

  printf "no mock generated ...\n"

fi
