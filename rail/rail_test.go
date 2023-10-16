package rail

import "testing"

func TestEncode(t *testing.T) {
	testcases := []struct {
		name      string
		plaintext string
		key       []int
		expected  string
	}{
		{"test0", "WEAREDISCOVEREDFLEEATONCE", []int{3, 1}, "WECRLTEERDSOEEFEAOCAIVDEN"},
		{"test1", "Hello, World!", []int{3, 1}, "Hoo!el,Wrdl l"},
		{"test2", "Hello, World!", []int{3, 2}, "Herlo!lWd o,l"},
		{"test_homework",
			"CRYPTOLOGY IS THE PRACTICE AND STUDY OF TECHNIQUES FOR SECURE COMMUNICATION IN THE PRESENCE OF THIRD PARTIES" +
				" CALLED ADVERSARIES",
			[]int{4, 5},
			"CT HEGCONAYQNENMDD FRLMUET EIIEE ISTOPCSRT CC FLRSIPVI TO IR  TH  OSFAAHNCOTC  DSRDRN SAIUY CDRTEAPHA CUPEEAEIOEOLSIETNEES YRRU"},
	}

	for _, e := range testcases {
		got := Encode(e.plaintext, e.key)
		if got != e.expected {
			t.Errorf("Test case %s failed. Expected %v, but got %v\n", e.name, e.expected, got)
		}
	}
}

func TestDecode(t *testing.T) {
	testcases := []struct {
		name       string
		ciphertext string
		key        []int
		expected   string
	}{
		{"test0", "WECRLTEERDSOEEFEAOCAIVDEN", []int{3, 1}, "WEAREDISCOVEREDFLEEATONCE"},
		{"test1", "Hoo!el,Wrdl l", []int{3, 1}, "Hello, World!"},
		{"test2", "Herlo!lWd o,l", []int{3, 2}, "Hello, World!"},
		{"test_homework",
			"TPSNIONFRMHANOIREEOIBTSEAKLAPSEHISOBPEGTBRQREPTTMHRTHTTAWE" +
				"AACTFVAECAOLHANSEEKFHALOITUEEAICNLEYOLTOLEPADFKATATSJMSIINSH" +
				"ROCTITRIEEAKYNHYUEOTTSTATSIIRSARERUYUEDPCLETEROINEIYEHC",
			[]int{3, 3},
			"THERAILFENCECIPHERISANEASYTOAPPLYTRANSPOSITIONCIPHERTHATJUMBLESUPTHEORDEROFTHELETTERSOFAMESSAGEINAQUICKCONVENIENTWAYITALSOHASTHESECURITYOFAKEYTOMAKEITALITTLEBITHARDERTOBREAK"},
	}

	for _, e := range testcases {
		got := Decode(e.ciphertext, e.key)
		if got != e.expected {
			t.Errorf("Test case %s failed. Expected %v, but got %v\n", e.name, e.expected, got)
		}
	}
}
