# loop
go-lang 반복문 흐름을 제어하는 예제 입니다.

## Examples
for loop
```
for i := 0; i < 5; i++ {
}
```

slice array loop
```
numbers := []int{1,2,3,4,5}
for i := range numbers {
    f.Println(i)	
}
```

map collection loop
```
capacityUnit := make(map[string]string)
capacityUnit["1KB"] = "1024 Byte"
capacityUnit["1MB"] = "1024 KB"
capacityUnit["1GB"] = "1024 MB"
capacityUnit["1TB"] = "1024 GB"
for key, result := range capacityUnit {
    f.Println(key, result)
}
```

## Build
```
go build -o target/sample sample.go
```

## Run
```
./target/sample
```