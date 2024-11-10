# FIND common use cases <sup><sup>[back](../README.md)</sup></sup>

Find files & folder that changing `less than 5 minutes` except ones in `/proc`, `/sys`, `/var`, `/run` & `/mnt` (`-prunt` option will ignore the whole tree directory)

```shell
find / -cmin -5 ! -path '/proc/*' -prune ! -path '/sys/*' -prune ! -path '/var/*' -prune ! -path '/run/*' -prune ! -path '/mnt/*' -prune
```

---

Find files in current folder that have name pattern ~ `sam*` & replace content in those file from `sample` to `sample2`

```shell
find . -name "sam*" -type f -exec sed -i 's/sample/sample2/g' {} \;
```

---

Find anything that `change more than 24 days` and `size more than 15Gb`

```shell
find . -ctime +24 -size +15G
```

---

Showing empty files content in level 1 of the home directory

```shell
find ~ -maxdepth 1 -size 0c -exec ls -lht {} \;
```

---

Finding the line that contain text "CLS" in all the executable file in current folder

```shell
find . -type f -executable -exec grep "CLS" {} \;
```

---

Execute complicate `cmd` command with some pdf files

```shell
find . -name '*pdf' -exec bash -c 'cmd {}' \;
```

---

Find all folders in 1 parent folder and list them to a text file

```shell
find parent-folder -maxdepth 1 -mindepth 1 -type d -printf '%f\n' >children_foldername
```
