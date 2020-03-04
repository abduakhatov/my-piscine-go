#!/bin/bash

go build
echo "1 + 1 = 2$"
echo $(./doop 1 + 1 | cat -e)
echo
echo "hello + 1 = 0$"
echo $(./doop hello + 1 | cat -e)
echo
echo "1 p 1 = 0$"
echo $(./doop 1 p 1 | cat -e)
echo
echo "1 / 0 = No division by 0$"
echo $(./doop 1 / 0 | cat -e)
echo
echo "1 % 0 = No modulo by 0$"
echo $(./doop 1 % 0 | cat -e)
echo
echo "9223372036854775807 + 1 = 0"
echo $(./doop 9223372036854775807 + 1)
echo
echo "-9223372036854775809 - 3 = 0"
echo $(./doop -9223372036854775809 - 3)
echo

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
echo
echo "9223372036854775807 * 3 = 0"
echo $(./doop 9223372036854775807 "*" 3)
echo
echo "-9223372036854775808 * 3 = 0"
echo $(./doop -9223372036854775808 "*" 3)
echo
echo "-9223372036854775808 * -3 = 0"
echo $(./doop -9223372036854775808 "*" -3)
echo
echo "-2 * -2 = 4"
echo $(./doop -2 "*" -2)
echo
echo "2 * 2 = 4"
echo $(./doop 2 "*" 2)
echo
echo "-2 * 2 = -4"
echo $(./doop -2 "*" 2)
echo
echo "2 * -2 = -4"
echo $(./doop 2 "*" -2)
echo
echo "-9223372036854775808 * -9223372036854775808 = 0"
echo $(./doop -9223372036854775808 "*" -9223372036854775808)
echo

echo "9223372036854775807 * 1 = 9223372036854775807"
echo $(./doop 9223372036854775807 "*" 1)
echo

echo "9223372036854775807 * -3 = 0"
echo $(./doop 9223372036854775807 "*" -3)
echo

echo "9223372036854775807 * -1 = -9223372036854775807"
echo $(./doop 9223372036854775807 "*" -1)
echo

echo "9223372036854775807 * 9223372036854775807 = 0"
echo $(./doop 9223372036854775807 "*" 9223372036854775807)
echo

echo "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
echo

echo "-9223372036854775807 + 3 = -9223372036854775804"
echo $(./doop -9223372036854775807 + 3)
echo
echo "9223372036854775807 - 3 = 9223372036854775804"
echo $(./doop 9223372036854775807 - 3)
echo
echo "9223372036854775804 + -3 = 9223372036854775801"
echo $(./doop 9223372036854775804 + -3)
echo
echo "9223372036854775807 + -3 = 9223372036854775804"
echo $(./doop 9223372036854775807 + -3)
echo
echo "-9223372036854775807 - 3 = 0"
echo $(./doop -9223372036854775807 - 3)
echo
echo "-9223372036854775807 - -9223372036854775806 = -1"
echo $(./doop -9223372036854775807 - -9223372036854775806)
echo
echo "-9223372036854775805 - -9223372036854775808 = 3"
echo $(./doop -9223372036854775805 - -9223372036854775808)
echo
echo "9223372036854775807 "*" 3 = 0"
echo $(./doop 9223372036854775807 "*" 3)
echo
echo "1 "*" 1 = 1"
echo $(./doop 1 "*" 1)
echo
echo "1 "*" -1 = -1"
echo $(./doop 1 "*" -1)
echo
echo "10 "/" 2 = 5"
echo $(./doop 10 "/" 2)
