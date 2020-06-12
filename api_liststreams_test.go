package multichain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ApiListstreams", func() {

	DescribeTable("getListStreamsParams function provides correct params for",

		func(streams string, verbose bool, expectedParams []interface{}) {
			actual := getListStreamsParams(streams, verbose)
			Expect(actual).To(Equal(expectedParams))
		},
		Entry("all streams (asterisk), verbosely",
			"*", true, []interface{}{"*", true}),
		Entry("all streams (empty string), verbosely",
			"", true, []interface{}{"*", true}),
		Entry("all streams (asterisk), non-verbosely",
			"*", false, []interface{}{"*", false}),
		Entry("all streams (empty string), non-verbosely",
			"", false, []interface{}{"*", false}),
	)

	DescribeTable("getListStreamsParams function provides correct params for",

		func(streams string, verbose bool, expectedParams []interface{}) {
			actual := getListStreamsParams(streams, verbose)
			Expect(len(actual)).To(Equal(len(expectedParams)))
			Expect(actual[0]).To(Equal(expectedParams[0]))
			Expect(actual[1]).To(Equal(expectedParams[1]))
		},
		Entry("single stream, verbosely",
			"my-stream", true, []interface{}{[]string{"my-stream"}, true}),
		Entry("multiple streams, non-verbosely",
			"my-stream-1,my-stream-2", false, []interface{}{[]string{"my-stream-1", "my-stream-2"}, false}),
	)

})
