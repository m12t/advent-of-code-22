#!/bin/bash
for day in {1..25}
do
   mkdir "dec-${day}"
   cd "dec-${day}"
   touch main.go && touch input.txt
   cd ..
done