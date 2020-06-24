# MygolangProject

Required installation Setup in Local:
==========================================
1. Download and Install go.(I installed go version go1.14.2 windows/amd64).
2. Set your go workspace path in environment variables (varName:GOPATH,varValue:C:\MyGoWorkspace) and this path can be any folder which contains pkg,bin,src folders.
3. I am Using GoLand IDE (by Jetbrains) to code.
4. Install Mysql 5.7

Project Setup:
=============
1. create a folder structure like this in your go workspace path "\MyGoWorkspace\src\github.com" and clone the project.
2. Execute the command "go build" and it downloads the dependencies mentioned in go.mod, compiles and creates a MygolangProject.exe file if compilation is successful.
3. run MygolangProject.exe in gitbash or cmd prompt.

