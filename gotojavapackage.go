package main

import (
	"fmt"
	"strings"
)

func main() {

	goPkgName := "github.com/golang/packages"
	
	javaPkgName := GoToJavaPackage(goPkgName)
	
	fmt.Println(goPkgName)
	fmt.Println(javaPkgName)
	
}

/* 
  Function converts go package to java package by performing reverse domain and / to . converstion 
  Assumes go package starts with domain name, in other cases, it returns blank string
*/
func GoToJavaPackage(goPkg string) (javaPkg string) {
	
	var slashIdx, dotIdx int
	
	dotIdx = strings.Index(goPkg, ".")
	
	if dotIdx > 1 {
	
		slashIdx = strings.Index(goPkg, "/")
	
		if dotIdx < slashIdx {
			javaPkg = goPkg[dotIdx + 1 :slashIdx] + "." + goPkg[:dotIdx] + goPkg[slashIdx:]
			javaPkg = strings.Replace(javaPkg , "/", ".", -1)
		}		
	
	}

	return
}
