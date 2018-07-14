package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	// create an app
	app := cli.App("cp", "Copy files around")

	// Here's what differentiates mow.cli from other CLI libraries:
	// This line is not just for help message generation.
	// It also validates the call to reject calls with less than 2 arguments
	// and split the arguments between SRC or DST
	app.Spec = "[-r] SRC... DST"

	var (
		// declare the -r flag as a boolean flag
		recursive = app.BoolOpt("r recursive", false, "Copy files recursively")
		// declare the SRC argument as a multi-string argument
		src = app.StringsArg("SRC", nil, "Source files to copy")
		// declare the DST argument as a single string (string slice) arguments
		dst = app.StringArg("DST", "", "Destination where to copy files to")
	)

	// Specify the action to execute when the app is invoked correctly
	app.Action = func() {
		fmt.Printf("Copying %v to %s [recursively: %v]\n", *src, *dst, *recursive)
		// err := CopyFile(os.Args[1], os.Args[2])
		// if err != nil {
		// 	fmt.Printf("CopyFile failed %q\n", err)
		// } else {
		// 	fmt.Printf("CopyFile succeeded\n")
		// }
	}

	// Invoke the app passing in os.Args
	app.Run(os.Args)
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	return
}
