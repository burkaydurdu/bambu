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

	k, v := hash.indexOf(0)

	t.Equal(k, "name")
	t.Equal(v, "burkay")

	k, v = hash.indexOf(1)

	t.Equal(k, "surname")
	t.Equal(v, "durdu")

	k, v = hash.indexOf(-1)

	t.Equal(k, "")
	t.Equal(v, "")

	k, v = hash.indexOf(2)

	t.Equal(k, "")
	t.Equal(v, "")
}

func (t *HashSuite) TestHash_Get() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")

	val := hash.get("surname")

	t.Equal(val, "durdu")

	val = hash.get("surnames")

	t.Equal(val, "")
}

func (t *HashSuite) TestHash_GetFirstAndLast() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")
	hash.Add("school", "KTU")

	k, v := hash.first()

	t.Equal(k, "name")
	t.Equal(v, "burkay")

	k, v = hash.last()

	t.Equal(k, "school")
	t.Equal(v, "KTU")

	hash.clear()

	k, v = hash.first()

	t.Equal(k, "")
	t.Equal(v, "")
	return
}

func (t *HashSuite) TestHash_GetLength() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")
	hash.Add("school", "KTU")

	length := hash.length()

	t.Equal(length, 3)

	hash.Add("age", "27")

	length = hash.length()

	t.Equal(length, 4)
}

func (t *HashSuite) TestHash_Remove() {
	hash := New[string, string]()
	hash.Add("name", "burkay")
	hash.Add("surname", "durdu")
	hash.Add("school", "KTU")

	hash.Remove("school")

	t.Equal(hash.length(), 2)
	t.Equal(hash.get("school"), "")

	hash.Remove("surname")

	t.Equal(hash.length(), 1)
	t.Equal(hash.get("surname"), "")

	hash.Remove("name")

	t.Equal(hash.length(), 0)
	t.Equal(hash.get("name"), "")
}

func TestHashSuite(t *testing.T) {
	suite.Run(t, new(HashSuite))
}
