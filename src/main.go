package main

import (
    "github.com/sethvargo/go-githubactions"
  )
  
  func main() {
    val := githubactions.GetInput("showGithubVars")
    output := ""
    if val == "true" {

      newLineEsc := "%0A"

      val2 := githubactions.GetInput("GITHUB_WORKFLOW")
      output += val2 + newLineEsc

      val2 = githubactions.GetInput("GITHUB_RUN_ID")
      output += val2 + newLineEsc      

    }
    githubactions.SetOutput("myOutput",output)
  }