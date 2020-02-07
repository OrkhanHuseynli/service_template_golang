## Docker <br/>
##### NOTE: Commands and instructions below are valid for Windows machines
#### 1. Basic commands
##### 1.1. Use the --rm flag to keep our filesystem clean after we stop our container. <br/>
Once we start our image with the following command it becomes a container <br/>

 `docker start -i --rm [image-name] sh`
##### 1.2 Start a sleeping container interactively <br/>
`docker start -i [container-id]`
* Type a linux command, i.e. `ls` to list all files/folders inside the current directory of the container
#### 2. Adding  Volumes
* bind mounting a host-directory in a container (`docker run -v /some/host/dir/:/container/path/`) uses the files that are
 present on the host. If the host-directory doesn't exist, a new, empty directory is created on the host and mounted in 
 the container (this will change in future, and an error is shown in stead)
* using a "nameless" volume (docker run -v /container/path) will create a new volume, and copy the contents of /container/path into that volume
* using a "named" volume (docker run -v somename:/container/path) will create a new volume, named "somename", or use 
the existing "somename" volume, and use the files that are present inside that volume. If the volume is newly created, 
it will be empty. <br/>
`docker run --rm -it -v C:/go-work/src/github.com/service_template/:go/src/service_template golang:alpine /bin/sh`

#### 3. Docker ports
Note: `pwd` is a source path for the app <br>

`docker run -it --rm -v $(pwd):/src -w /src -p 8080:8080 golang:alpine /bin/sh`
* Now you can run your app inside the docker simply by running `go run main.go` 
*The -w flag we are passing is to set the working directory that means that any command we
run in the container will be run inside this folder.
* We are not using alpine:latest, which is a lightweight
version of Linux, we are using golang:alpine, which is a version of Alpine with the most
recent Go tools installed.
* We will add the -p argument to specify the port. Like volumes,
this takes a pair of values separated by a colon (:). The first is the destination port on the
host that we would like to bind to the second is the source port on the Docker container to
which our application is bound.

#### 4. Docker networking

By default, Docker supports the following
network modes:
* bridge
* host
* none
* overlay

##### 4.1 Connect docker to the custom network
1. Create network <br/>
`docker network create testnetwork`<br/>
2. Attach your service to the network <br/>
`docker run --rm -it -v C:/go-work/src/github.com/service_template/:/go/src/mic_tem  -w /go/src/service_template/src --name server_a --network=testnetwork golang:alpine /bin/sh` <br/>
3. Run the service <br/>
`go run main.go` <br/>
4. Open a new terminal and do a curl request to the service in the specified network (i.e. "testnetwork") <br/>
`docker run --rm --network=testnetwork appropriate/curl:latest curl -i -XPOST server_a:8080/product -d "{\"product\":\"Nike\"}"` 

#### 5. Running the Dockerfile
1. Build the image: <br/>
`docker build ./ -t imagename`
2. Create container from the image <br/>
`docker run -p 8080:8080 imagename`
3. Do POST request to `localhost:8080/product` with the body `{"product": "some product"}`. <br/>
You should get the response: `{"message":"new product name: new product","date":"","id":"0"}`

NOTE: If you requests fail in Windows machines then you can check in the Oracle WM if guest ports are linked to the ones in the host. 
Go to VM's settings (usually with the general name "default") -> Network -> Advanced -> Port Forwarding -> Adds new port forwarding rule 
-> put `Host IP` 127.0.0.1 and `Host port` 8080, also add `Guest port` 8080

#### 6. docker-compose 
Inside `docker-compose.yml` file you can define services that will
make up your application. In example, we are describing two services:
our example code that we have just built and the second is a simple service that curls
this API.

To avoid conflicts when starting a stack, we can pass -p projectname to the docker-compose up 
command; this will prefix the name of any of our containers with the specified project name.

To remove any stopped container that you have started with dockercompose,
we can use the particular compose command rm and pass the -v argument to
remove any associated volumes: <br/>
`docker-compose rm -v`

#### Appendix

##### 1. Go Build
##### 1.1 go versions <=1.1
`CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o server` <br/>
* `CGO_ENABLED=0`  Cgo enables the creation of Go packages that call C code. By setting it to `0`, we make it unavailable. 
*  `GOARCH=386` This is pronounced “gore-ch”, and stands for Go Architecture. 
* `GOOS=linux GOARCH=386` The key-value pair starts with the `GOOS`, which in this example would be linux, referring to the `Linux OS`. The `GOARCH`
 here would be `386`, which stands for the `Intel 80386 microprocessor`.
 * `-ldflags` for details follow the [link](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications).
 * `-ldflags '-s'`, this argument passes the -s argument to the linker when we build the application and tells it to 
 statically link all dependencies.
 * `-o`  output
 * `-a` force rebuilding of packages that are already up-to-date
 
 Note: A **statically linked binary** has all of the required library code built in, so it's big but will run on just 
 about any system of the same type it was compiled on. If the binary is dynamically linked (uses shared libraries), all
 systems it is to run on must have a copy of all the required libraries. 
 For more info read ["Static and Dynamic Libraries"](https://www.geeksforgeeks.org/static-vs-dynamic-libraries/). 
 
 ##### 1.2 go versions > 1.1
`GOOS=linux GOARCH=386 go build -a -ldflags '-s' -o server` <br/>