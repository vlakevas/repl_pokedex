package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://site1.dev",
			val: []byte("testCase1"),
		},
		{
			key: "https://site2.dev/page3/photos/image1.png",
			val: []byte("testCase2"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case #%d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const intervalTime = 5 * time.Millisecond
	const waitTime = 2 * intervalTime

	testCache := NewCache(intervalTime)

	testCache.Add("testKey", []byte("dataForTesting"))

	if _, exists := testCache.Get("testKey"); !exists {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	if _, exists := testCache.Get("testKey"); exists {
		t.Errorf("expected key to be missing")
		return
	}

}
