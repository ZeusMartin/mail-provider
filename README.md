mail-provider
=============

把Gomail封装为一个简单http接口，可以为falcon-plus中的alarm用来发送报警邮件

## 编译步骤

```
# Please make sure that you have set `$GOPATH` and `$GOROOT` correctly.
# If you have not golang in your host, please follow [https://golang.org/doc/install] to install golang.

mkdir -p $GOPATH/src/github.com/open-falcon
cd $GOPATH/src/github.com/open-falcon
git clone https://github.com/ZeusMartin/mail-provider.git
cd $GOPATH/src/github.com/open-falcon/mail-provider
go get ./...
./control build
./control pack
然后你就可以拿着生成的tar.gz包去部署了。

```
## 部署启动步骤

```
export WorkDir="$HOME/open-falcon"
mkdir -p $WorkDir
tar -xzvf falcon-mail-provider-0.0.1.tar.gz -C $WorkDir
cd $WorkDir

# 启动
# 修改cfg.json配置文件中的smtp信息
./control start

# check  status
./control status

```

## 使用方法

```
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"

```
