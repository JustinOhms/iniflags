package iniflags

import (
	"flag"
	"testing"
)

func TestRemoveTrailingComments(t *testing.T) {
	hash_commented := "v = v # test_comment"
	clean := removeTrailingComments(hash_commented)
	if clean != "v = v" {
		t.Fatalf("Supposed to get 'v = v ', got '%s'", clean)
	}
	colon_commented := "v = v ; test_comment"
	clean = removeTrailingComments(colon_commented)
	if clean != "v = v" {
		t.Fatalf("Supposed to get 'v = v ', got '%s'", clean)
	}

}

func TestBOM(t *testing.T) {
	args, ok := getArgsFromConfig("test_bom.ini")
	if !ok {
		t.Fail()
	}
	if len(args) != 1 {
		t.Fatalf("Unexpected number of args parsed: %d. Expected 1", len(args))
	}
	if args[0].Key != "bom" {
		t.Fatalf("Unexpected key name parsed: %q. Expected \"bom\"", args[0].Key)
	}
	if args[0].Value != "привет" {
		t.Fatalf("Unexpected value parsed: %q. Expected \"привет\"", args[0].Value)
	}
}

func TestUnquoteValue(t *testing.T) {
	val := "\"val#;\\\"\\n\"    # test\n"
	fixed_val, ok := unquoteValue(val, 0, "")
	if !ok || fixed_val != "val#;\"\n" {
		t.Fatalf("Value should be unquoted and stripped, got '%s'", fixed_val)
	}
}

func TestGetFlags(t *testing.T) {
	parsed = false
	Parse()
	missingFlags := getMissingFlags()
	if _, found := missingFlags["config"]; !found {
		t.Fatalf("'config' flag should be missing in tests")
	}
}

func TestGetArgsFromConfig(t *testing.T) {
	args, ok := getArgsFromConfig("test_config.ini")
	if !ok {
		t.Fail()
	}
	var checked_var0, checked_var1, checked_var2, checked_var3 bool
	for _, arg := range args {
		t.Log(arg.Key, arg.Value)
		if arg.Key == "var0" {
			if arg.Value != "val0" {
				t.Fatalf("Val of 'var0' should be 'val0', got '%s'", arg.Value)
			} else {
				checked_var0 = true
			}
		}
		if arg.Key == "var1" {
			if arg.Value != "val#1\n\\\"\nx" {
				t.Fatalf("Invalid val for var1='%s'", arg.Value)
			} else {
				checked_var1 = true
			}
		}
		if arg.Key == "var2" {
			if arg.Value != "1234" {
				t.Fatalf("Val of 'var2' should be '1234', got '%s'", arg.Value)
			} else {
				checked_var2 = true
			}
		}
		if arg.Key == "var3" {
			if arg.Value != "" {
				t.Fatalf("Val of 'var3' should be '', got '%s'", arg.Value)
			} else {
				checked_var3 = true
			}
		}
	}
	if !checked_var0 || !checked_var1 || !checked_var2 || !checked_var3 {
		t.Fatalf("Not all vals checked: args=[%v], %v, %v, %v, %v", args, checked_var0, checked_var1, checked_var2, checked_var3)
	}
}

func TestIsHttp(t *testing.T) {
	if !isHttp("http://example.com") {
		t.Fatalf("http://example.com should must be recognized as http path")
	}
	if !isHttp("hTtpS://example.com") {
		t.Fatalf("hTtpS://example.com should must be recognized as http path")
	}
}

var x = flag.String("x", "baz", "for TestSetConfigFile")

func TestSetConfigFile(t *testing.T) {
	parsed = false
	SetConfigFile("./test_setconfigfile.ini")
	Parse()
	if *x != "foobar" {
		t.Fatalf("Unexpected x=[%s]. Expected [foobar]", *x)
	}
}
