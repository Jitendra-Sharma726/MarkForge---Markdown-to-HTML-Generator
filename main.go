package main 

import (
  "bufio"
  "fmt"
  "html"
  "os"
  "strings"
)

//processMarkdown reads line-by-line and applies formatting rules
func processMarkdown(scanner *bufio.Scanner, file *os.File) {
  inList := false

  for scanner.scan()
  //Clean up spacing and sanitize input to prevent XSS attacks
  rawText := strings.Trim
