package main

import (
	"flag"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/gocarina/gocsv"
	"github.com/kondroid00/sample-server-2022/package/errorcode"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := flag.String("dir", ".", "target directory")
	packageName := flag.String("package", "main", "package name")
	flag.Parse()

	dirPath, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatalln(err)
	}

	errorCodes, err := read(filepath.Join(dirPath, "list.csv"))
	if err != nil {
		log.Fatalln(err)
	}
	f := build(*packageName, errorCodes)
	fmt.Printf("%#v", f)
	if err := f.Save(filepath.Join(dirPath, "list.go")); err != nil {
		log.Fatalln(err)
	}
}

func read(filePath string) ([]*errorcode.ErrorCode, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	errorCodes := make([]*errorcode.ErrorCode, 0)
	if err := gocsv.UnmarshalFile(file, &errorCodes); err != nil {
		return nil, err
	}

	return errorCodes, nil
}

func write(filePath string, f *File) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	if err := f.Render(file); err != nil {
		return err
	}
	return nil
}

func build(packageName string, errorCodes []*errorcode.ErrorCode) *File {
	f := NewFile(packageName)
	errorCodePkg := "github.com/kondroid00/sample-server-2022/package/errorcode"
	f.ImportName(errorCodePkg, "errorcode")
	f.Var().Op("(")
	for _, v := range errorCodes {
		if v == nil {
			continue
		}
		f.Id(strings.ToUpper((*v).Code)).
			Op("=").
			Qual(errorCodePkg, "New").
			Call(
				Lit((*v).Code),
				Lit((*v).Message),
			)
	}
	f.Op(")")
	return f
}
