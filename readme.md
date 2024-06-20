打包:
https://www.baifachuan.com/posts/4862a3b1.html

CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ go build -o processor

CGO_ENABLED=1 GOOS=linux  GOARCH=amd64  CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++  go build -ldflags '-extldflags "-static"' -o processor


chmod +x processor


### 整理需求
1. 读取配置
2. 扫描出指定目录下所有markdown文件所在目录, 按照目录进行迭代处理
3. 替换图片源, 并将图片采用md5方式替换文件名后保存到指定文件夹
   - 图片存储位置命名规则, 文件夹命名: MD5编码md所在文件夹的相对路径
   - 图片名命名: MD5编码对md文件名+图片原路径
4. 分离markdown头部与主体
   - 需要支持多种配置格式
5. 对头部进行处理
   - 统一输出为yaml格式[暂定]
6. 重新写入文件
7. 根据配置添加search.md和archives.md文件
8. 删除原目录下所有的图片文件 
9. 9删除原目录下所有空文件夹
10. 输出运行日志结果
11. 额外需求: 支持目录级的总体头信息定义, 文件名以.head.md[暂定]命名


