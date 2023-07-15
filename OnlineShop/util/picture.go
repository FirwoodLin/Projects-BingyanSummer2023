package util

import (
	"OnlineShop/config"
	"OnlineShop/global"
	"OnlineShop/model"

	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

//type UriBundle struct {
//	OriginPicUri     string
//	CompressedPicUri string
//}

func UploadFile(fp *multipart.FileHeader) (uriBundle model.UriBundle, err error) {
	// Multipart:FileHeader -> File
	file, err := fp.Open()
	if err != nil {
		log.Printf("[error]util:UploadFile-打开文件错误:%v\n", err.Error())
		return
	}
	defer file.Close()
	// os:download file
	fileName := uuid.New().String() + filepath.Ext(fp.Filename) // 生成文件名称 uuid,避免重复
	localFile, err := os.Create(filepath.Join("./temp/", fileName))
	if err != nil {
		log.Printf("[error]util:UploadFile-创建文件错误:%v\n", err.Error())
		return
	}

	defer localFile.Close()
	// 写入 localFile
	_, err = io.Copy(localFile, file)
	if err != nil {
		log.Printf("[error]util:UploadFile-写入文件错误:%v\n", err.Error())
		return
	}
	stat, _ := localFile.Stat()
	log.Printf("[info]util:UploadFile-写入%v文件成功,大小:%v B\n", stat.Name(), stat.Size())
	// 上传原图到 COS
	client := global.GCos
	//_, err = client.Object.Put(context.Background(), stat.Name(), localFile, nil)// 行不通
	_, err = client.Object.PutFromFile(context.Background(), stat.Name(), localFile.Name(), nil)
	if err != nil {
		fmt.Printf("[error]商品图片上传失败：%v\n", err.Error())
		return
	}
	uriBundle.OriginPicUri = config.ProjectConfig.COS.Url + "/" + stat.Name()
	// 压缩 localFile
	fileNameCompressed := uuid.New().String() + filepath.Ext(fp.Filename)
	fileCompressed := CompressPic(localFile, fileNameCompressed)
	stat, _ = fileCompressed.Stat()
	log.Printf("[info]util:UploadFile-压缩文件成功,压缩后名称/大小:%v,%v B\n", stat.Name(), stat.Size())
	// 上传 thumbnail 到 COS
	//client := global.GCos
	//_, err = client.Object.Put(context.Background(), stat.Name(), localFile, nil)
	_, err = client.Object.PutFromFile(context.Background(), stat.Name(), fileCompressed.Name(), nil)
	if err != nil {
		fmt.Printf("[error]商品图片上传失败：%v\n", err.Error())
		return
	}
	uriBundle.CompressedPicUri = config.ProjectConfig.COS.Url + "/" + stat.Name()

	return
}

// CompressPic @param:fp原始文件指针，fn目标文件名称
func CompressPic(fp *os.File, fn string) (fileCompressed *os.File) {
	//stat, _ := fp.Stat()
	//log.Printf("[info]util:CompressPic-传入文件%v,大小:%v B\n", stat.Name(), stat.Size())
	//TODO:疑惑：为什么无法通过文件指针进行 decode 只能通过文件名称进行 Open
	//srcImage, err := imaging.Decode(fp)
	//srcImage, _, err := image.Decode(fp)
	srcImage, err := imaging.Open(fp.Name())
	if err != nil {
		log.Printf("[error]util-CompressPic:解码原始图片错误 %v\n", err.Error())
	}
	dstImage100 := imaging.Resize(srcImage, 100, 0, imaging.Lanczos) // 高度自动按比例缩放
	err = imaging.Save(dstImage100, filepath.Join("./temp/", fn))
	if err != nil {
		log.Printf("[error]util-CompressPic:存储压缩后图片失败failed to save image: %v", err)
	}
	fileCompressed, err = os.Open(filepath.Join("./temp/", fn))
	if err != nil {
		log.Printf("[error]util-CompressPic:打开压缩后图片失败failed to save image: %v", err)
	}
	return
}
