## 介绍
采用fyne开发GUI程序实现PPPoE账号密码恢复功能实现
* 视频教程视频链接 https://www.douyin.com/video/7136831314338073867
![img](https://user-images.githubusercontent.com/15072465/187062440-88f4dfc1-3439-408f-8c50-1b50aa30fb50.png)
## 仿造华为路由器实现的功能
* ![image](https://user-images.githubusercontent.com/15072465/187063603-c1958c8d-0d64-4d9a-a758-69ecee4c991e.png)
* java 方式实现 https://github.com/codewindy/Mikrotik-Phicomm-Backup

## 用法
直接下载执行[pppoe-hijack-go.exe](https://github.com/codewindy/RouterPPPoERecover20220828/releases/download/v0.1/pppoe-hijack-go.exe) 文件即可，恢复的账号密码会自动弹出
![img](https://user-images.githubusercontent.com/15072465/187062556-a8207ceb-3c8f-432c-9208-66c5a3e2bb10.gif)

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
