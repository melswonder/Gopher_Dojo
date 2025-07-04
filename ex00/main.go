package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("error: invalid argument")
        os.Exit(1)
    }

    dirPath := os.Args[1]

    if err := checkDir(dirPath); err != nil {
        fmt.Printf("error: %v\n", err)
        os.Exit(1)
    }

    if err := processDirectory(dirPath); err != nil {
        fmt.Printf("エラーが発生しました: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("ファイルの処理が正常に完了しました")
}

func checkDir(dirPath string) error {
    fileInfo, err := os.Stat(dirPath)
    if err != nil {
        if os.IsNotExist(err) { //中身が存在するか
            return fmt.Errorf("%s: no such file or directory", dirPath)
        }
        return err
    }
    if !fileInfo.IsDir() {
        return fmt.Errorf("%s is not a directory", dirPath)
    }
    return nil
}

func processDirectory(dirPath string) error {
    return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error { //ディレクトリを再帰的に
        if err != nil {
            return err
        }
        
        // ディレクトリはスキップ
        if info.IsDir() {
            return nil
        }

        // 有効な画像ファイルかチェック
        if !isValidImageFile(info.Name()) {
            fmt.Printf("error: %s is not a valid file\n", path)
            return fmt.Errorf("invalid file found: %s", path)
        }

        // JPG/JPEGファイルの場合は変換処理
        ext := strings.ToLower(filepath.Ext(path))
        if ext == ".jpg" || ext == ".jpeg" {
            fmt.Printf("変換中: %s\n", path)
            // TODO: 実際の変換ロジックをここに実装
        }

        return nil
    })
}

func isValidImageFile(filename string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}
