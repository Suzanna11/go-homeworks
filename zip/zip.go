package main
 
import (
   "archive/zip"
   "log"
   "os"
)
 
func main() {
 
   outFile, err := os.Create("test.zip")
   if err != nil {
      log.Fatal(err)
   }
   defer outFile.Close()
 
   // Create a new zip archive.
   w := zip.NewWriter(outFile)
 
   // Add some files to the archive.
   var files = []struct {
      Name, Body string
   }{
      {"file1.txt", "This is a file1"},
      {"file2.txt", "This is a file1"},
   }
   for _, file := range files {
      f, err := w.Create(file.Name)
      if err != nil {
         log.Fatal(err)
      }
      _, err = f.Write([]byte(file.Body))
      if err != nil {
         log.Fatal(err)
      }
   }
 
   err = w.Close()
   if err != nil {
      log.Fatal(err)
   }
}