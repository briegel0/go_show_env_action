package main

import (
    "os"
    "strings"
    "fmt"
    "github.com/sethvargo/go-githubactions"
  )
  
  func main() {
    val := githubactions.GetInput("showGithubVars")
    output := ""

    newLineEsc := "\n"

    if val == "true" {

      // newLineEsc := "%0A"
      //newLineEsc := "\n"

      output += " CI="

      val2 := os.Getenv("CI")
      output += val2 

      output += newLineEsc

      output += " GITHUB_WORKFLOW="      

      val2 = os.Getenv("GITHUB_WORKFLOW")
      
      output += val2 + newLineEsc      

      output += " GITHUB_RUN_ID="      

      val2 = os.Getenv("GITHUB_RUN_ID")
      
      output += val2 + newLineEsc        

      output += " GITHUB_RUN_NUMBER="      

      val2 = os.Getenv("GITHUB_RUN_NUMBER")
      
      output += val2 + newLineEsc
      
      output += " GITHUB_ACTION="      

      val2 = os.Getenv("GITHUB_ACTION")
      
      output += val2 + newLineEsc    
      
      output += " GITHUB_ACTIONS="      

      val2 = os.Getenv("GITHUB_ACTIONS")
      
      output += val2 + newLineEsc        
      
      output += " GITHUB_ACTOR="      

      val2 = os.Getenv("GITHUB_ACTOR")
      
      output += val2 + newLineEsc     
      
      output += " GITHUB_REPOSITORY="      

      val2 = os.Getenv("GITHUB_REPOSITORY")
      
      output += val2 + newLineEsc     
      
      output += " GITHUB_EVENT_NAME="      

      val2 = os.Getenv("GITHUB_EVENT_NAME")
      
      output += val2 + newLineEsc     
      
      
      output += " GITHUB_EVENT_PATH="      

      val2 = os.Getenv("GITHUB_EVENT_PATH")
      
      output += val2 + newLineEsc     
      
      output += " GITHUB_WORKSPACE="      

      val2 = os.Getenv("GITHUB_WORKSPACE")
      
      output += val2 + newLineEsc     
      
      output += " GITHUB_SHA="      

      val2 = os.Getenv("GITHUB_SHA")
      
      output += val2 + newLineEsc     
      
      output += " GITHUB_REF="      

      val2 = os.Getenv("GITHUB_REF")
      
      output += val2 + newLineEsc           

      output += " GITHUB_HEAD_REF="      

      val2 = os.Getenv("GITHUB_HEAD_REF")
      
      output += val2 + newLineEsc   

      output += " GITHUB_BASE_REF="      

      val2 = os.Getenv("GITHUB_BASE_REF")
      
      output += val2 + newLineEsc   
      
      output += " GITHUB_SERVER_URL="      

      val2 = os.Getenv("GITHUB_SERVER_URL")
      
      output += val2 + newLineEsc   
      
      output += " GITHUB_API_URL="      

      val2 = os.Getenv("GITHUB_API_URL")
      
      output += val2 + newLineEsc   
      
      output += " GITHUB_GRAPHQL_URL="      

      val2 = os.Getenv("GITHUB_GRAPHQL_URL")
      
      output += val2 + newLineEsc         


    }

    // any app env's ? 
    valApp := githubactions.GetInput("listOfAppEnvVars")
    if len(valApp) > 0 { 
       keyArray := strings.Split(valApp, ",")
       for i, s := range keyArray {

        fmt.Println(i, s)

        output += s + "="      

        val2 := os.Getenv(s)
        
        output += val2 + newLineEsc   

       }       
    }


    githubactions.SetOutput("myOutput",output)
  }