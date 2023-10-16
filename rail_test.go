package rail

import "testing"

func TestEncode(t *testing.T) {
	got := Encode("WEAREDISCOVEREDFLEEATONCE", 3)
	if got != "WECRLTEERDSOEEFEAOCAIVDEN" {
		println("error #1\n\tgot !" + got + "!\n\twant!WECRLTEERDSOEEFEAOCAIVDEN!")
	}
	got = Encode("Hello, World!", 3)
	if got != "Hoo!el,Wrdl l" {
		println("error #2=>" + got)
	}

}

func TestDecode(t *testing.T) {
	got := Decode("WECRLTEERDSOEEFEAOCAIVDEN", 3)
	if got != "WEAREDISCOVEREDFLEEATONCE" {
		println("error #1=>" + got)
	}
	got = Decode("Hoo!el,Wrdl l", 3)
	if got != "Hello, World!" {
		println("error #2=>" + got)
	}

}
