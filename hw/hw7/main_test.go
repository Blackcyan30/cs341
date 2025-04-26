package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// setupSampleMap reads sample.txt and populates passwordMap.
func setupSampleMap(t *testing.T) {
	passwordMap = make(map[string]EntrySlice)
	f, err := os.Open("sample.txt")
	if err != nil {
		t.Fatalf("opening sample.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) != 3 {
			t.Fatalf("invalid sample line: %s", scanner.Text())
		}
		entry := Entry{site: parts[0], user: parts[1], password: parts[2]}
		passwordMap[parts[0]] = append(passwordMap[parts[0]], entry)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("scanning sample.txt: %v", err)
	}
}

// helper to simulate stdin via os.Pipe()
func withStdin(input string, fn func()) error {
	// create pipe
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	defer r.Close()

	// backup and replace Stdin
	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	// write input and close writer
	_, err = w.Write([]byte(input))
	w.Close()
	if err != nil {
		return err
	}

	// invoke function that reads from stdin
	fn()
	return nil
}

func TestInitializeMap(t *testing.T) {
	setupSampleMap(t)

	exCount := len(passwordMap["example.com"])
	if exCount != 4 {
		t.Errorf("expected 4 entries for example.com; got %d", exCount)
	}

	sbCount := len(passwordMap["sitebuilder.xyz"])
	if sbCount != 1 {
		t.Errorf("expected 1 entry for sitebuilder.xyz; got %d", sbCount)
	}
}

func TestAddEntry(t *testing.T) {
	setupSampleMap(t)
	input := "newsite.com newuser newpass\n"

	// first addition
	err := withStdin(input, addEntry)
	if err != nil {
		t.Fatalf("setup stdin failed: %v", err)
	}
	entries, ok := passwordMap["newsite.com"]
	if !ok {
		t.Fatal("newsite.com key not created in map")
	}
	if len(entries) != 1 {
		t.Errorf("expected 1 entry for newsite.com; got %d", len(entries))
	}
	if entries[0].user != "newuser" || entries[0].password != "newpass" {
		t.Errorf("unexpected entry for newsite.com: %+v", entries[0])
	}

	// duplicate addition should not change count
	err = withStdin(input, addEntry)
	if err != nil {
		t.Fatalf("setup stdin failed: %v", err)
	}
	if len(passwordMap["newsite.com"]) != 1 {
		t.Errorf("duplicate addition should not change count; got %d", len(passwordMap["newsite.com"]))
	}
}

func TestRemoveEntry(t *testing.T) {
	// remove specific user
	setupSampleMap(t)
	err := withStdin("example.com user01\n", removeEntry)
	if err != nil {
		t.Fatalf("setup stdin failed: %v", err)
	}
	for _, e := range passwordMap["example.com"] {
		if e.user == "user01" {
			t.Error("user01 was not removed from example.com slice")
		}
	}

	// remove site when only one entry
	setupSampleMap(t)
	err = withStdin("sitebuilder.xyz\n", removeEntry)
	if err != nil {
		t.Fatalf("setup stdin failed: %v", err)
	}
	if _, ok := passwordMap["sitebuilder.xyz"]; ok {
		t.Error("sitebuilder.xyz should be deleted when removing sole entry")
	}

	// attempt to remove site with multiple users: should leave intact
	setupSampleMap(t)
	err = withStdin("example.com\n", removeEntry)
	if err != nil {
		t.Fatalf("setup stdin failed: %v", err)
	}
	if len(passwordMap["example.com"]) != 4 {
		t.Errorf("example.com should remain unchanged; expected 4, got %d", len(passwordMap["example.com"]))
	}
}
