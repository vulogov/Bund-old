#!/bin/bash

rm -f ../../Documentation/BNF/*.png

for i in `ls *.dot`; do
  dot -Tpng $i > ../../Documentation/BNF/$i.png
done
