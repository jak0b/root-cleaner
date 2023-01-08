package utils

import (
  "os"
  "log"
  "fmt"
  "strings"
)

var ProtectedFiles = []string{"dwn", "img", "mnt", "shr", "vid", "doc", "inf", "not", "pkg", "vms",
                            "dsk", "mus", "prg", "wrk"}

func directoryContent(path string) {
  dir, err := os.Open(path)
  if err != nil { log.Fatalf("not such directory %s", err) }
  defer dir.Close()

  directoryFiles, _ := dir.ReadDir(0)

  for _, entry := range directoryFiles {
    entryName := entry.Name()
    if !IsHiddenEntry(entryName) && !InProtectedFiles(entryName) {
      fmt.Println(entryName)
    }
  }
}

func InProtectedFiles(elem string) bool {
  for _, prtFile := range ProtectedFiles {
    if elem == prtFile {
      return true
    }
  }
  return false
}

func IsHiddenEntry (elem string) bool {
    return strings.HasPrefix(elem, ".")
}
