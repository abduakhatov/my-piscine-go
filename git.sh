#!/bin/bash

echo $(git status)
echo $(git add .)
echo $(git commit -m "autopush")
echo $(git push)
