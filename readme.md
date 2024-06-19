打包:
https://www.baifachuan.com/posts/4862a3b1.html

CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ go build -o processor

CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  go build -o processor


chmod +x processor