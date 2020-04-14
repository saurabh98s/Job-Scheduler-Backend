# Handling Jobs/tasks on Server

### Solution Description
----------------------------------



What the task declared was to devise a method where a user has the option<br>
to stop or abort any long running task if he/she finds some error due to any reason<br>
So, the Api given in the repo is designed so that a simple request to the desired job id<br>
will stop the task when it is running. The API is designed in GoLang, which is as far<br>
I have worked is an amazing language for backend tasks.

#### COMMANDS FOR API:


You can test it by running the server and then using cURL from another terminal:
```
run the go-docker.exe
in another terminal:

curl "localhost:8080/start?id=1" 
curl "localhost:8080/start?id=2"
curl "localhost:8080/stop?id=1"
```
In the server terminal you should see the following:
```
Doing job id 1
Doing job id 1
Doing job id 1
Doing job id 2
Doing job id 1
Doing job id 2
Cancelling job id 1
Doing job id 2
Doing job id 2
Doing job id 2
...
```

### How it Works?
For understanding the context in golang package there are 2 concepts that you should be familiar with.<br>
* goroutine
* channels <br>
These two concepts which are used to build a concurrent API helped me in the development

#### Resources used:
* http://p.agnihotry.com/post/understanding_the_context_package_in_golang/
* https://golang.org/pkg/context/

### What gave me trouble:<br>
Dockerising the image and sending it, there was an error an error which i couldn't solve<br>
```
docker: Error response from daemon: OCI runtime create failed: container_linux.go:345: starting container process caused "exec: \"/app/go-docker\": stat /app/go-docker:
 no such file or directory": unknown.
```
As i use Docker Toolbar, my Windows version isn't suitable for docker environment.<br>
Nevertheless,it was great working on the project
## Alternative to Docker
Binaries built in Golang can be run freely across any of the platforms without installing anything , so i am uploading the executable
file, so that you may run and then verify the code.
