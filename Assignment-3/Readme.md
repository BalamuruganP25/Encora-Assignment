 

    Prerequisites 
    OS - ubuntu 
1.Golang version go1.15 required

2.Before Run program please follow the steps

    2.1.install golang  (if your already install golang please skip the step)
  
    Note  - User - root
   
    Download  Golang package 
		https://golang.org/doc/install
	Unzip the download folder 
   	   tar -xzf golangpackage name 
	Copy go folder 
		cp -rf  go /user/local
	Set golang path 
 		vim ~/.bashrc
	Add the below line in the bashrc 
		export GOROOT=/usr/local/go
		export GOPATH=$HOME/Goprojects
		export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
	Save the file
		source  ~/.bashrc
	Check go version in comment line
		
		go version 
	 Out put 
		go version go1.15 linux/amd64

 2.2 Add environment variable in bashrc file 
    vim ~/.bashrc

	
	 export API_IP=192.168.43.100
	 export API_PORT=8081


	
	source  ~/.bashrc
2.3 Run the program  
    
    Go to project src folder 
    go run main.go
 
