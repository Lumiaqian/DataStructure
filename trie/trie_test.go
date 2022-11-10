package trie

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrie(t *testing.T) {
	Convey("Test Trie", t, func() {
		words := []string{"how", "hi", "her", "hello", "so", "see"}
		var trie Trie
		trie = *trie.NewTrie()
		trie.Init(words)
		Convey("The rsult Find is true", func() {
			So(trie.Find("her"), ShouldBeTrue)
		})
		Convey("The result Find is true", func() {
			trie.Insert("say")
			So(trie.Find("say"), ShouldBeTrue)
		})
		Convey("The result Find is false", func() {
			So(trie.Find("he"), ShouldBeFalse)
		})
		Convey("The result Find nil is false", func() {
			So(trie.Find("say"), ShouldBeFalse)
		})
	})
}
