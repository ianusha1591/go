package main

import (
    "fmt"
    "html/template"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/ec2"
"github.com/aws/aws-sdk-go/aws/credentials" 
 "github.com/aws/aws-sdk-go/aws/session"

    "log"
    "net/http"
//"/root/.aws/credentials"


"os"

 "encoding/json"
)
type Data struct {
    Name string
}



func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("create.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("insname:", r.Form["insname"])
name:= r.FormValue("insname")
data  := &Data{name}
b, err :=json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
f, err := os.OpenFile("loginoutput.go", os.O_APPEND|os.O_WRONLY, 0644)

 f.Write(b)

sess, err := session.NewSession(&aws.Config{
    Region:   aws.String("us-west-2"),
    Credentials: credentials.NewSharedCredentials("", "default"),
 })
print(sess)
svc := ec2.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
    // Specify the details of the instance that you want to create.
    runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
        // An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
        ImageId:      aws.String("ami-e7527ed7"),
        InstanceType: aws.String("t2.micro"),
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
    })

    if err != nil {
        log.Println("Could not create instance", err)
        return
    }

    log.Println("Created instance", *runResult.Instances[0].InstanceId)
        
}  
}   

func main() {
//    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/ins", login)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

