package main

import (
	"fmt"
	"github.com/vmarkovtsev/licenseng"
	"os"
	"sort"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: licenseng <path>")
		os.Exit(1)
	}
	licenses, err := licenseng.InvestigateProjectLicenses(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	licenseStrs := make([]string, 0, len(licenses))
	for key, val := range licenses {
		licenseStrs = append(licenseStrs, fmt.Sprintf("%-20s\t%.1f%%", key, val*100))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(licenseStrs)))
	for _, s := range licenseStrs {
		fmt.Println(s)
	}
}
