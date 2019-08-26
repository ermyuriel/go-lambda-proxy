# go-lambda-proxy
A simple Go AWS Lambda utility package and template for API Gateway accesible implementations

# Use

```go
func f(ctx *context.Context) (interface{},error){
    //Define the flow you want the API to execute here. You will find the request function inside the Context with a "request" key
}


func main(){
    lambda.Start(lambdaproxy.ProxyFunction(f))
}
```

That's it. 

# Why

This repository was made mostly just to make my life easier and development of single purpose lambda APIs faster.

Any suggestions or corrections are very welcome.
