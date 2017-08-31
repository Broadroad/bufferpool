package bufferpool

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleNew() {
	bp := NewBufferPool(10, 1024)
	buffer := bp.Get()
	buffer.WriteString("Hello World!")
	bp.Put(buffer)

	buffer1 := bp.Get()
	buffer1.WriteString("World Hello!")

	// buffer1 is World Hello!
	fmt.Println(buffer1)
	// Length of buffer pool is 1
	fmt.Println(len(bp.c))
}

func Testbp(t *testing.T) {
	size := 4
	capacity := 1024

	bp := NewBufferPool(size, capacity)

	b := bp.Get()

	// Check whether the capacity is correct before we use it
	if cap(b.Bytes()) != capacity {
		t.Fatalf("buffer capacity is incorrect: got %v want %v", cap(b.Bytes()), capacity)
	}

	// Increase the buffer beyond our capacity and then return it to the buffer poool
	b.Grow(capacity * 3)
	bp.Put(b)

	for i := 0; i < size; i++ {
		bp.Put(bytes.NewBuffer(make([]byte, 0, bp.a*2)))
	}

	// Check that buffers are full or not
	if len(bp.c) < size {
		t.Fatalf("buffer pool is too small, got %v wat %v", len(bp.c), size)
	}

	close(bp.c)

	// Check that there are buffers of the correct capacity in the pool
	for buffer := range bp.c {
		if cap(buffer.Bytes()) != bp.a {
			t.Fatalf("returned buffers wrong capacity: got %v want %v", cap(buffer.Bytes()), capacity)
		}
	}
}
