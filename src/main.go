package main

import (
    "github.com/sethvargo/go-githubactions"
  )
  
  func main() {
    val := githubactions.GetInput("showGithubVars")
    output := ""
    if val == "true" {

      //newLineEsc := "%0A"

      output += " one "

      val2 := os.Getenv("GITHUB_WORKFLOW")
      output += val2 

      output += " two "

      //val2 = os.Getenv("GITHUB_RUN_ID")
      //output += val2 + newLineEsc      

    }
    githubactions.SetOutput("myOutput",output)
  }