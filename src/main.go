package main

import (
    "os"
    "github.com/sethvargo/go-githubactions"
  )
  
  func main() {
    val := githubactions.GetInput("showGithubVars")
    output := ""
    if val == "true" {

      newLineEsc := "%0A"

      output += " GITHUB_WORKFLOW="

      val2 := os.Getenv("GITHUB_WORKFLOW")
      output += val2 

      output += newLineEsc

      output += " GITHUB_RUN_ID="      

      val2 = os.Getenv("GITHUB_RUN_ID")
      
      output += val2 + newLineEsc      

    }
    githubactions.SetOutput("myOutput",output)
  }