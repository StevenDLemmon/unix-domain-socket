The intent of this project was to setup two processes communicating by means of a UNIX domain socket with a UDP connection.  This was developed and tested on WSL 2.


Setup:

Run server.go using "go run server.go".

In another terminal, run client.go using "go run client.go".


You should now see the output on the server terminal, which is a ndJSON with every process and the current memory that process is using.  The memory displayed is in kilobytes.  If you want, you can keep running the client over and over again until you shut down the server.
