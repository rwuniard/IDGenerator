# IDGenerator


## Running the code
You can run with this on the command line after the dependencies are setup.<br/>
It will re-compile/run it if there is any code changes.<br/>
CompileDaemon -command="./go_rest_gin"<br/>
<br/>
## Setting up the dependencies
Watches your go files in a directory and invokes go build if file changed.<br/>  
Github location:<br/>
https://github.com/githubnemo/CompileDaemon<br/>
  
From your project folder where the go.mod is located:<br/>  
go get github.com/githubnemo/CompileDaemon<br/>
go install github.com/githubnemo/CompileDaemon<br/>
Dependency will be added to your go.mod<br/>

The executable will be put in your $HOME/go/bin<br/>
### Web Framework - Gin
GitHub location:<br/>
https://github.com/gin-gonic/gin<br/>
<br/>
<br/>
Load it from your project folder where the go.mod is located:<br/>
go get -u github.com/gin-gonic/gin<br/>
<br/>
Dependency will be added to your go.mod.<br/>
<br/>