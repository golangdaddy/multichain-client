package multichain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

const (
	key1       = "key1"
	key2       = "key2"
	publisher1 = "publisher1"
	publisher2 = "publisher2"
)

type queryKeys struct {
	KeysKey, KeysKeyNotExisting, PublishersKey, PublishersKeyNotExisting string
}

var _ = Describe("ApiListstreamqueryitems", func() {

	DescribeTable("getQuery func provides correct query object for",
		func(keys, publishers []string, qKeys queryKeys, expKey, expPublishers interface{}) {
			q, err := getQuery(keys, publishers)
			Expect(err).ToNot(HaveOccurred())
			Expect(q[qKeys.KeysKey]).To(Equal(expKey))
			Expect(q[qKeys.PublishersKey]).To(Equal(expPublishers))
			Expect(q[qKeys.KeysKeyNotExisting]).To(BeNil())
			Expect(q[qKeys.PublishersKeyNotExisting]).To(BeNil())
		},
		Entry("keys and publishers slices each having just one object",
			[]string{key1},
			[]string{publisher1},
			queryKeys{"key", "keys", "publisher", "publishers"},
			key1,
			publisher1,
		),
		Entry("keys slice having multiple, publishers slices just one object",
			[]string{key1, key2},
			[]string{publisher1},
			queryKeys{"keys", "key", "publisher", "publishers"},
			[]string{key1, key2},
			publisher1,
		),
		Entry("keys slice having just one, publishers slices multiple objects",
			[]string{key1},
			[]string{publisher1, publisher2},
			queryKeys{"key", "keys", "publishers", "publisher"},
			key1,
			[]string{publisher1, publisher2},
		),
		Entry("keys and publishers slices each having multiple objects",
			[]string{key1, key2},
			[]string{publisher1, publisher2},
			queryKeys{"keys", "key", "publishers", "publisher"},
			[]string{key1, key2},
			[]string{publisher1, publisher2},
		),
	)

	DescribeTable("getQuery func provides correct query object without publishers",
		func(keys, publishers []string, expKey string, expValue interface{}) {
			q, err := getQuery(keys, publishers)
			Expect(err).ToNot(HaveOccurred())
			Expect(q[expKey]).To(Equal(expValue))
			Expect(q["publisher"]).To(BeNil())
			Expect(q["publishers"]).To(BeNil())
		},
		Entry("just one keys", []string{key1}, []string{}, "key", key1),
		Entry("multiple keys", []string{key1, key2}, []string{}, "keys", []string{key1, key2}),
	)

	DescribeTable("getQuery func provides correct query object without keys",
		func(publishers, keys []string, expKey string, expValue interface{}) {
			q, err := getQuery(keys, publishers)
			Expect(err).ToNot(HaveOccurred())
			Expect(q[expKey]).To(Equal(expValue))
			Expect(q["key"]).To(BeNil())
			Expect(q["keys"]).To(BeNil())
		},
		Entry("just one publisher", []string{publisher1}, []string{}, "publisher", publisher1),
		Entry("multiple publishers", []string{publisher1, publisher2}, []string{}, "publishers", []string{publisher1, publisher2}),
	)

	Specify("getQuery func provides an error if keys and publishers slices are empty", func() {
		keys := []string{}
		publishers := []string{}
		_, err := getQuery(keys, publishers)
		Expect(err).To(HaveOccurred())
	})
	
})
