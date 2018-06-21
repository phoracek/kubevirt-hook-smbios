package main

import (
	hook "github.com/phoracek/kubevirt-hook-smbios/pkg"
)

func main() {
	hook.NewHookServer().Run()
}
