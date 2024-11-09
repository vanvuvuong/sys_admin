# AWK variable abbriviations

| Variables   | Description                                                                                                                               |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| ENVIRON     | An associative array of environment variables.                                                                                            |
| FIELDWIDTHS | Specify the field width                                                                                                                   |
| FS          | The field seperator                                                                                                                       |
| RS          | The record seperator, default is new line                                                                                                 |
| NF          | The number of fields in the current record                                                                                                |
| NR          | The number of the current record                                                                                                          |
| FNR         | Similar to NR, but relative to the current file. It is useful when AWK is operating on multiple files. Value of FNR resets with new file. |
| OFS         | The output field seperator, space by default                                                                                              |
| ORS         | The output record seperator                                                                                                               |
| FILENAME    | Hold the processed file name                                                                                                              |
| IGNORECASE  | Ignore character case                                                                                                                     |
| $n          | The n<sup>th</sup> field in the current record where the fields are separated by FS                                                       |

#W Usage examples

---

- Basic usage:

```shell
awk 'BEGIN {FS=":"} {OFS=".|--|"} {print NR, $1} END {print FILENAME, "Total: " NR}' /etc/passwd
```

- Condition AWK:

```shell
awk 'BEGIN {FS=":"} $3 > 100 {print FNR, $1, $3} END {print FILENAME, "Total: " FNR}' /etc/passwd
```

- Basic format output

```shell
awk 'BEGIN {FS=":"} {printf "%20s %6d %25s\n", $1, $4, $7}' /etc/passwd
```

- Basic format output with header (-20 mean: 20 characters left-align, 6 means: 6 characters right-align)

```shell
awk 'BEGIN {FS=":";printf "%-20s %6s %25s\n", "Name", "UID", "Shell"} $4 >= 1000 {printf "%-20s %6d %25s\n", $1, $4, $7}' /etc/passwd
```

- Basic format output using file `passwd.awk` (file content below)

```shell
awk -f passwd.awk /etc/passwd
```

> Content of `passwd.awk` file

```shell
cat <<EOF > passwd.awk
function green(s) {
        printf "\033[1;32m" s "\033[0m\n"
}
BEGIN {
        FS=":"
        green("     Name:   UID:             Shell:")
}
{
        if ($3 > 999)
        printf "%10s %6d %18s\n", $1, $3, $7
}
EOF
```
