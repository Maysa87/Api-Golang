# Api-Golang
RESTful API server with two endpoints; a repository of data about vintage jazz records.
In this project I'll be storing the data in memory.
* Storing data in memory means that the set of albums will be lost each time you stop the server, then recreated when you start it.

### Prerequisites
* Go install for Mac
  ```sh
  curl -O https://dl.google.com/go/go1.20.5.darwin-amd64.tar.gz
  ```

  2. Clone the repo
   ```sh
   git clone https://github.com/Maysa87/api-golang.git
   ```
3. Install Gin package
   ```sh
   go get -u github.com/gin-gonic/gin
   ```
