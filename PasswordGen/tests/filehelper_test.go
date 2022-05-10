package tests

import (
	"testing"

	"passwordgen/filehelper"
)

func TestFileHelper(t *testing.T) {

	name := "../JsonFile/newfile.json"
	
	filehelper.GetFileInfo(name)

}
