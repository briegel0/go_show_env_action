package main

import (
    "github.com/sethvargo/go-githubactions"
  )
  
  func main() {
    val := githubactions.GetInput("showGithubVars")
    if val != "true" {
      githubactions.Fatalf("missing 'showGithubVars'")
    }
    githubactions.setEnv("myOutput","billKey=XXX")
  }