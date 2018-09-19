package commonprefix

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var ret string

	/*
		fmt.Println("strs", strs)
		fmt.Println("len, cap of strs", len(strs),cap(strs))
		fmt.Println("ret", ret)

		for i:=0; i<len(strs); i++{
			fmt.Println("i, str", i,strs[i], strs[i][1:])
		}
	*/

	var shortstridx int

	var strlen = len(strs[0])
	shortstridx = 0

	for i, v := range strs {
		if strlen > len(v) {
			strlen = len(v)
			shortstridx = i
		}
	}

	fmt.Println("shortstridx, str", shortstridx, strs[shortstridx])

	return ret
}
