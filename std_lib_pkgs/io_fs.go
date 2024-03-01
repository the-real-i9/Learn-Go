package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

/*
- Package fs defines basic interfaces to a file system.
- It provides functions to access and operate on the file system's tree
*/
func io_fs_pkg() {
	// access the contents (file names ("-" prefix) and folder names ("d" prefix)) of a directory in a file system(FS)
	// os.DirFs("/*") turns dir into a file system for which /* is it root
	dir_entries, err := fs.ReadDir(os.DirFS("/home"), "i9")
	if err != nil {
		log.Fatal(err)
	}

	var dir_str []string
	for _, dir_entry := range dir_entries {
		dir_str = append(dir_str, fs.FormatDirEntry(dir_entry)) // FormatDirEntry() for human readability
	}

	bt, err := json.MarshalIndent(dir_str, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	io.Discard.Write(bt)

	// retrieve the names of all direnctories and files  matching the glob pattern in the file system
	matches, err := fs.Glob(os.DirFS("/home/i9/Code/Go"), "*/*")
	if err != nil {
		log.Fatal(err)
	}

	mat_bt, err := json.MarshalIndent(matches, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	io.Discard.Write(mat_bt)

	// read the file specified in a file system
	file_data, err := fs.ReadFile(os.DirFS("/home/i9"), ".bash_history")
	if err != nil {
		log.Fatal(err)
	}

	io.Discard.Write(append(file_data, '\n'))

	// get the stat of a file in the file system
	// it returns a FileInfo type from which you can retrive specific information
	// Name(), Size(), Mode(), ModTime(), IsDir(), Sys()
	fs_stat, err := fs.Stat(os.DirFS("/home/i9"), ".bashrc")
	if err != nil {
		log.Fatal(err)
	}

	// print the file information in a human-readable format
	// "Mode() Size() ModeTime() Name()"
	f_info := fs.FormatFileInfo(fs_stat) + "\n"
	fmt.Fprintln(io.Discard, f_info)

	/* ------------ */
	// test if a given path name is valie for use in a call to os.Open()
	// pathname must neither start nor end with "/"
	fmt.Fprintln(io.Discard, fs.ValidPath("home/i8/.bashrc"))

	/* ------------ */
	// Walk the file system tree (all nodes) starting from the specified root
	// if the root is ".", then it starts from the file system as its root
	// if the root is a specified folder in the file system, it starts from that folder as the root
	fs.WalkDir(os.DirFS("/home/i9/Code/Go"), "src", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(fs.FormatDirEntry(d)) // either a folder name ("d" prefix) or a file name ("-" prefix)
		return err
	})
}
