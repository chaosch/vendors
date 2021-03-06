# YIYIDB - A fast NoSQL database for storing big list of data

[![Author](https://img.shields.io/badge/author-@jacoblai-blue.svg?style=flat)](http://www.icoolpy.com/) [![Platform](https://img.shields.io/badge/platform-Linux,%20OpenWrt,%20Android,%20Mac,%20Windows-green.svg?style=flat)](https://github.com/jacoblai/dhdb) [![NoSQL](https://img.shields.io/badge/db-NoSQL-pink.svg?tyle=flat)](https://github.com/jacoblai/dhdb)


YIYIDB is a high performace NoSQL database

## Features

* Pure Go 
* Big data list to 10 billion
* Queue FIFO
* KV list (z-list)
* KV list TTL time expirse auto del and event
* Android or OpenWrt os supported (ARM/MIPS)

## import
```
import "egov/localdb"
```
## KET VALUE LIST
## open or create database
```
dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
if err != nil {
	panic(err)
}
kv, err := yiyidb.OpenKvdb(dir + "/kvdata", true) //path, enable ttl
if err != nil {
	fmt.Println(err)
	return
}
defer kv.Close()
```
## reg an event func hook on TTL delete
```
kv.OnExpirse = func(key, value []byte) {
   fmt.Println("exp:", string(key), string(value))
}
```

## insert One key value data
```
kv.Put([]byte("hello1"), []byte("hello value"), 0)
```

## insert One key value data TTL 3 seconds expirse auto delete
```
kv.Put([]byte("hello1"), []byte("hello value"), 3)
```

## set an exists key enable TTL 8 seconds expirse auto delete
```
kv.SetTTL([]byte("hello1"), 8)
```

## batch operation
```
items := make([]BatItem,0)
for i := 1; i < 5; i++ {
    item := BatItem{
    	Op: "put",
    	Ttl: 1,
    	Key: []byte("test" + strconv.Itoa(i)),
    	Value: []byte("bat values"),
    }
    items = append(items, item)
}
for i := 1; i < 5; i++ {
    item := BatItem{
    	Op: "del",
    	Ttl: 1,
    	Key: []byte("test" + strconv.Itoa(i)),
    	Value: []byte("bat values"),
    }
    items = append(items, item)
}
kv.BatPutOrDel(&items)
```

## get data
```
vaule, err := kv.Get([]byte("hello1"))
if err != nil {
		fmt.Println(err)
}
```

## all keys
```
all := kv.AllKeys()
for _, k := range all {
	fmt.Println(k)
}
```

## all keys and struct by value
```
type object struct {
	Value int
}
var o object
all := kv.AllByObject(o)
for _, v := range all {
	fmt.Println(string(v.Key), v.Object)
}
```

## all keys and raw by value
```
all := kv.AllByKV()
for _, v := range all {
	fmt.Println(string(v.Key), string(v.Value))
}
```

## keys start with filter and struct by value
```
type object struct {
	Value int
}
var o object
all, err := kv.KeyStartByObject([]byte("key"), o)
if err !=nil{
	fmt.Println(err)
}
for _, v := range all {
	fmt.Println(string(v.Key), v.Object)
}
```

## keys range with filter and struct by value
```
type object struct {
	Value int
}
var o object
all, err := kv.KeyRangeByObject([]byte("minkey"),[]byte("maxkey123"), o)
if err !=nil{
	fmt.Println(err)
}
for _, v := range all {
	fmt.Println(string(v.Key), v.Object)
}
```

## keys start with 
```
searchkeys := kv.KeyStart([]byte("hello1"))
for _, v := range all {
	fmt.Println(string(v.Key), string(v.Value))
}
```

## keys range with
```
randkeys, err := kv.KeyRange([]byte("2017-06-01T01:01:01"), []byte("2017-07-01T01:01:01"))
if err !=nil{
	fmt.Println(err)
}
for _, v := range all {
	fmt.Println(string(v.Key), string(v.Value))
}
```

## QUEUE LIST (FIFO)
## open or create queue
```
dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
if err != nil {
	panic(err)
}
queue, err := yiyidb.OpenQueue(dir + "/queuedata")
if err != nil {
	fmt.Println(err)
	return
}
defer queue.Close()
```

## enqueue push string
```
item, err = q.EnqueueString("value")
```

## enqueue push object
```
type object struct {
	Value int
}
item, err = q.EnqueueObject(object{Value:1})
```

## enqueue push []byte 
```
item, err := q.Enqueue([]byte("value"))
if err != nil{
	t.Error(err)
}
fmt.Print(string(item.Value))
```

## enqueue push batch of []byte
```
vals := make([][]byte, 0)
for i := 1; i < 500; i++ {
	vals = append(vals, []byte("test values" + strconv.Itoa(i)))
}

err = q.EnqueueBatch(vals)
if err != nil {
	t.Error(err)
}

item, err := q.PeekByID(499)
if err != nil {
	t.Error(err.Error())
}
if string(item.Value) != "test values499"{
	t.Error("not encode all")
}
```

## dequeue pop item
```
deqItem, err := q.Dequeue()
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(deqItem.Value))
```

## peek get item (just see get on by auto remove it)
```
peekItem, err := q.Peek()
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(peekItem.Value))
```

## peekbyoffset
```
peekFirstItem, err := q.PeekByOffset(0)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(peekFirstItem.Value))
```

## update queue item bytes value
```
updatedItem, err := q.Update(item.ID, []byte(newCompStr))
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(updatedItem.Value))
```

## update queue item string value
```
updatedItem, err := q.UpdateString(item.ID, "new values")
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(updatedItem.Value))
```

## update queue item object value
```
type object struct {
	Value int
}
updatedItem, err := q.UpdateObject(item.ID, object{Value:1})
if err != nil {
	fmt.Println(err)
}
var obj object
if err := updatedItem.ToObject(&obj); err != nil {
	fmt.Println(err)
}
```

## More sameple code in project *_test.go files

## Perfromace 
## Enqueue or insert kv list bench test
300,000	      5865ns/op	     516B/op	       9allocs/op
## Dequeue or get kv list bench test
200,000	     14379ns/op	    1119B/op	      20allocs/op

## Authors

@jacoblai

## Thanks

* syndtr, github.com/syndtr/goleveldb
