package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func runMain(args ...string) string {
	oldArgs := os.Args
	oldOut := os.Stdout
	os.Args = append([]string{"wdmatch"}, args...)
	r, w, _ := os.Pipe()
	os.Stdout = w
	done := make(chan string, 1)
	go func() {
		var b bytes.Buffer
		io.Copy(&b, r)
		done <- b.String()
	}()
	func() {
		defer func() {
			os.Args = oldArgs
			os.Stdout = oldOut
			w.Close()
		}()
		main()
	}()
	return <-done
}

func TestLevel1_Wdmatch_Visible(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"example_1_same_numbers", []string{"123", "123"}, "123\n"},
		{"example_2_match", []string{"faya", "fgvvfdxcacpolhyghbreda"}, "faya\n"},
		{"example_3_no_match", []string{"faya", "fgvvfdxcacpolhyghbred"}, ""},
		{"example_4_not_enough_letters", []string{"error", "rrerrrfiiljdfxjyuifrrvcoojh"}, ""},
		{"example_5_with_spaces", []string{"quarante deux", "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "}, "quarante deux\n"},
		{"example_6_no_arguments", []string{}, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := runMain(tc.args...)
			if got != tc.expected {
				t.Fatalf("wdmatch %q printed %q; want %q", tc.args, got, tc.expected)
			}
		})
	}
}
