#!/bin/bash

declare -a FILENAME

sed -i -e 's|\\|/|g' "${1}".jmx

FILENAME=$(grep -o '>[^`]*[^`]*/*.csv' "${1}".jmx  | xargs -I{} basename {})
FILENAME=($FILENAME)

for m in "${FILENAME[@]}"
do
  echo "before $m"
  sed -i -e "s|$m|/opt/apache-jmeter-5.3/bin/&|g" ${1}.jmx
done
cat ${1}.jmx | grep csv
echo "after"
FILENAME=$(grep -o '>[^`]*[^`]*/opt' "${1}".jmx | head -1)
#echo $FILENAME
sed -i -e "s|${FILENAME}|>/opt|g" ${1}.jmx
cat ${1}.jmx | grep csv

echo "done modifying csv path"