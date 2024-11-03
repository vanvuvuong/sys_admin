function ecf
    # encrypt files with gpg private key, ignore the encrypted & images files
    # $1: gpg keyname
    # $2: file path
    set -l checking (string match -er '.*asc|.*jpg|.*png' $argv[2])
    if test (count $checking) -eq 0
        gpg --batch --yes --armor --encrypt-files --recipient $argv[1] $argv[2] 2>/dev/null
    end
end

function dcf
    # decrypt acs files with gpg private key
    # $1: gpg keyname
    # $2: file path
    set -l checking (string match -er '.*asc|' $argv[2])
    if test (count $checking) -gt 0
        gpg --batch --yes --armor --decrypt-files --recipient $argv[1] $argv[2] 2>/dev/null
    end
end

function eccf
    # encrypt files in current folder (ignore files with extension asc, md, png)
    # args is gpg private encryption key name
    for file in ./*
        ecf $argv $file
    end
end

function dccf
    # decrypt files in currcent folder
    # args is gpg private encryption key name
    for file in ./*
        dcf $argv $file
    end
end

function ecfp
    # encrypt files in current folder depend on pattern
    # args[1] is gpg private encryption key name
    # args[2] is file pattern
    for file in ./$argv[2]
        ecf $argv[1] $file
    end
end

function dcfp
    # decrypt files in currcent folder
    # args[1] is gpg private encryption key name
    # args[2] is file pattern
    for file in ./$argv[2]
        dcf $argv[1] $file
    end
end

function ecasf
    # encrypt all files in current folder & sub-folder
    set -l current_folder (pwd)
    read -l -P "Process ENCRYPTING data. Do you want to continue? (y/n) ->" confirm
    switch $confirm
        case y Y
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
        case '*'
            echo "Do nothing..."
    end
end

function dcasf
    # decrypt all files in current folder & sub-folder
    set -l current_folder (pwd)
    read -l -P "Process DECRYPTING data & OVERRIDE old files. Continue? (y/n) ->" confirm
    switch $confirm
        case y Y
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
        case '*'
            echo "Do nothing..."
    end
end
