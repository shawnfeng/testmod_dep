package testmod_dep

import (
    "fmt"
    "github.com/shawnfeng/testmod"
)


func Hidep(name string, arg string) string {
	fmt.Println(testmod.Hi("testmod_dep mod roberto 1", "AAA", "bbbb"))
	return fmt.Sprintf("testmod_dep Hi dep, %s arg:%s", name, arg)
}

