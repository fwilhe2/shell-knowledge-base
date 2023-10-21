#!/usr/bin/env bash

A=$([ -f src/A ] && cat src/A || echo 'B' )

B=$([ -f src/B ] && cat src/B || echo 'A' )
echo $A $B

