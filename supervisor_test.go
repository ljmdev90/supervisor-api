package supervisor_test

import (
	"testing"

	"github.com/ljmdev90/supervisor-api"
)

var sv *supervisor.Supervisor

func init() {
	var err error
	if sv, err = supervisor.New("http://localhost:9001/RPC2"); err != nil {
		panic(err)
	}
}

func TestListMethods(t *testing.T) {
	_, err := sv.ListMethods()
	if err != nil {
		t.Errorf("ListMethods error: %s", err)
	}
}
