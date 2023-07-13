package util

/*
func UploadFile(fp *multipart.FileHeader) (uri string, err error) {
	file, err := fp.Open()
	if err != nil {
		log.Printf("[error]util:UploadFile-打开文件错误:%v\n", err.Error())
		return "", err
	}
	defer file.Close()

	// local file 是指针
	localFile, err := os.Create(filepath.Join("./temp/", fp.Filename))
	if err != nil {
		log.Printf("[error]util:UploadFile-创建文件错误:%v\n", err.Error())
		return "", err
	}
	defer localFile.Close()
	// 写入
	_, err = io.Copy(localFile, file)
	if err != nil {
		log.Printf("[error]util:UploadFile-写入文件错误:%v\n", err.Error())
		return "", err
	}
	fileCompressed := CompressPic(localFile)

	return "", nil
}

func CompressPic(fp *os.File) (fileCompressed *os.File) {
	scrImage, _, _ := image.Decode(fp)
	dstImage := image.NewRGBA(image.Rect(0, 0, 80, 80))
	graphics.Thumbnail(dstImage, scrImage)

	return nil
}
*/
