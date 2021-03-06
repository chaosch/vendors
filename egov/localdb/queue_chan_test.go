package localdb

import (
	"fmt"
	"time"
	"testing"
	"strconv"
)

func TestChanQueue_EnqueueObject(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	q, err := OpenChanQueue(file, 10)
	if err != nil {
		t.Error(err)
	}
	defer q.Close()

	type object struct {
		Value []byte
		Key   string
	}

	msg := &object{}
	msg.Key= "dkfjdkf"
	msg.Value = []byte("ddddd")

	for i := 1; i <= 5; i++ {
		msg.Key= "dkfjdkf" + strconv.Itoa(i)
		msg.Value = []byte("ddddd"+ strconv.Itoa(i))
		if item, err := q.EnqueueObject("jac", *msg); err == nil {
			rmsg := &object{}
			item.ToObject(rmsg)
			fmt.Println("out",rmsg.Key, string(rmsg.Value))
			//t.Error(err)
		}
	}

	vals, err := q.PeekStart("jac")
	if err != nil {
		if err.Error() != "out of len" {
			fmt.Println("send offline err:", err)
		}
		//return
	}
	for _, v := range vals {
		remsg := &object{}
		err = v.ToObject(remsg)
		if err != nil {
			fmt.Println("send offline to object err:", err)
		}
		fmt.Println(remsg.Key, string(remsg.Value))
	}
}

func TestQueueChan_Enqueue(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	q, err := OpenChanQueue(file, 10)
	if err != nil {
		t.Error(err)
	}
	defer q.Close()

	for i := 1; i <= 5; i++ {
		if _, err = q.Enqueue("jac", []byte(fmt.Sprintf("value for item %d", i))); err != nil {
			t.Error(err)
		}
	}

	for i := 1; i <= 8; i++ {
		if _, err = q.Enqueue("quy", []byte(fmt.Sprintf("value for item %d", i))); err != nil {
			t.Error(err)
		}
	}

	for i := 1; i <= 5; i++ {
		deqItem, err := q.Dequeue("jac")
		if err != nil {
			t.Error(err)
		}
		fmt.Println("deq:", deqItem.ID, string(deqItem.Value))
	}

	for i := 1; i <= 8; i++ {
		deqItem, err := q.Dequeue("quy")
		if err != nil {
			t.Error(err)
		}
		fmt.Println("deq:", deqItem.ID, string(deqItem.Value))
	}
}

func TestChanQueue_Clear(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	q, err := OpenChanQueue(file, 10)
	if err != nil {
		t.Error(err)
	}
	defer q.Close()
	for i := 1; i <= 5; i++ {
		if _, err = q.Enqueue("jac", []byte(fmt.Sprintf("value for item %d", i))); err != nil {
			t.Error(err)
		}
	}

	q.Clear("jac")

	_, err = q.Dequeue("jac")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println("deq:",deqItem.ID, string(deqItem.Value))
}

func BenchmarkQueueChan_Dequeue(b *testing.B) {
	// Open test database
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	q, err := OpenChanQueue(file, 3)
	if err != nil {
		b.Error(err)
	}
	defer q.Drop()

	// Fill with dummy data
	for n := 0; n < b.N; n++ {
		if _, err = q.Enqueue("jac", []byte("value")); err != nil {
			b.Error(err)
		}
	}

	// Start benchmark
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_, _ = q.Dequeue("jac")
	}
}
