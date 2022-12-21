## Teetime
Like `tee` but adds an timestamp for each row.


### Basic Usage
```
#  echo "test" | ./teetime  test.txt   
1671657010238: test
```
```
#   echo "test" | ./teetime  test.txt -H                        
2022-12-21 22:10:07: test
```