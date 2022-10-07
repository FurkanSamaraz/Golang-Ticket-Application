# Golang-TaskApp

# Project Title
```
The project I developed is an automated ticket marketing algorithm.
This software is a simple backend architecture with no frontend.

* Pay attention to the notes.
```
## What Does It Do?
```
1- Developing new projects.

2- To manage the process quickly and without errors in long-term projects

3- Detecting and correcting the mistakes made thanks to its sophistication.

4- Working more systematically.

5- Analyzing reporting and business data.

6- Import databases.
```
## Getting Started

### Prerequisites

```
In order to run the project successfully, the requirements are golang version 1.18.2 
as programming language, a suitable IDE and postman or similar Api test tool for testing 
are required.
```

### Installing

```
After the Go installation, I will transfer what I did to set the correct GOROOT and GOPATH variables,
install 3rd party packages, and set up the project. I'm writing for MacOS, I think it can be solved 
with similar configurations in other UNIX distributions. First of all, I downloaded Go from this 
address. Then I installed the package, next, next, next (my favorite) It gets added to your Go PATH
along with the package installation, so you can run it to the command line to check it out.

- go version
You can write. Now we need to assign values for the GOROOT and GOPATH environment variables. 
It is useful to give brief information about these variables, GOROOT is used to reach the 
directory where Go is installed. For macOS, this value should be /usr/local/go. To check if 
this value is correct "which go" you can find out where Go is installed with the command. 
As a result, I got the following response "/usr/local/go/bin/go" based on this answer, 
I learned which directory the executable file is in. There is no problem, now the 
GOROOT value "/usr/local/go" I can assign as.
for this i open my profile file "nano $HOME/.bash_profile" I enter this line in it 
"export GOROOT=/usr/local/go".

Yes, now I need a directory where I will host Go projects, in this directory there will
be Go projects that I will create or clone, as well as the 3rd party packages I have 
downloaded in this directory. For this, I will open a directory on the desktop, you 
can edit it according to your wishes. I prefer to install on desktop.

-cd /Users/furkan/Desktop/
-mkdir GoProjects
-cd GoProjects
-mkdir src 
-mkdir pkg

In line with these commands, I created a folder called GoProjects on the desktop and 
I created 2 more folders called src and pkg. Now GOPATH, that is, the folder where my 
Go projects and 3rd party packages will be hosted, is ready. In order for Go to access 
the packages, I need to add them to the PATH. I open my profile file again "nano $HOME/.bash_profile"
I add the following line "export GOPATH=/Users/furkan/Desktop/GoProjects"

Yes, Go installation and environment variables are also completed, the packages I 
downloaded with go get and my own packages are working correctly in line with this 
article. In order for our variables to become permanent "source $HOME/.bash_profile" 
We enter the command.

for mod and sum.
-go mod init main.go
-go mod tidy


You can check the installation repo details.

https://github.com/FurkanSamaraz/GoKurulum

1- Go to the location of main.go

2- Download repository.

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"testing"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

3- go run main.go


```
## Running the project


```
1- First create a user by going to the "/users" extension in Postman.

2- In Postman go to "/ticket_options" extension and add ticket

3- Go to "/ticket/{id}" extension in postman and check by entering ticket id

4-Go to "/ticket_options/1/purchases" option in postman and buy tickets with the user.

```


## Diagrams
Check out the flowcharts of the APIs for the project;

![diagram_id](https://user-images.githubusercontent.com/92402372/187967417-fb2cc3c7-dc0b-4ae7-a3de-cc8fc38d87c5.png)

![diagram_purchases](https://user-images.githubusercontent.com/92402372/187967439-a8243173-d9cd-4189-ab95-8ac3b93b6d3d.png)

![diagram_options](https://user-images.githubusercontent.com/92402372/187967430-f9ee9533-2287-4276-87f4-d77c1e91db74.png)

![diagram_user](https://user-images.githubusercontent.com/92402372/187967446-5300547d-ccdf-4915-b94d-089b3dab6223.png)


## Used technologies

```
Technologies used in the project;

1- vscode

2- Golang

3- Golang fiber and gorilla library for APIs

4- httptest library for Golang tests

5- postgresql as database to store data

6- Postman to test API

7- Flow chart for drawing the project chart and ease of visual understanding

```
## Running the tests

```
_test.go tests were written for the files with the .go extension in the repository. 
You can run and check the tests by going to the appropriate file location with the 
'go test -v' command or by opening the _test.go file and clicking the play button 
next to the functions.

```

### And coding style tests

```
There are two kinds of test algorithms in the _test.go files in the repository. 
1 for Get requests. 2. Test algorithms for Post requests. httptest library is 
used for these test algorithms.

```


## Docker

```
Let's create our dockerfile and docker-compose.yml files.

Dockerfile is the file without extension that describes all the layers that 
exist to create a particular image image. Which Image will be used, which 
files it will contain, and which application will run with which parameters 
are written in the Dockerfile.

Docker Compose is in yaml file format. YAML is a data exchange format that 
can be easily read by humans and easily processed numerically and can be used
by all programming languages.

Open terminal in project location and dockerize.

-  docker-compose up -d --build

Note: "Error response from daemon: Ports are not available: listen tcp 0.0.0.0:5432: bind: address already in use" If you get the error, free the 5432 port. Don't forget to backup the DB. The db will be deleted when emptying the 5432 port.

- sudo lsof -nP -i4TCP:5432 | grep LISTEN
- sudo kill -9 <pid>

Dockerize it again, then restore the db you backed up and you can go to the APIs and use it.

Note: "http: panic serving 172.19.0.1:port dial tcp 127.0.0.1:5432: connect: connection refused" If you get an error, change the host address to db-service in the sql connection function in the code and specify it in the .yml file.

```
## Build With

```
Specify OS or Architecture in build:
env GOOS=linux go build . # builds for Linux
env GOARCH=arm go build . # builds for ARM architecture

Build multiple files:
go build main.go assets.go # outputs an executable: main

Building a package:
go build . # outputs an executable with name as the name of enclosing folder

```

