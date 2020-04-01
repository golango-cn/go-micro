package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T)  {

	//sp := string(filepath.Separator)

	fmt.Println(string(filepath.Separator))
	fmt.Println(filepath.Abs(filepath.Dir(".")))

	os.Chdir("/Users/ningzi/go/src/go-micro/bb")
	file, err := os.Create("a.txt")

	if err != nil {
		fmt.Println(err)
	}

	var bf bytes.Buffer
	bf.WriteString("12321312321")

	bf.WriteTo(file)


	//appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))
	//
	//pt := filepath.Join(appPath, "conf")
	//os.Chdir(appPath)

	//fmt.Println(pt)
}

func TestConfig02(t *testing.T)  {

	err := os.MkdirAll("./a", os.ModePerm)
	//os.Chdir("./a")
	os.Create("file.txt")
	fmt.Println(err)//成功创建文件file.txt，返回nil
	//os.Chdir("../")
	//err = os.RemoveAll("./a")
	//fmt.Println(err)//成功删除目录a，返回nil

}