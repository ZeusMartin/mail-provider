mail-provider
=============

把Gomail封装为一个简单http接口，可以为falcon-plus中的alarm用来发送报警邮件

## 使用方法

```
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```
