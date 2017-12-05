package main

import (
	"./utils"
	"bufio"
	"fmt"
	"os"
)

type InputArgs struct {
	OutputPath string /** 输出目录 */
	LocalPath  string /** 输入的目录或文件路径 */
	LogoPath   string /** 水印图片名称*/
}

var inputArgs InputArgs

func main() {
	debug := true
	if debug {
		inputArgs.LocalPath = "/Volumes/Go/test/"
		inputArgs.OutputPath = "/Volumes/Go/test/"
		inputArgs.LogoPath = inputArgs.OutputPath + ".logo.png"
	} else {
		inputArgs.LocalPath = utils.GetCurPath()
		inputArgs.OutputPath = utils.GetOutPath("doing")
		inputArgs.LogoPath = inputArgs.OutputPath + ".logo.png"
	}

	//创建输出目录
	os.MkdirAll(inputArgs.OutputPath, 0777)
	//检查创建水印图片
	utils.CheckLogoFile(inputArgs.LogoPath)

	//获取图片和音频
	images, audios := utils.GetImageAudio(inputArgs.LocalPath)
	hasMagick := utils.CheckHasMagick()

	fmt.Println(">>>>>>>>>>   开始处理图片（压缩加水印）...")
	for i := 0; i < len(images); i++ {
		var name = images[i]
		fmt.Print(name + " 处理中...")
		if debug {
			utils.ResizeImg(inputArgs.LocalPath, inputArgs.OutputPath, name, "n-"+name)
			utils.ResizeImgByMagick(inputArgs.LocalPath, inputArgs.OutputPath, name, "m-"+name)
		} else {
			if hasMagick {
				utils.ResizeImgByMagick(inputArgs.LocalPath, inputArgs.OutputPath, name, name)
				utils.LogoImgByMagick(inputArgs.LogoPath, inputArgs.OutputPath, inputArgs.OutputPath, name+".jpg", name+".jpg")
			} else {
				utils.ResizeImg(inputArgs.LocalPath, inputArgs.OutputPath, name, name)
				utils.LogoImg(inputArgs.LogoPath, inputArgs.OutputPath, inputArgs.OutputPath, name, name)
			}
		}
		fmt.Println(" 完成.")
	}
	fmt.Println("<<<<<<<<<<   图片处理成功...")

	hasffmpeg := utils.CheckHasFFMPEG()
	if hasffmpeg {
		fmt.Println(">>>>>>>>>>   开始处理音频（压缩）...")
		for i := 0; i < len(audios); i++ {
			var name = audios[i]
			fmt.Print(name + " 处理中...")
			if debug {
				utils.ResizeAudio(inputArgs.LocalPath, inputArgs.OutputPath, name, "f-"+name)
			} else {
				utils.ResizeAudio(inputArgs.LocalPath, inputArgs.OutputPath, name, name)
			}

			fmt.Println(" 完成.")
		}
		fmt.Println("<<<<<<<<<<   音频处理成功...")
	}

	fmt.Println("请输入 end ,结束本程序...")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
	}
}