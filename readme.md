# Blog Processor
针对hugo的Blog文件预处理器, 使用本项目可以更加便捷的组织您的hugo博客, 简化头部信息的定义, 一键替换图片源.

### How it works
1. 读取config.yaml配置
2. 扫描出指定目录下所有markdown文件所在目录, 按照目录进行迭代处理
3. 替换图片源, 并将图片采用md5方式替换文件名后保存到指定文件夹
   - 图片存储位置命名规则, 文件夹命名: MD5编码md所在文件夹的相对路径
   - 图片名命名: MD5编码对md文件名+图片原路径
4. 分离markdown头部与主体
   - 支持yaml, json, toml头部格式
5. 对头部进行处理
   - 程序首先创建全局头部, 然后扫描每个目录下头部文件生成目录级头部.
   - 单个文件的头部会首先和所在目录的头部进行merge操作, 然后和全局头部进行merge.
6. 重新写入文件, 头部统一输出为toml格式
7. 根据配置添加search.md和archives.md文件
8. 删除原目录下所有的图片文件 
9. 删除原目录下所有空文件夹
10. 支持目录级的总体头信息定义, 文件名默认以.head.toml命名, 可在config.yaml中修改

### config.yaml
```yaml
Basic:
  BlogDir: "../content" # 博客文件所在目录
  ImageDir: "../content/assert" # 废弃配置项
  TemplateFile: "/" # 暂未使用的配置项
  OutputDir: "../static"  # 静态文件输出目录
  HttpBasePath: "https://123/static/" # 静态文件所在url, 用于一键替换图片源
  DateLayout: "2006-01-02"  # 日期格式
  CommonHeaderFileName: ".header" # 目录级头部信息定义文件名
  CommonHeaderFileExt: ".toml"    # 目录级头部信息定义文件后缀
  CommonHeaderFileFormat: "toml"  # 目录级头部信息定义文件内容格式

DataBase:
  DBFile: "../meta.db"  # sqlite 数据库文件所在路径

SearchPageConfig: # search页面配置
   Require: true  # 是否需要自动生成search页面, 以下为search页面头部信息
   Title: "Search"
   Layout: "search"
   Summary: "search"
   Placeholder: "search..."
   Type: "special"

ArchivesPageConfig: # archives页面配置
   Require: true  # 是否需要自动生成archives页面, 以下为archives页面头部信息
   Title: "Archives"
   Layout: "archives"
   Url: "/archives/"
   Summary: "archives"
   Type: "special"

GlobalHeaderConfig: # 全局头部信息
   Author: "RLTEA"
```

### 打包发布

```shell
CGO_ENABLED=1 GOOS=linux  GOARCH=amd64  CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++  go build -ldflags '-extldflags "-static"' -o processor

chmod +x processor
```
