package ntp

import (
	"fmt"

	"github.com/beevik/ntp"
)

var ntpServer = "0.beevik-ntp.pool.ntp.org"

func GetTime() error {
	time, err := ntp.Time(ntpServer)
	if err != nil {
		return err
	}

	fmt.Println(time)

	return nil
}
