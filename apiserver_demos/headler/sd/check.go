package headler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

//使用 OK 作为 ping 通网络的返回值
func HealthCheck(c *gin.Context)  {
	message := "OK"
	c.String(http.StatusOK,"\n" + message)
}

//检查磁盘的使用情况
func DiskCheck(c *gin.Context)  {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95{
		status = http.StatusTooManyRequests
		text = "危险-磁盘存储空间不足"
	}else if usedPercent >= 90{
		status = http.StatusTooManyRequests
		text = "警告-磁盘使用率过高"
	}

	message := fmt.Sprintf("%s - 剩余空间: %dMB (%dGB) / %dMB (%dGB) | 已使用 %d%%",text,usedMB,usedGB,totalMB,totalGB,usedPercent)
	c. String(status,"\n" + message)
}

// 检查 CPU 的使用情况.
func CPUCheck(c *gin.Context) {

	
}