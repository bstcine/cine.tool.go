package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
	"os"
	"strconv"
	"bufio"
	"log"
)


/// 将100以内的int数据转换为string（显示两位）
func ChangeInt(value int) string {

	s := strconv.Itoa(value)

	if value < 10 {

		s = "0"+s

	}

	return  s
}


/// 读取标准用户输入流（包含提示用户的内容）
/**
 *
 */
func ClientInputWithMessage(message string,endbyte byte) string {

	// 提示用户需要输入信息
	fmt.Println(message)

	value,err := ClientInput(endbyte)

	if err != nil {

		// 告知用户标准输入流异常
		fmt.Println("标准输入流异常，输入信息获取失败")

		return  ""
	}

	return  value
}

/// 读取标准用户输入流(提示用户输入信息)
/**
 * @param endbyte 字符串结束符 char类型 如 '\n', '\t',' '等
 * @return string 除掉结束字符的输入字符串
 * @return error 标准输入流报错
 */
func ClientInput(endbyte byte) (string, error) {

	clientReader := bufio.NewReader(os.Stdin)

	input,err := clientReader.ReadString(endbyte)

	endData := []byte{endbyte,}
	input = strings.Replace(input,string(endData[:]),"",-1)

	return input,err
}

/**
获取日志打印器
 */
func GetLogger(outPath string) *log.Logger {
	dirPath := outPath[0:strings.LastIndex(outPath, "/")+1]
	if _, err := os.Stat(dirPath); err != nil {
		os.MkdirAll(dirPath, 0777)
	}

	logFile, err := os.Create(outPath)

	if err != nil {
		log.Fatalln("open file error !")
	}

	logger := log.New(logFile, "[Info]", log.Lshortfile)
	return logger
}

/**
获取配置文件参数
 */
func GetConfArgs(path string) (argsMap map[string]string) {
	if _, err := os.Stat(path); err != nil {
		argsMap = nil
	} else {
		args, _ := ReadLines(path)
		argsMap = make(map[string]string)

		for i := 0; i < len(args); i++ {
			arg := args[i]
			if !strings.Contains(arg, "#") {
				argSplit := strings.Split(arg, "=")
				if len(argSplit) > 1 {
					argsMap[argSplit[0]] = argSplit[1]
				} else {
					argsMap[argSplit[0]] = ""
				}
			}
		}
	}
	return argsMap
}

/**
获取图片和音频
*/
func GetImageAudio(path string) (images, audios []string) {
	//读取当前目录文件列表
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil
	}

	for i := 0; i < len(files); i++ {
		var info = files[i]
		var name = info.Name()

		if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "resize-") || strings.HasPrefix(name, "logo-") ||
			strings.HasPrefix(name, "m-") || strings.HasPrefix(name, "n-") || strings.HasPrefix(name, "f-") {
			continue
		}

		//图片处理
		if strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".png") {
			images = append(images, name)
		}

		//音频处理
		if strings.HasSuffix(name, ".mp3") {
			audios = append(audios, name)
		}
	}

	return images, audios
}

/**
运行命令
*/
func CineCMD(command string) bool {
	var result bool = true

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			result = false
			fmt.Print(err)
		}
		fmt.Print(string(out))
	} else {
		cmd := exec.Command("bash", "-c", command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			result = false
			fmt.Print(err)
		}
		fmt.Print(string(out))
	}
	return result
}

/**
运行命令
*/
func RunCMD(command string) (string, error) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", command)
		out, err := cmd.CombinedOutput()
		return string(out), err
	} else {
		cmd := exec.Command("bash", "-c", command)
		out, err := cmd.CombinedOutput()
		return string(out), err
	}
}
