package ordereddict_test

import (
	"github.com/ivanovaleksey/algogo/ordereddict"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderedDict(t *testing.T) {
	t.Run("with empty dict", func(t *testing.T) {
		od := ordereddict.NewOrderedDict()

		assert.Empty(t, od.Keys())
	})

	t.Run("can add keys", func(t *testing.T) {
		od := ordereddict.NewOrderedDict()

		od.Add("foo", "bar")
		require.Equal(t, []string{"foo"}, od.Keys())

		od.Add("foo2", "bar")
		od.Add("foo1", "baz")
		require.Equal(t, []string{"foo", "foo2", "foo1"}, od.Keys())
	})

	t.Run("can delete first key", func(t *testing.T) {
		od := ordereddict.NewOrderedDict()
		od.Add("foo1", "bax")
		od.Add("foo2", "bay")

		ok := od.Delete("foo1")

		require.True(t, ok)
		require.Equal(t, []string{"foo2"}, od.Keys())
	})

	t.Run("can delete middle key", func(t *testing.T) {
		od := ordereddict.NewOrderedDict()

		od.Add("foo1", "bax")
		od.Add("foo2", "bay")
		od.Add("foo3", "baz")

		ok := od.Delete("foo2")

		require.True(t, ok)
		require.Equal(t, []string{"foo1", "foo3"}, od.Keys())
	})

	t.Run("can delete last key", func(t *testing.T) {
		od := ordereddict.NewOrderedDict()
		od.Add("foo1", "bax")
		od.Add("foo2", "bay")

		ok := od.Delete("foo2")

		require.True(t, ok)
		require.Equal(t, []string{"foo1"}, od.Keys())
	})

	t.Run("can delete everything and add new keys", func(t *testing.T) {
		od := ordereddict.NewOrderedDict()
		od.Add("foo1", "bax")
		od.Add("foo2", "bay")

		od.Delete("foo1")
		od.Delete("foo2")

		require.Equal(t, []string{}, od.Keys())

		od.Add("foo2", "new-bax")
		od.Add("foo1", "new-bay")

		require.Equal(t, []string{"foo2", "foo1"}, od.Keys())

		foo1, ok := od.Get("foo1")
		require.True(t, ok)
		require.Equal(t, "new-bay", foo1)

		foo2, ok := od.Get("foo2")
		require.True(t, ok)
		require.Equal(t, "new-bax", foo2)
	})
}
