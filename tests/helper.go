package tests

import (
	"testing"

	"github.com/jpcaissy/gotcha/worklist"
	"github.com/stretchr/testify/assert"
)

// For all testdata checks
// a) returned value is of Type errInFlow (we are expecting flows)
// b) the number of reported flows is equal to the expected number
func check(testData []*testDataStruct, t *testing.T) {
	for _, td := range testData {
		//t.Log(td)
		//t.Logf("path: %s | sourceFilesFlag: %s | ssf: %s | allpkgs: %t |ptr: %t \n", path, td.sourceFile, taintFile, false, true)
		err := worklist.DoAnalysis(path, td.sourceFile, taintFile, false, "", true)
		t.Log(err)
		if td.expectedFlows > 0 {
			if assert.NotNil(t, err) {
				if assert.IsType(t, (*worklist.ErrInFlows)(nil), err) {
					e, _ := err.(*worklist.ErrInFlows)
					getFlows := e.NumberOfFlows()
					t.Logf("expectedFlows: %d | getFlows: %d \n", td.expectedFlows, getFlows)
					assert.Equal(t, td.expectedFlows, getFlows, "get: "+e.Error())
				}
			}
		} else {
			assert.Nil(t, err)
		}
	}
}

var taintFile = "./sourcesAndSinksTest.txt"
var path = "github.com/jpcaissy/gotcha"

type testDataStruct struct {
	sourceFile    []string
	expectedFlows int
}
