package main

import(
	"fmt"
	"image"
	"errors"
	"os"
)

func file_check() error {
	path := os.Args[1]
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileInfo.IsDir(){
		return errors.New("指定されたパスはディレクトリです")
	}
	
	mode := fileInfo.Mode()
	if mode&0666 != 0666 {
		return errors.New("読み込みまたは書き込みの権限がありません")
	}
	return nil
}

func main(){
	if len(os.Args) < 2{
		fmt.Println("error: invalid argment")
		os.Exit(1)
	}

	if err := file_check(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	inputFile, err := os.Open(os.ARgs[1])
	if err != nil {
		fmt.Printf("入力ファイルのオープンに失敗しました: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	fmt.Println("ファイルの処理が正常に完了しました")
}