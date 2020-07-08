package main

import (
	"fmt"
	"github.com/google/btree"
	ulid "github.com/oklog/ulid/v2"
	"github.com/pborman/uuid"
	"math/rand"
	"strings"
	"time"
)

type AutoIntKey struct {
	val int
}

func (ai AutoIntKey) Less(b btree.Item) bool {
	return ai.val < b.(AutoIntKey).val
}

type StringKey struct {
	val string
}

func (sk StringKey) Less(b btree.Item) bool {
	return strings.Compare(sk.val, b.(StringKey).val) < 0
}

func main() {
	bt := btree.New(6)

	for i := 0; i < 2000000; i++ {
		bt.ReplaceOrInsert(StringKey{val: uuid.New()})
	}

	fmt.Println("uuid nodes", bt.NodeCount())
	//bt.Print()

	bt = btree.New(6)

	for i := 0; i < 2000000; i++ {
		bt.ReplaceOrInsert(AutoIntKey{val: i})
	}

	fmt.Println("auto int nodes", bt.NodeCount())
	//bt.Print()

	bt = btree.New(6)

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	for i := 0; i < 2000000; i++ {
		u, _ := ulid.New(ulid.Timestamp(time.Now()), entropy)
		bt.ReplaceOrInsert(StringKey{val: u.String()})
		if i%2000 == 0 {
			time.Sleep(time.Millisecond)
		}
	}

	fmt.Println("ulid nodes", bt.NodeCount())
	//bt.Print()

	bt = btree.New(6)

	for i := 0; i < 2000000; i++ {

		bt.ReplaceOrInsert(AutoIntKey{val: rand.Int()})
		//time.Sleep(time.Millisecond)
	}

	fmt.Println("rand int nodes", bt.NodeCount())
	//bt.Print()
}
