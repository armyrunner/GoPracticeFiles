package generator

import (
	"testing"
)


func TestGenertator(t *testing.T){
	tests := []struct{
		length    int
		minChar   int
		minNum    int
		minUpper  int
		minLower  int
	}{
		{length: 8,  minChar: 2, minNum: 1, minUpper: 3, minLower: 1},
		// {length: 10, minChar: 3, minNum: 3, minUpper: 2, minLower: 1},
		// {length: 16, minChar: 4, minNum: 3, minUpper: 3, minLower: 4},
	}
	
	for _, test := range tests{
		pass, _ := GeneratePassword(test.length, test.minChar, test.minNum, test.minUpper, test.minLower)
		if len(pass) != test.length{
			t.FailNow()
		}
	}

}