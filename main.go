package main

import (
	"cps/models"
	"easygf"
	"os"
)

func main() {
	s := easygf.NewServer()
	initArgs(s)
}

func initArgs(s *easygf.Server) {
	args := os.Args
	for _, v := range args {
		if v == "--syncdb" {
			s.SyncDB(models.GetModels())
			os.Exit(0)
		}
	}

}
