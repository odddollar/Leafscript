package main

import (
	"Leafscript/supporting"
	"archive/zip"
	"fmt"
	"github.com/akamensky/argparse"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// create argparse
	parser := argparse.NewParser("Leafscript", "The backend of the Leafscript programming language")
	run := parser.String("r", "run", &argparse.Options{Help: "The path to the .lfs file to run"})
	debug := parser.Flag("d", "debug", &argparse.Options{Default: false, Help: "Include to run program in debug mode"})
	compile := parser.String("b", "build", &argparse.Options{Help: "The path to the .lfs file to compile to .exe"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
	}

	// check if compiling or running
	if *run != "" && *compile == ""{
		// ensure correct file format is used
		if strings.Contains(*run, ".lfs") {
			// parse file to language
			fileLines := supporting.ParseFile(*run)

			if *debug {
				supporting.Lex(fileLines, true)
			} else {
				supporting.Lex(fileLines, false)
			}
		} else {
			fmt.Println("Invalid file format. Please use the format .lfs")
		}
	} else if *run == "" && *compile != ""{
		// ensure correct file format is used
		if strings.Contains(*compile, ".lfs") {
			// create zip file to copy to and zip writer
			zipFile, err := os.Create("LeafscriptTemp.zip")
			if err != nil {
				log.Fatal(err)
			}
			defer zipFile.Close()

			writer := zip.NewWriter(zipFile)
			defer func() {
				writer.Close()
				err = exec.Command("PackageFile", "-run", "run.bat", "-autotemp", "-sm", "-sp", "LeafscriptTemp.zip").Run()
				if err != nil {
					log.Fatal(err)
				}

				//rename file
				_ = os.Rename("LeafscriptTemp.exe", strings.ReplaceAll(*compile, ".lfs", ".exe"))
			}()

			// copy .lfs file to zip
			if err = AddFileToZip(writer, *compile); err != nil {
				log.Fatal(err)
			}

			// create run.bat, copy to zip then remove original
			_ = ioutil.WriteFile("run.bat", []byte(fmt.Sprintf("@ECHO OFF\nleafscriptTemp -r %v", *compile)), 0755)
			if err = AddFileToZip(writer, "run.bat"); err != nil {
				log.Fatal(err)
			}
			_ = os.Remove("run.bat")

			// create copy of leafscript exe, move to zip then delete
			progName, _ := os.Executable()
			path, _ := os.Getwd()
			Copy(progName, path+"/leafscriptTemp.exe")
			if err = AddFileToZip(writer, "leafscriptTemp.exe"); err != nil {
				log.Fatal(err)
			}
			_ = os.Remove("leafscriptTemp.exe")
		} else {
			fmt.Println("Invalid file format. Please use the format .lfs")
		}
	}
}

func Copy(src string, dst string) {
	// Read all content of src to data
	data, _ := ioutil.ReadFile(src)
	// Write data to dst
	_ = ioutil.WriteFile(dst, data, 0644)
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}