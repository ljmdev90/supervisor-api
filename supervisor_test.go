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

func TestStartProcessGroup(t *testing.T) {
	list, err := sv.StartProcessGroup("cat")
	if err != nil {
		t.Errorf("StartProcessGroup error: %s", err)
	} else {
		t.Log(len(list), list)
	}
}

func TestStopProcessGroup(t *testing.T) {
	list, err := sv.StopProcessGroup("cat")
	if err != nil {
		t.Errorf("StopProcessGroup error: %s", err)
	} else {
		t.Log(len(list), list)
	}
}
