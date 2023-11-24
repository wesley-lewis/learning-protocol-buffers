First we create a Server and use the unimplementedUserServiceServer(generated for us)

Then we implement the Service for the above struct

We listen on a port, we get a listener

Then we create a grpc.Server(I think it's a default server)

register the server we made above using the RegisterUserServiceServer()

then we use the server.Serve method and pass in the tcp listener
