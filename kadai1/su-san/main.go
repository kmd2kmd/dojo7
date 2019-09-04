package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/gopherdojo/dojo7/kadai1/su-san/image"
)


func main() {

	inputExt := flag.String("i", "jpg", " extension to be converted ")
	outputExt := flag.String( "o", "png", " extension after conversion")

	// Usageメッセージ
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage : cmd [-i] [-o] target_dir")
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	targetDir := args[0]

	// ディレクトリがなければ知らせる
	if _, err := os.Stat(targetDir); err != nil {
		fmt.Printf("%v\n",  err)
		return
	}
	targetFiles := dirwalk(targetDir)

	convExts := image.NewConvExts(*inputExt, *outputExt)
	for _, f := range targetFiles {

		err := image.FmtConv(f, convExts)
		if err != nil {
			fmt.Println(f, err)
		}
	}
}

// dirwalk は与えられたディレクト以下のファイルパスをリストで返します
func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
