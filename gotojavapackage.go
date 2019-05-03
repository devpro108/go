package main

import (
	"fmt"
	"strings"
)

func main() {

	goPkgName := "github.com/devpro108/go"
	
	javaPkgName := GoToJavaPackage(goPkgName)
	
	fmt.Println(goPkgName)
	fmt.Println(javaPkgName)
	
	goPkgName = ""

	goPkgName = JavaToGoPackage(javaPkgName)

	fmt.Println(goPkgName)
	
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
	
		if ( slashIdx - dotIdx ) > 1 {
			javaPkg = goPkg[dotIdx + 1 :slashIdx] + "." + goPkg[:dotIdx] + goPkg[slashIdx:]
			javaPkg = strings.Replace(javaPkg , "/", ".", -1)
		}		
	
	}

	return
}

/* 
  Function converts java package to go package by adjusting domain name and . to / converstion 
  Assumes java package starts with domain name, in other cases, it returns blank string
*/
func JavaToGoPackage(javaPkg string) (goPkg string) {
	
	var dotIdx1, dotIdx2 int
	
	dotIdx1 = strings.Index(javaPkg, ".")
	
	if dotIdx1 > 1 {
	
		dotIdx2 = strings.Index(javaPkg[dotIdx1 + 1:], ".")
		dotIdx2 = dotIdx2 + dotIdx1 + 1 
			
		if ( dotIdx2 - dotIdx1 ) > 1 {
		
			goPkg = strings.Replace(javaPkg[dotIdx2:], ".", "/", -1)
			goPkg = javaPkg[dotIdx1 + 1 :dotIdx2] + "." + javaPkg[:dotIdx1] + goPkg
		}		
	
	}

	return
}
