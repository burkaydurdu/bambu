//go:build unit

package bambu

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type HashSuite struct {
	suite.Suite
}

func (t *HashSuite) TestHash_IndexOf() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")

	k, v := hash.IndexOf(0)

	t.Equal(k, "name")
	t.Equal(v, "burkay")

	k, v = hash.IndexOf(1)

	t.Equal(k, "surname")
	t.Equal(v, "durdu")

	k, v = hash.IndexOf(-1)

	t.Equal(k, "")
	t.Equal(v, "")

	k, v = hash.IndexOf(2)

	t.Equal(k, "")
	t.Equal(v, "")
}

func (t *HashSuite) TestHash_Get() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")

	val := hash.Get("surname")

	t.Equal(val, "durdu")

	val = hash.Get("surnames")

	t.Equal(val, "")
}

func (t *HashSuite) TestHash_GetFirstAndLast() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")
	hash.Add("school", "KTU")

	k, v := hash.First()

	t.Equal(k, "name")
	t.Equal(v, "burkay")

	k, v = hash.Last()

	t.Equal(k, "school")
	t.Equal(v, "KTU")

	hash.Clear()

	k, v = hash.First()

	t.Equal(k, "")
	t.Equal(v, "")
	return
}

func (t *HashSuite) TestHash_GetLength() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")
	hash.Add("school", "KTU")

	length := hash.Length()

	t.Equal(length, 3)

	hash.Add("age", "27")

	length = hash.Length()

	t.Equal(length, 4)
}

func (t *HashSuite) TestHash_Remove() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")
	hash.Add("school", "KTU")

	hash.Remove("school")

	t.Equal(hash.Length(), 2)
	t.Equal(hash.Get("school"), "")

	hash.Remove("surname")

	t.Equal(hash.Length(), 1)
	t.Equal(hash.Get("surname"), "")

	hash.Remove("name")

	t.Equal(hash.Length(), 0)
	t.Equal(hash.Get("name"), "")
}

func TestHashSuite(t *testing.T) {
	suite.Run(t, new(HashSuite))
}
