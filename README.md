# mqtt-golang-vernemq
The vernemq lightweight protocol for mqtt broker in golang


# NOTE

If your running docker you dont need to install golang and vernemq.


# Install go

Install Golang from the link.

https://golang.org/doc/install

add export PATH=$PATH:/usr/local/go/bin to .bashrc   
then do 

source .bashrc  

mkdir go  

then install go dep dependency manager (golang package manager)  

$ go get -u github.com/golang/dep/cmd/dep

then run 

$ dep ensure -v



# Install vernemq

Download Vernemq from here https://vernemq.com/downloads/  

Then Run -->  
$ sudo dpkg -i vernemq_1.4.1-1_amd64.deb  

$ sudo service vernemq start  

Then you can copy the vernemq.conf from path go/src/mhdbs/mqtt-golang-vernemq/config/v/vernemq.conf to /etc/vernemq/
by this command 

Then you can copy the vernemq.conf from path 
$ sudo cp go/src/mhdbs/mqtt-golang-vernemq/config/v/vernemq.conf /etc/vernemq/


# Start the Go Sever

$ go run main.go 

Now the server is started open up the another terminal tab and run the client


# Start the Mqtt go client

$ cd {PWD}/go/src/mhdbs/mqtt-golang-vernemq/client

$ go run clientSub.go


# Running this service by Docker

Install docker from the link 

https://www.docker.com/

$ cd {PWD}/go/src/mhdbs/mqtt-golang-vernemq/

$ docker-compose up

Now the containers are up and running

To show list of containers 

$ docker container ls 

