package main

import (
	"fmt"
	"mc-camera/Capture"
	"mc-camera/Upload"
	"path/filepath"
	"os"
)
func GetFileInFolder(folder string) ([]string, error) {
	var files []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() == false {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
const(
	Cam_ID string ="rtsp://172.168.0.3:554/user=admin_password=tlJwpbo6_channel=1_stream=0.sdp?real_stream/tcp/av0_0"
	XML = "haarcascade_frontalface_default.xml"
)
func main() {
	for{
		Capture.Take_video(Cam_ID)
		Capture.DetectImage(XML,"Image/anh_selfile6.jpg")
		filename,err:= GetFileInFolder("Image")
		if err == nil{
			Capture.DNN_detection(filename[0], "res10_300x300_ssd_iter_140000.caffemodel", "deploy.prototxt.txt")
		}else {
			continue;
		}
		for {
			filename,err=GetFileInFolder("Detected")
			if err == nil && filename != nil{
				fmt.Printf("Da lay file %s trong Detected",filename)
				Upload.Upfile(filename[0])
			}else {
				break;
			}
		}
		
		
	}
}


