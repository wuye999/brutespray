package brute

import (
	"strconv"
	"time"

	"github.com/jlaffaye/ftp"
)

func BruteFTP(host string, port int, user, password string, timeout time.Duration) (bool, bool) {
	conn, err := ftp.Dial(host+":"+strconv.Itoa(port), ftp.DialWithTimeout(timeout))
	if err != nil {
		return false, false
	}
	defer conn.Quit()

	err = conn.Login(user, password)
	if err != nil {
		return false, true
	}

	return true, true
}
