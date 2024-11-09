#!/bin/bash
# or #!/bin/sh

# condition
if [ -z $0 ]; then
    echo "Run command 1"
elif [ -n $0 ]; then
    echo "Run command 2"
else
    echo "Run command 3"
fi

# switch ... case - bash style
case $# in
3) ;; ## We need 3 args, so do nothing
*)
    printf "%s\n" "Please provide three names" >&2
    exit 1
    ;;
esac

# read a file line by line assigning the value to a variable:
while read -r line; do
    echo "Text read from file: $line"
done <my_filename.txt

# loop string in array
array=("string1" "string2")
for string in ${array[@]}; do
    echo $string
done

# loop with array index
allThreads=("string1" "string2")
for i in ${!allThreads[@]}; do
    echo ${allThreads[$i]}
done
