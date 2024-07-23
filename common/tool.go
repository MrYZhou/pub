package common

import (
	// . "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

func GetServer(s ServerConfig) *Cli {
	con, _ := Server(s.Host, s.User, s.Password)
	return con
}
