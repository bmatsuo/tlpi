/*
A low-level implementation of the `copy` command.
*/
package main

import (
	. "arguments"
	. "check"
	"syscall"
)

type OpenFlag int

const (
	Create    OpenFlag = syscall.O_CREAT
	Exclusive OpenFlag = syscall.O_EXCL
	Truncate  OpenFlag = syscall.O_TRUNC
	Append    OpenFlag = syscall.O_APPEND
	WriteOnly OpenFlag = syscall.O_WRONLY
	ReadOnly  OpenFlag = syscall.O_RDONLY
	ReadWrite OpenFlag = syscall.O_RDWR
)

type FilePerm uint32

const (
	UserRead   FilePerm = syscall.S_IRUSR
	UserWrite  FilePerm = syscall.S_IWUSR
	UserExec   FilePerm = syscall.S_IWUSR
	GroupRead  FilePerm = syscall.S_IRGRP
	GroupWrite FilePerm = syscall.S_IWGRP
	GroupExec  FilePerm = syscall.S_IWGRP
	OtherRead  FilePerm = syscall.S_IROTH
	OtherWrite FilePerm = syscall.S_IWOTH
	OtherExec  FilePerm = syscall.S_IWOTH
)

func main() {
	args := Arguments()

	infile, args, err := args.String("in-file", nil)
	Check(err)
	outfile, args, err := args.String("out-file", nil)
	Check(err)

	infd, err := syscall.Open(infile, int(ReadOnly), 0)
	Check(err)
	defer func() { Check(syscall.Close(infd)) }()

	flags := Create | WriteOnly | Truncate
	perms := 0 |
		UserRead | UserWrite |
		GroupRead | GroupWrite |
		OtherRead | OtherWrite
	outfd, err := syscall.Open(outfile, int(flags), uint32(perms))
	Check(err)
	defer func() { Check(syscall.Close(outfd)) }()

	bufSize := 1024
	buf := make([]byte, bufSize)
	var n int
	for {
		n, err = syscall.Read(infd, buf)
		Check(err)
		if n == 0 {
			break
		}
		_, err = syscall.Write(outfd, buf[:n])
		Check(err)
	}
}
