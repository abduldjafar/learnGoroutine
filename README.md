# learnGoroutine [ still develop]

This repo i create for learn how goroutine work. I use word count as use case for it.

## Result
* without Concurrency
```
C:\Users\rsudo_30\go\src\learnGoroutine\noConcurrency>go run main.go
2020/12/21 15:00:10 Readfile
2020/12/21 15:00:10 processing file
2020/12/21 15:01:30 map[saya:17694720 makan:17694720 nasi:17694720]
2020/12/21 15:01:30 processing took 1m19.7723504s
C:\Users\rsudo_30\go\src\learnGoroutine>
```
* With Concurrency
```
C:\Users\rsudo_30\go\src\learnGoroutine>go run main.go
2020/12/21 15:10:34 Readfile
2020/12/21 15:10:34 processing file
2020/12/21 15:11:21 map[saya:17694720 makan:17694720 nasi:17694720]
2020/12/21 15:11:21 processing took 46.84417s
```
