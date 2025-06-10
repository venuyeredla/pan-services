package learngo

import "testing"

func TestSort(t *testing.T) {
	SortPrint()
	t.Log("Sucess")
}

func TestFunctions(t *testing.T) {
	closureExample()
	TestCallBack()
	ReturnFunc()("venugopal")

	testRegex()
}
