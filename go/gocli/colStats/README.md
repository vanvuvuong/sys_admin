
# Create benchmark CSV file
```
#/bin/fish
for i in (seq 0 2501); echo "col1,col2" > file$i.csv; end
for i in (seq 0 2501); for ii in (seq 0 2501); echo "data$ii,$ii" >> file$i.csv; end; end
```