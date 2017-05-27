mail-provider
=============

为Exchange no ssl 25端口提供发送报警邮件功能模块，封装为一个简单http接口。可以和falcon-plus中的alarm配合使用

## 编译步骤

```
# Please make sure that you have set `$GOPATH` and `$GOROOT` correctly.
# If you have not golang in your host, please follow [https://golang.org/doc/install] to install golang.

mkdir -p $GOPATH/src/github.com/open-falcon
cd $GOPATH/src/github.com/open-falcon
git clone git url
cd $GOPATH/src/github.com/open-falcon/mail-provider
go get ./...
./control build
./control pack
然后你就可以拿着生成的tar.gz包去部署了。

```
## 部署启动步骤

```
export WorkDir="$HOME/open-falcon"
mkdir -p $WorkDir/mail-provider
tar -xzvf mail-provider-0.0.1.tar.gz -C $WorkDir/mail-provider/
cd $WorkDir/mail-provider

# 启动
# 修改cfg.json配置文件中的smtp信息
./control start

# check  status
./control status

```

## 使用方法


```
## 配置alarm的cfg.json文件中的 mail 地址
##测试方式
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"

```
