package utils

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

func ShowInterfaces() []string {
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range interfaces {
		fmt.Printf("%+v\n", item)
	}
	nameArr := make([]string, 0, len(interfaces))
	println(interfaces)
	for _, data := range interfaces {
		nameArr = append(nameArr, data.Name)
		//nameArr = append(nameArr, "网卡名: "+data.Name+" | 网卡描述: "+data.Description)
	}
	//println(nameArr)
	return nameArr
}
