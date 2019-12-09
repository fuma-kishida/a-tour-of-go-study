package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (a MyReader) Read(rb []byte) (n int, e error) {
    ///@alt for i, _ := range rb { rb[i] = 'A'; n++; }
    for n, e = 0, nil; n < len(rb); n++ {
        rb[n] = 'A'
    }
    return
}

func main() {
    reader.Validate(MyReader{})
}