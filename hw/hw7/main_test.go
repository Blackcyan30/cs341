package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// sampleData holds the entries from sample.txt
const sampleData = `example.com user01 Pass1234
mysite.dev jane_doe RedFish#99
example.com johndoe testPASS2023
dummyapp.org testuser LetMeIn!
sitebuilder.xyz alice89 S1tebuildR
fakesite.io dev_guy pass_word1
example.com betaTester Beta$2024
mysite.dev user_test1 Test456
mocksite.co hello_world qwerty!@
dummyapp.org coder123 C0deC0de
testcloud.app cloudguy CloudyDay99
fakebank.net userbank $BankSecure
maildemo.com inboxme mailbox_2024
dummyapp.org streamerx passPASS123
example.com portalUser WelCome1!
mysite.dev safe_user Safe4U2024
betaweb.dev betaOnly b3ta_test
tryapp.net demo_user TrymeNow
gamespotter.co gamer001 playGame$
datatester.com jsonlover JSONlove88
example.com admin_demo Adm1n123
signupnow.dev signup1 Signup123
clickandgo.net testerGuy clicknGO!
uiuxworld.com ux_queen designQueen9
mockupplace.org mockuser mockU99!
samplezone.dev zoneUser ZoneIt!2025
codespace.net coderGirl CodeSp@ce
netalpha.org alphaTester Alphaz#
mysite.dev hubUser Hub_Life21
devconsole.io consoleAdmin Cons0le4U
trydemo.dev trial_user trial_n_error
startuphub.org founderx hustleHard2024
example.com serverAdmin adminSERV3R
testpage.io pageViewer viewerOnly
browserdemo.com surfTester surfwave!!
quickmock.net quickie Quick123
emailtest.org mailtest mail4test
chatbotdemo.dev chattie ch@tbox
learnhub.co student1 learn@now
dummyapp.org api_user API_KEY_789
loginzone.net logme l0gMe1n
fakeprofile.org profake fak3it123
uidev.xyz ui_dev UIxUX2024
trythisapp.io appuserx trytryagain
samplestore.dev buyer001 B!Gpurchase
devplayground.org playgroundDev PlayPlay09
mocksite.co siteTester M0ck@site
webmockup.dev designerx des1gnerPASS
localhost:3000 localUser rootlocal
fakeintranet.org corpUser corpSecure55
`

// helper to populate passwordMap with sampleData
func setupSampleMap(t *testing.T) {
	passwordMap = make(map[string]EntrySlice)
	r := bufio.NewReader(strings.NewReader(sampleData))
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			t.Fatalf("invalid sample line: %s", line)
		}
		entry := Entry{site: parts[0], user: parts[1], password: parts[2]}
		passwordMap[parts[0]] = append(passwordMap[parts[0]], entry)
	}
	if err := s.Err(); err != nil {
		t.Fatalf("scanning sample data: %v", err)
	}
}

func TestInitializeMap(t *testing.T) {
	setupSampleMap(t)
	// example.com appears 4 times in sampleData
	exCount := len(passwordMap["example.com"])
	if exCount != 4 {
		t.Errorf("expected 4 entries for example.com; got %d", exCount)
	}
	// sitebuilder.xyz appears once
	sbCount := len(passwordMap["sitebuilder.xyz"])
	if sbCount != 1 {
		t.Errorf("expected 1 entry for sitebuilder.xyz; got %d", sbCount)
	}
}

func TestAddEntry(t *testing.T) {
	setupSampleMap(t)
	// prepare to add a new entry via addEntry()
	input := "newsite.com newuser newpass\n"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = strings.NewReader(input)
	addEntry()
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
	// test duplicate does not increase slice
	os.Stdin = strings.NewReader(input)
	addEntry()
	entries2 := passwordMap["newsite.com"]
	if len(entries2) != 1 {
		t.Errorf("duplicate addition should not change count; got %d", len(entries2))
	}
}

func TestRemoveEntry(t *testing.T) {
	setupSampleMap(t)
	// remove specific user
	input := "example.com user01\n"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = strings.NewReader(input)
	removeEntry()
	for _, e := range passwordMap["example.com"] {
		if e.user == "user01" {
			t.Error("user01 was not removed from example.com slice")
		}
	}
	// remove site when only one entry
	// pick sitebuilder.xyz which has exactly one
	os.Stdin = strings.NewReader("sitebuilder.xyz\n")
	removeEntry()
	if _, ok := passwordMap["sitebuilder.xyz"]; ok {
		t.Error("sitebuilder.xyz should be deleted when removing sole entry")
	}
	// attempt to remove site with multiple users: should leave intact
	setupSampleMap(t)
	os.Stdin = strings.NewReader("example.com\n")
	removeEntry()
	if len(passwordMap["example.com"]) != 4 {
		t.Errorf("example.com should remain unchanged; expected 4, got %d", len(passwordMap["example.com"]))
	}
}
