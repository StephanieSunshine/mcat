package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "os"
)

func main() {
  flag.Parse()

  for _, file := range flag.Args() {
    fmt.Fprintln(os.Stderr, file)

    fh, err := os.Open(file)
    if err != nil {
      panic("Could not open file: "+file)
    }
    defer fh.Close()

    reader := bufio.NewReader(fh)
    
    for {
      line, _, err := reader.ReadLine()
      if err != nil {
        if err == io.EOF { break }
        panic("A real error happend with file: "+file)
      }
      fmt.Println(string(line))
    }
  }
}
