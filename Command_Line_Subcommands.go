package main

// Some command-line tools, like the go tool or git have many subcommands, each with its own set of flags.
// For example, go build and go get are two different subcommands of the go tool.
//  The flag package lets us easily define simple subcommands that have their own flags.

import (
	"flag"
	"fmt"
	"os"
)

// The easiest way to create a temporary file is by calling os.CreateTemp. It creates a file and opens it for reading and writing.
//  We provide "" as the first argument, so os.CreateTemp will create the file in the default location for our OS.
func main() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Display the name of the temporary file. On Unix-based OSes the directory will likely be /tmp.
	// The file name starts with the prefix given as the second argument to
	//  os.CreateTemp and the rest is chosen automatically to ensure that concurrent calls will always create different file names.
	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

// Clean up the file after we’re done. The OS is likely to clean up temporary files by itself after some time.
// We can write some data to the file.If we intend to write many temporary files, we may prefer to create a temporary directory.
//  os.MkdirTemp’s arguments are the same as CreateTemp’s, but it returns a directory name rather than an open file.
//  Now we can synthesize temporary file names by prefixing them with our temporary directory.

// output :
// $ ./command-line-subcommands foo -enable -name=joe a1 a2
// subcommand 'foo'
//   enable: true
//   name: joe
//   tail: [a1 a2]

//   $ ./command-line-subcommands bar -level 8 a1
// subcommand 'bar'
//   level: 8
//   tail: [a1]

//   $ ./command-line-subcommands bar -enable a1
// flag provided but not defined: -enable
// Usage of bar:
//   -level int
//         level
