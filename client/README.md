Create the .proto file with the services that are required

Then we use grpc.Dial mentioning the grpc address, use empty credentials that disables TLS, 

We create a NewUserClient which is made for us by the go-code-gen

Then we create a context to pass to the Service. The context is always passed on to the service.
