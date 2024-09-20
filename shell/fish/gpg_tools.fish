#!/bin/fish
# encrypt files with gpg private key, ignore the encrypted & images files
# $1: gpg public keyname
# $2: file path
function ecf
    set -l checking (string match -er '.*asc|.*jpg|.*png' $argv[2])
    if test (count $checking) -eq 0
        gpg --batch --yes --armor --encrypt-files --recipient $argv[1] $argv[2] 2>/dev/null
    end
end

# decrypt `.asc` files with gpg private key
# $1: gpg private keyname
# $2: file path
function dcf
    set -l checking (string match -er '.*asc|' $argv[2])
    if test (count $checking) -gt 0
        gpg --batch --yes --armor --decrypt-files --recipient $argv[1] $argv[2] 2>/dev/null
    end
end

# encrypt files in current folder (ignore files with extension asc, png, jpg)
# args is gpg private encryption key name
function eccf
    for file in ./*
        ecf $argv $file
    end
end

# decrypt files in currcent folder
# args is gpg private encryption key name
function dccf
    for file in ./*
        dcf $argv $file
    end
end

# encrypt all files in current folder & sub-folder
function ecasf
    set -l current_folder (pwd)
    for path in ./**/*
        if test -d $path
            cd $path
            eccf $argv
        else if test -f $path
            ecf $argv $path
        end
        cd $current_folder
    end
    cd $current_folder
end


# decrypt all files in current folder & sub-folder
function dcasf
    set -l current_folder (pwd)
    for path in ./**/*
        if test -d $path
            cd $path
            dccf $argv
        else if test -f $path
            dcf $argv $path
        end
        cd $current_folder
    end
    cd $current_folder
end
