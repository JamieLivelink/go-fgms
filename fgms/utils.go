
package fgms

import(
	"fmt"
	"time"

	"github.com/FreeFlightSim/go-fgms/message"
)

//= Returns an int64 with epoch (should be UTC ?)
func Now() int64{
	//time.UTC().unix() ??
	return time.Now().Unix()
} 

// pete FAIL FAIL FAIL's 
func GetProtocolVersionString() string {
	//return "1.1"
	major := message.PROTOCOL_VER >> 16
	minor := message.PROTOCOL_VER & 0xffff
	return fmt.Sprintf("%d.%d", major, minor) 
}


// TODO There has Got to be a better way
// Should the bytes be a pointer ?
func BytesToString(bites []byte) string{
	for n, b := range bites {
		if b == 0 {
			return string(bites[:n])
		}
	}
	return string(bites[:])
}

