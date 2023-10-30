package cached_test

import (
	"log"
	"testing"

	"notchman.tech/chat-tkb/src/cached"
)

func TestMemcache(t *testing.T) {

	m := cached.NewMemcached("memcached:11211")
	if e := m.Save("sample", []byte("This is the SHIOKAZE limited express bound for Okayama")); e != nil {
		log.Println(e.Error())
	}

	if v, e := m.Get("sample"); e != nil {
		log.Println(e.Error())
	} else if v == nil {
		log.Println("value is empty")
	} else {
		log.Println(string(v))
	}
}