package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	// panic
	// 	// can do anything
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	// panic
	// 	// can do anything

	// 	panic(err)
	// }

	// fmt.Println(fileInfo.Name())         // output will be the name of the file name
	// fmt.Println(fileInfo.Size())         // output will be the size of the file
	// fmt.Println(fileInfo.ModTime())      // output will be the modification time of the file
	// fmt.Println(fileInfo.IsDir())        // output will be the directory of the file
	// fmt.Println(fileInfo.Mode())         // output will be the mode of the file
	// fmt.Println(fileInfo.Mode().Perm())  // output will be the permission of the file
	// fmt.Println(fileInfo.Mode().IsDir()) // output will be the directory of the file (check is it file or folder)
	// fmt.Println(fileInfo.Mode().Perm())  // output will be the permission of the file

	// defer f.Close() // close the file
	// read file and store buffer (buffer is a temporary storage into ram)
	// buff := make([]byte, 12)
	// // read function reads the file and store the data into buffer
	// d, err := f.Read(buff)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < d; i++ {
	// 	fmt.Println("data ", d, string(buff[i]))
	// }

	// fmt.Println("data ", d, string(buff))

	// ! most simple way

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	// log the error
	// 	// panic
	// 	// can do anything
	// 	panic(err)
	// }
	// fmt.Println(string(data))
	// readFile use when file size is small
	// it load into full file into memory
	// and return error when memory is full

	// ! read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()
	// // read dir
	// fileInfo, err := dir.ReadDir(-1)
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// * how to create a file

	// file, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// 	fmt.Println("error creating file", err)
	// }

	// defer file.Close()
	// file.WriteString("hi golang")
	// byte := []byte(" hello bro learn go lang")
	// file.Write(byte)

	// * how to copy data from one file to another file

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()
	destFile, err := os.Create("example3.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()

		// copy data from source file to destination file
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}
	writer.Flush()
	fmt.Println("retuen to new file successfully")
}
