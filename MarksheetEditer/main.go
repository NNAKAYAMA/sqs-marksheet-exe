package main

//go:generate statik --src static/SourceEditor.jar
import (
	"fmt"
	"io"
	"os"
	"os/exec"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func Exists(filename string) bool {
	fmt.Println(filename)
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	appdata := os.Getenv("APPDATA")
	dirName := "sqs"
	fileName := "SourceEditor.jar"
	if (!Exists(fmt.Sprintf("%s\\%s",appdata,dirName))){
		os.Mkdir(fmt.Sprintf("%s\\%s",appdata,dirName), 0777)
	}
	file, _ := os.OpenFile(fmt.Sprintf("%s\\%s\\%s", appdata, dirName, fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if fileInfo, err := file.Stat(); err == nil {
		if fileInfo.Size() == 0 {
			FS, _ := fs.New()
			f, _ := FS.Open("/" + fileName)
			_, err = io.Copy(file, f)
			fmt.Println("OverWrite File")
		}
	} else {
		FS, _ := fs.New()
		f, _ := FS.Open("/" + fileName)
		_, err = io.Copy(file, f)
		fmt.Println("Write File  New")
	}
	file.Close()
	exec.Command("javaw", "-jar", fmt.Sprintf("%s\\%s\\%s", appdata, dirName, fileName)).Start()
}
