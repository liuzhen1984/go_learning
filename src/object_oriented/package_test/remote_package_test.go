package package_test

import "testing"
import cm "github.com/easierway/concurrent_map"

//go get -u github.com/golang....

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
