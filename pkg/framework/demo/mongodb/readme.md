```go
type test struct {
    UserName string 
    Phone    []int64
    Subject  string
}
```
```go
//驼峰命名对应mgo字段全部切换成小写
//UserName--->username
//插入slice对应mgo时array
```


