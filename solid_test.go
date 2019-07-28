package gon3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

var baseURIPathSolid = "https://solid.community"

func TestSolidPositive(t *testing.T) {
	verbosity := 1
	for _, testName := range positiveParserSolidTests {
		basePath := "./tests/solid/"
		ttlFile := basePath + testName + ".ttl"
		ntFile := basePath + testName + ".out"
		baseURI := baseURIPathSolid + testName + ".ttl"
		// ttlFd, err := os.Open(ttlFile)
		ttlFd, err := ioutil.ReadFile(ttlFile)
		if err != nil {
			t.Fatalf("Error reading file %s", ttlFile)
		}
		ttlReader := bytes.NewReader(ttlFd)
		ntFd, err := ioutil.ReadFile(ntFile)
		if err != nil {
			t.Fatalf("Error reading file %s", ntFile)
		}
		ntReader := bytes.NewReader(ntFd)
		if verbosity > 0 {
			fmt.Printf("\nStarting test %s\n", testName)
		}
		ttlGraph, err := NewParser(baseURI).Parse(ttlReader)
		if err != nil {
			t.Fatalf("Test %s failed. Error parsing %s\n(%s)", testName, ttlFile, err)
		}
		ntGraph, err := NewParser(baseURI).Parse(ntReader)
		if err != nil {
			t.Fatalf("Test %s failed. Error parsing %s\n(%s)", testName, ntFile, err)
		}
		if !ntGraph.IsomorphicTo(ttlGraph) {
			t.Fatalf("Test %s failed. Graphs not isomorphic.\nttl graph:\n%s\nnt graph:\n%s", testName, ttlGraph, ntGraph)
		}
		if verbosity > 0 {
			fmt.Printf("Test %s passed.\n", testName)
		}
	}
}

func TestSolidPositiveNoIso(t *testing.T) {
	verbosity := 1
	for _, testName := range positiveParserSolidTestsNoIso {
		basePath := "./tests/solid/"
		ttlFile := basePath + testName + ".ttl"
		baseURI := baseURIPathSolid + testName + ".ttl"
		ttlBytes, err := ioutil.ReadFile(ttlFile)
		if err != nil {
			t.Fatalf("Error reading file %s", ttlFile)
		}
		if verbosity > 0 {
			fmt.Printf("\nStarting test %s\n", testName)
		}
		ttlReader := bytes.NewReader(ttlBytes)
		_, err = NewParser(baseURI).Parse(ttlReader)
		if err != nil {
			t.Fatalf("Test %s failed. Error parsing %s\n(%s)", testName, ttlFile, err)
		}
		if verbosity > 0 {
			fmt.Printf("Test %s passed.\n", testName)
		}
	}
}

func TestSolidNegative(t *testing.T) {
	verbosity := 0
	for _, testName := range negativeParserSolidTests {
		testFile := "./tests/solid/" + testName
		baseURI := baseURIPathSolid + testName + ".ttl"
		b, err := ioutil.ReadFile(testFile)
		if err != nil {
			t.Fatalf("Error reading test file %s", testFile)
		}
		br := bytes.NewReader(b)
		if verbosity > 0 {
			fmt.Printf("\nStarting test %s\n", testName)
		}
		p := NewParser(baseURI)
		_, err = p.Parse(br)
		if err == nil {
			t.Fatalf("Test %s failed: %s", testName, err)
		}
		if verbosity > 0 {
			fmt.Printf("Test %s passed.\n", testName)
		}
	}
}

var negativeParserSolidTests []string = []string{
}

var positiveParserSolidTests []string = []string{
}

var positiveParserSolidTestsNoIso []string = []string{
	"test-00",
	"test-01",
	"test-02",
	"test-03",
	"test-04",
	"test-05",
	"test-06",
	"test-07",
	"test-08",
	"test-09",
	"test-10",
}
