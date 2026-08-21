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
  if strings.HasPrefix(line, "- ") {
    if !inList {
      file.WriteString("<ul>\n")
      inList= true
    }
    content := strings.TrimPrefix(line, "- ")
    file.WriteString("  <li>" + parseInline(content) + "</li>\n")
    continue
  }

  //Close list if paragraph starts
  if inList {
    file.WriteString("</ul>\n")
    inList = false
  }

  //Handle standard paragraphs
  file.WriteString("<p> + parseInline(line) + "</p>\n")
}
                   
//Close trailing list
if inList {
  file.WriteString("</ul>\n")

//check for scan failures
if err:= scanner.Err(); err != nil {
  fmt.Printf("Error reading file : %v\n", err)
  }
}

//parseInline handles bold, italic, and inline code formatting
func parseInline(text string) string {
  //Step 1: Parse bold(**text**)
  parts := strings.Split(text, "**")
  for i := 1; i < len(parts); i += 2 {
    parts[i] = "<strong>" + parts[i] + "</strong>"
}
text = strings.Join(parts, "")

//Step 2: Parse italic (*text*)
parts = strings.Split(text, "*")
for i := 1; i < len(parts); i +=2 {
   parts[i] = "<em>" + parts[i] + "<em>"
}
text = strings.Join(parts, "")

//Step 3: Parse inline code (`code`)
parts = strings.Split(text, "`")
for i := 1; i < len(parts); i +=2 {
     parts[i] = "<code>" + parts[i] + "</code>"
 }
 text = strings.Join(parts, "")
  
return text
}

func main() {
  fmt.Println("=== Markdown to HTML Generator ===")

  //1.Open input file
  inputFile, err := os.Open("input.md:")
  if err != nil {
    fmt.Printf("Error opening file :%v\n", err)
    os.Exit(1)
    }
    defer inputFile.Close()
  
//2. Create output file
 outputFile, err := os.Create("output.html")
 if err != nil {
    fmt.Printf("Error creating file: %v\n", err)
    os.Exit(1)
   }
defer outputFile.Close()

//3. Process the file
writeHeader(outputFile)

scanner := bufio.NewScanner(inputFile)
processMarkdown(scanner, outputFile)
  
writeFooter(outputFile)
  
fmt.Println("Success! HTML generated at output.html")
}

//writeHeader outputs HTML boilerplate and CSS
func writeHeader(file *os.File) {
  file.WriteString("<!DOCTYPE html>\n<html>\n<head>\n")
  file.WriteString("<title>Generated Page</title>\n")
  file.WriteString("<style>\n")
  file.WriteString("     body { font-family : 'Segoe UI', Tahoma, sans-serif; max-width: 800px; margin: 40px auto; line-height: 1.6; color: #333; padding: 0 20px; }\n")
  file.WriteString("     h1 { border-bottom: 2px solid #eee; padding-bottom: 0.3em; }\n")
  file.WriteString("     h2 { border-bottom: 1px solid #eee; padding-bottom: 0.2em; }\n")
  file.WriteString("     code { background: #f4f4f4; padding: 2px 6px; border-radius: 4px; font-family: monospace; font-size: 0.9em; }\n")
  file.WriteString("     ul { padding-left: 2em; }\n")
  file.WriteString("     li {margin: 4px 0; }\n")
  file.WriteString("</style>\n")
  file.WriteString("</head>\n<body>\n")

  // writeFooter closes HTML tags
  func writeFooter(file *os.File) {
    file.WriteString("</body>\n</html>\n")
  }










  










  
  















                   

                   
















                   







                                          





  















                                  
