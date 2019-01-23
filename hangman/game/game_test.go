package game

import "testing"

func TestParseWord(t *testing.T) {
	input := []string{
		"asDfGh",
		"Asd fgh",
		"",
		"asd fgh!",
		"123",
		"asdf ghทดสอบ",
		"123 asdf ghjK 3",
	}
	expected := []string{
		"______",
		"___ ___",
		"",
		"___ ___!",
		"123",
		"____ __ทดสอบ",
		"123 ____ ____ 3",
	}

	if len(input) != len(expected) {
		t.Fatal("length of input != expected, please check test file")
	}

	for i := 0; i < len(input); i++ {
		output := parseWord(input[i])
		if output != expected[i] {
			t.Errorf("from input `%s`: expected `%s`, but get `%s`", input[i], expected[i], output)
		}
	}
}

func TestFormatDisplay(t *testing.T) {
	input := []string{
		"asDfGh",
		"Asd fgh",
		"",
		"asd fgh!",
		"_sD__h",
		"Asd __h",
		"_",
		"___",
	}
	expected := []string{
		"a s D f G h",
		"A s d   f g h",
		"",
		"a s d   f g h !",
		"_ s D _ _ h",
		"A s d   _ _ h",
		"_",
		"_ _ _",
	}

	if len(input) != len(expected) {
		t.Fatal("length of input != expected, please check test file")
	}

	for i := 0; i < len(input); i++ {
		output := formatDisplay(input[i])
		if output != expected[i] {
			t.Errorf("from input `%s`: expected `%s`, but get `%s`", input[i], expected[i], output)
		}
	}
}
