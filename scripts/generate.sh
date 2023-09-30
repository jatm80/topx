#!/bin/bash

set -euo pipefail

:> numbers

for i in {0..10000000}
do
  echo $(( $RANDOM % 100000000 + 1 )) >> numbers
done

