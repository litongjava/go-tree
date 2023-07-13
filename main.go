package main

import (
  "flag"
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

func main() {
  excludeDirs := flag.String("I", "", "Exclude directories (comma-separated)")
  flag.Parse()

  directory := "."
  args := flag.Args()
  if len(args) > 0 {
    directory = args[0]
  }
  var excludeList []string
  if *excludeDirs != "" {
    excludeList = strings.Split(*excludeDirs, ",")
  }

  walkFunc := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      fmt.Printf("error accessing a path %q: %v\n", path, err)
      return err
    }

    relPath, _ := filepath.Rel(directory, path)

    if info.IsDir() {
      for _, excludeDir := range excludeList {
        if strings.Contains(relPath, excludeDir) {
          return filepath.SkipDir
        }
      }

      dirCount := countDir(relPath)
      indent := "│   "
      if dirCount > 0 {
        indent = strings.Repeat("│   ", dirCount-1)
        fmt.Printf("%s├── %s\n", indent, filepath.Base(path))
      } else {
        fmt.Printf("%s\n", filepath.Base(path))
      }
    } else {
      dirCount := countDir(relPath)
      indent := strings.Repeat("│   ", dirCount-1)
      fmt.Printf("%s├── %s\n", indent, filepath.Base(path))
    }
    return nil
  }

  err := filepath.Walk(directory, walkFunc)
  if err != nil {
    fmt.Printf("error walking the path %v: %v\n", directory, err)
  }
}

func countDir(path string) int {
  if path == "." {
    return 0
  } else {
    count := strings.Count(path, string(os.PathSeparator))
    return count + 1
  }
}
