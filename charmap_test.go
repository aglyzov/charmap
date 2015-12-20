package charmap

import "testing"

var TEST_DATA = map[string][]byte{
	"CP1251" : []byte(
		"\xd2\xe5\xf1\xf2\xee\xe2\xe0\xff \xf1\xf2\xf0\xee\xea\xe0 \xed\xe0 \xf0\xf3\xf1\xf1\xea\xee\xec",
	),
	"IBM866" : []byte(
		"\x92\xa5\xe1\xe2\xae\xa2\xa0\xef \xe1\xe2\xe0\xae\xaa\xa0 \xad\xa0 \xe0\xe3\xe1\xe1\xaa\xae\xac",
	),
	"iso-8859-5" : []byte(
		"\xc2\xd5\xe1\xe2\xde\xd2\xd0\xef \xe1\xe2\xe0\xde\xda\xd0 \xdd\xd0 \xe0\xe3\xe1\xe1\xda\xde\xdc",
	),
	"koi8r" : []byte(
		"\xf4\xc5\xd3\xd4\xcf\xd7\xc1\xd1 \xd3\xd4\xd2\xcf\xcb\xc1 \xce\xc1 \xd2\xd5\xd3\xd3\xcb\xcf\xcd",
	),
	"KOI8-U" : []byte(
		"\xf4\xc5\xd3\xd4\xcf\xd7\xc1\xd1 \xd3\xd4\xd2\xcf\xcb\xc1 \xce\xc1 \xd2\xd5\xd3\xd3\xcb\xcf\xcd",
	),
}
const EXPECTED = "Тестовая строка на русском"


func Test_ANY_to_UTF8_Works(t *testing.T) {
	for charset, data := range TEST_DATA {
		result, err := ANY_to_UTF8(data, charset)
		if err != nil {
			t.Errorf("%v", err)
		}
		if string(result) != EXPECTED {
			t.Errorf("Expected %#v, got %#v", EXPECTED, string(result))
		}
	}
}

func Test_CP1251_to_UTF8_Works(t *testing.T) {
	result := CP1251_to_UTF8(TEST_DATA["CP1251"])
	if string(result) != EXPECTED {
		t.Errorf("Expected %#v, got %#v", EXPECTED, string(result))
	}
}

func Test_CP866_to_UTF8_Works(t *testing.T) {
	result := CP866_to_UTF8(TEST_DATA["IBM866"])
	if string(result) != EXPECTED {
		t.Errorf("Expected %#v, got %#v", EXPECTED, string(result))
	}
}

func Test_KOI8R_to_UTF8_Works(t *testing.T) {
	result := KOI8R_to_UTF8(TEST_DATA["koi8r"])
	if string(result) != EXPECTED {
		t.Errorf("Expected %#v, got %#v", EXPECTED, string(result))
	}
}

