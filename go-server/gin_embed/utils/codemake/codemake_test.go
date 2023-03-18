package reportno

import "testing"

func TestMakeNewCodeMaker(t *testing.T) {

	_, err := MakeNewCodeMaker("YC,YYMM,1231223423")
	if err == nil {
		t.Fail()
	}

}
