## 介绍
采用fyne开发GUI程序实现PPPoE账号密码恢复功能实现


## 用法
直接下载执行[pppoe-hijack-go.exe](https://github.com/codewindy/RouterPPPoERecover20220828/releases/download/v0.1/pppoe-hijack-go.exe) 文件即可，恢复的账号密码会自动弹出


```bash
git clone https://github.com/codewindy/RouterPPPoERecover20220828
cd pppoe-hijack-go
go build cmd/pppoe-hijack-go.go
```
模仿[PPPoE-hijack](https://github.com/Karblue/PPPoE-hijack)的golang版本

原理: [解密古老又通杀的路由器攻击手法：从嗅探PPPoE到隐蔽性后门](http://www.freebuf.com/articles/wireless/163480.html)

### 列网卡

```bash
./pppoe-hijack-go -l
```

### 嗅探

```bash
./pppoe-hijack-go -i 网卡名

./pppoe.exe -i \\Device\\NPF_{554B45C9-3490-48B6-9CE7-AE72625FDA16}

```
### 参考文档
* https://www.topgoer.cn/docs/fyne/fyne-1e2fi8sdb32j3
* https://www.topgoer.com/
* https://go101.org/article/101.html
* https://www.screentogif.com/
* https://github.com/LuckyC4t/pppoe-hijack-go.git
