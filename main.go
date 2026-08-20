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
  rawText := strings.TrimSpace(scanner.Text())
  line := html.EscapeString(rawText)

  // Skip empty lines and close open lists
  if line == "" {
    if inList {
        file.WriteString("</ul>\n")
        inList = false
      }
    continue
  }

  //Parse heading (H6 down to H1 to avoid false matches)
  if strings.HasPrefix(line, "###### ") {
    content := strings.TrimPrefix(line, "###### ")
    file.WriteString("<h6>" + parseInLine(content) + "</h6>\n")
    continue
    }
  if strings.HasPrefix(line, "##### ") {
    content := strings.TrimPrefix(line, "##### ")
    file.WriteString("<h5" + parseInline(content) + "</h5>\n")
    continue
  }
  if strings.HasPrefix(line, "#### ") {
    content := strings.TrimPrefix(line, "#### ")
    file.WriteString("<h4>" + parseInline(content) + "</h4>\n")
    continue
  }
  if strings.HasPrefix(line, "### ") {
    content := strings.TrimPrefix(line, "### ")
    file.WriteString("<h3>" + parseInline(content) + "</h3>\n")
    continue
  }
  if strings.HasPrefix(line, "## "){
    content := strings.TrimPrefix(line, "## ")
    file.WriteString("<h2>" + parseInline(content) + "</h2>\n")
    continue
  }
  if strings.HasPrefix(line, "# ") {
    content := strings.TrimPrefix(line, "# ")
    file.WriteString("<h1>" + parseInline(content) + "</h1>\n")
    continue
  }

  //Handle list items








                                          





  















                                  
