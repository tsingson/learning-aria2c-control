package main

import (
	"flag"
	"fmt"
)

/*
	parser.add_argument('gid', help='gid of the download in aria2')
    parser.add_argument('num_files', help='number of files')
    parser.add_argument('target_file', help='name of first file on filesystem')
*/

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	args := flag.Args()
	for i := 0; i < len(args); i++ {
		fmt.Println(i, args[i])
	}

}
