package gon3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

const baseUriPath = "http://www.w3.org/2001/sw/DataAccess/df1/tests/"

func TestW3Positive(t *testing.T) {
	verbosity := 1
	for _, testName := range positiveParserW3Tests {
		basePath := "./tests/w3/"
		ttlFile := basePath + testName + ".ttl"
		ntFile := basePath + testName + ".out"
		baseURI := baseUriPath + testName + ".ttl"
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

func TestW3PositiveNoIso(t *testing.T) {
	verbosity := 0
	for _, testName := range positiveParserW3TestsNoIso {
		basePath := "./tests/w3/"
		ttlFile := basePath + testName + ".ttl"
		baseURI := baseUriPath + testName + ".ttl"
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

func TestW3Negative(t *testing.T) {
	verbosity := 0
	for _, testName := range negativeParserW3Tests {
		testFile := "./tests/w3/" + testName
		baseURI := baseUriPath + testName + ".ttl"
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

var negativeParserW3Tests []string = []string{
	"bad-00.ttl",
	"bad-01.ttl",
	"bad-02.ttl",
	"bad-03.ttl",
	"bad-04.ttl",
	"bad-05.ttl",
	"bad-06.ttl",
	"bad-07.ttl",
	"bad-08.ttl",
	"bad-09.ttl",
	"bad-10.ttl",
	"bad-11.ttl",
	"bad-12.ttl",
	"bad-13.ttl",
	// "bad-14.ttl",
}

var positiveParserW3Tests []string = []string{
  "rdf-schema",
	// "rdfq-results",
	"rdfs-namespace",
	"test-00",
	"test-01",
	"test-02",
	"test-03",
	"test-04",
	// "test-05",
	"test-06",
	// "test-07",
	// "test-08",
	"test-09",
	"test-10",
	"test-11",
	"test-12",
	"test-13",
	"test-14",
	"test-15",
	"test-16",
	"test-17",
	"test-18",
	"test-19",
	"test-20",
	"test-21",
	"test-22",
	"test-23",
	"test-24",
	"test-25",
	"test-26",
	"test-27",
	// "test-28",
	// "test-29",
	"test-30",
}

var positiveParserW3TestsNoIso []string = []string{
	"manifest-bad",
	"manifest",
}
