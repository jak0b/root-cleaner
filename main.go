package main

import (
  "os"
  "github.com/jak0b/root-cleaner/pkg/utils"
)



func main() {
  homeEnv := os.Getenv("HOME")
  utils.directoryContent(homeEnv)
}
