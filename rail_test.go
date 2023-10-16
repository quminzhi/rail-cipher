package rail

import "testing"

func TestEncode(t *testing.T) {
	got := Encode("WEAREDISCOVEREDFLEEATONCE", []int{3, 1})
	if got != "WECRLTEERDSOEEFEAOCAIVDEN" {
		println("error #1\n\tgot !" + got + "!\n\twant!WECRLTEERDSOEEFEAOCAIVDEN!")
	}
	got = Encode("Hello, World!", []int{3, 1})
	if got != "Hoo!el,Wrdl l" {
		println("error #2=>" + got)
	}

}

func TestDecode(t *testing.T) {
	got := Decode("WECRLTEERDSOEEFEAOCAIVDEN", []int{3, 1})
	if got != "WEAREDISCOVEREDFLEEATONCE" {
		println("error #1=>" + got)
	}
	got = Decode("Hoo!el,Wrdl l", []int{3, 1})
	if got != "Hello, World!" {
		println("error #2=>" + got)
	}
}
