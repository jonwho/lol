package lol

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	cli, err := NewClient("test_key")
	if err != nil {
		t.Error(err)
	}
	expected := "test_key"
	actual := cli.Token
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
	expected = "na"
	actual = cli.Region
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	cli, err = NewClient("to_be_overwritten", WithToken("foobar"))
	if err != nil {
		t.Error(err)
	}
	expected = "foobar"
	actual = cli.Token
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	cli, err = NewClient("to_be_overwritten", WithRegion("mynewregion"))
	if err != nil {
		t.Error(err)
	}
	expected = "mynewregion"
	actual = cli.Region
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}
