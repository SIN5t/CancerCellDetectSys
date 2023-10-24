// Code generated by hertz generator.

package upload

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"log"
	"os"

	upload "github.com/SIN5t/CancerCellDetectSys/cmd/api/biz/model/upload"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UploadFile .
// @router ccds/upload [POST]
func UploadFile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req upload.UploadRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	file, err := c.FormFile("file_content")
	fileSavePath := "./temp/" + file.Filename // TODO 写入配置文件
	createDir("./temp/")

	err = c.SaveUploadedFile(file, fileSavePath)
	if err != nil {
		hlog.Error(err.Error())
		c.JSON(consts.StatusInternalServerError, &upload.UploadResponse{
			StatusCode: 1,
			StatusMsg:  "文件上传失败",
		})
		return
	}

	fmt.Println(req.FileName)
	c.JSON(consts.StatusOK, &upload.UploadResponse{
		StatusCode: 0,
		StatusMsg:  "文件成功上传",
	})
}

func createDir(dirPath string) {
	// 判断目录是否存在
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	} else if err != nil {
		log.Fatalf("Failed to check directory: %v", err)
	}

	log.Println("Directory exists or created successfully")
}
