package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/LuckyC4t/pppoe-hijack-go/internal/actions"
	"github.com/LuckyC4t/pppoe-hijack-go/internal/global"
	"github.com/LuckyC4t/pppoe-hijack-go/internal/utils"
	"github.com/flopp/go-findfont"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func init() {
	//设置中文环境
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	os.Setenv("FYNE_THEME", "light")

	a := app.NewWithID("icons")
	w := a.NewWindow("路由器PPPoE密码恢复器 -2022")
	w.Resize(fyne.NewSize(1000, 600))
	// set icon
	r, _ := fyne.LoadResourceFromPath("Main.ico")
	w.SetIcon(r)
	nameArr := utils.ShowInterfaces()

	var selectDevice string
	label := widget.NewLabel("请选择一个网卡 ")
	combo := widget.NewSelect(nameArr, func(device string) {
		selectDevice = device
		log.Println(selectDevice)
	})
	combo.PlaceHolder = "Select Network Interface  here..."
	btn := widget.NewButton("点击开始恢复PPPoE密码", func() {
		// Open device
		global.Device = selectDevice
		log.Println(selectDevice)
		handle, err := pcap.OpenLive(global.Device, 1024, false, -1*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		global.Handle = handle
		defer handle.Close()
		getwd, err := os.Getwd()
		// Use the handle as a packet source to process all packets
		fmt.Println("start sniffing...")
		currentTime := time.Now()
		format := currentTime.Format("2006-01-02 15:04:05")
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			if packetPayload := actions.Hijack(packet); packetPayload != nil {
				fmt.Println("---------------------------------------")
				fmt.Println("报文解析时间: ", format)
				fmt.Println("username:", packetPayload[0])
				fmt.Println("password:", packetPayload[1])

				os.WriteFile(getwd+"/pppoe.txt", []byte("报文解析时间: "+format+"\n"+"=========="+"\n"+"上网宽带账号:"+packetPayload[0]+"\n"+"密码:"+packetPayload[1]), 0666)
				exec.Command("cmd", "/C", "start", "C:\\Windows\\System32\\notepad.exe", getwd+"/pppoe.txt").Run()
				//exec.Command("cmd", "/C", "start", "C:\\Users\\Administrator\\AppData\\Local\\CentBrowser\\Application\\notepad.exe", getwd+"/pppoe.txt").Run()
				fmt.Println("---------------------------------------")
				break
			}
		}
	})
	c := container.NewVBox(label, combo, btn)
	w.SetContent(c)
	w.ShowAndRun()

}
