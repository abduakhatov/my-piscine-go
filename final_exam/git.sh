#!/bin/bash

echo $(git status)
echo ">>>[Adding files ...]"
git add .
echo ">>>[Commit message: 'autopush']"
git commit -m "autopush"
echo ">>>[Pushing ...]"
git push
