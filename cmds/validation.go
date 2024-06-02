package cmds

import "log"

func checkRequiredFlags(errMessage string, flags ...string) {
	for _, flag := range flags {
		if flag == "" {
			log.Fatal(errMessage)
		}
	}
}
