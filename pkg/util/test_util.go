package util

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/mock"
	"sync"
)

var BeforeAll = func(beforeAllFunc func()) {
	var once sync.Once

	BeforeEach(func() {
		once.Do(func() {
			beforeAllFunc()
		})
	})
}


func IsContain(calls []mock.Call, method string) bool {
	for _, v := range calls {
		if v.Method == method {
			return true
		}
	}
	return false
}

