package main

import "fmt"

type Input struct {
	FileName string
}

type OutPut struct {
	FileName string
}

type Reader interface {
	Read(input *Input) string
}

type XMLReader struct {
}

func (reader *XMLReader) Read(input Input, output *OutPut) error {
	fmt.Println("xml reader ...")
	output.FileName = input.FileName
	return nil
}
