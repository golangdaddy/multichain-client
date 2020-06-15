package multichain

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ApiSubscribes", func() {
	DescribeTable("appendInnerParams function derives inner params correctly for",
		func(indexTypes []IndexType, retrieveAllOffchain bool, expected []interface{}) {
			actual, err := appendInnerParams(indexTypes, retrieveAllOffchain, []interface{}{})
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(expected))
		},
		Entry("no index types and retrieveAllOffchain = false",
			[]IndexType{}, false,
			[]interface{}{}),
		Entry("no index types and retrieveAllOffchain = true",
			[]IndexType{}, true,
			[]interface{}{"retrieve"}),
		Entry("one index type and retrieveAllOffchain = false",
			[]IndexType{ IndexItems }, false,
			[]interface{}{ string(IndexItems) }),
		Entry("one index type and retrieveAllOffchain = true",
			[]IndexType{ IndexItems }, true,
			[]interface{}{ fmt.Sprintf("%s,%s", IndexItems, "retrieve") }),
		Entry("some index types and retrieveAllOffchain = false",
			[]IndexType{ IndexItems, IndexKeys, IndexKeysLocal }, false,
			[]interface{}{ fmt.Sprintf("%s,%s,%s", IndexItems, IndexKeys, IndexKeysLocal) }),
		Entry("some index types and retrieveAllOffchain = true",
			[]IndexType{ IndexItems, IndexKeys, IndexKeysLocal }, true,
			[]interface{}{ fmt.Sprintf("%s,%s,%s,%s", IndexItems, IndexKeys, IndexKeysLocal, "retrieve") }),
	)
})
