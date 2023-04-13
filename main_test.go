package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
    "bytes"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {

	oldOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Could create a pipe %v", err.Error())
	}

	os.Stdout = w

	prompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("incorrect promt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()
	out, _ := io.ReadAll(r)

	os.Stdout = oldOut

	expectedOutput := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "

	if string(out) != expectedOutput {
		t.Errorf("incorrect output: expected %q but got %q", expectedOutput, string(out))
	}
}

func TestCheckNumbers(t *testing.T) {
	testCases := []struct {
		input string
		want  string
		done  bool
	}{
		{"2\n", "2 is a prime number!", false},
		{"3\n", "3 is a prime number!", false},
		{"4\n", "4 is not a prime number because it is divisible by 2!", false},
		{"-1\n", "Negative numbers are not prime, by definition!", false},
		{"abc\n", "Please enter a whole number!", false},
		{"q\n", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tc.input))
			res, done := checkNumbers(scanner)

			if res != tc.want {
				t.Errorf("checkNumbers(%s) = %s; want %s", tc.input, res, tc.want)
			}

			if done != tc.done {
				t.Errorf("checkNumbers(%s) = %t; want %t", tc.input, done, tc.done)
			}
		})
	}
}

func TestReadUserInput(t *testing.T) {
	doneChan := make(chan bool)
  
	var stdin bytes.Buffer
  
	stdin.Write([]byte("1\nq\n"))
  
	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
  }