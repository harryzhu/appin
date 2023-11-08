package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	//"strings"
)

func AdbExec(pkg string, action string) (err error) {
	if AndroidVersion < 9 || AndroidVersion > 14 || pkg == "" || action == "" {
		err = errors.New("--android-version= and --package= cannot be empty")
		fmt.Println(err.Error())
		return err
	}

	action = strings.ToLower(action)

	var cmdExec *exec.Cmd
	var pincode string

	if pincode == "pin" {
		pincode = "1"
	}
	if pincode == "unpin" {
		pincode = "0"
	}

	fmt.Printf("Android:%d\nPackage:%s\nAction:%s\n", AndroidVersion, pkg, action)

	switch AndroidVersion {
	case 13:
		cmdExec = exec.Command(ADBPATH, "shell", "service", "call", "package", "133", "s16", pkg, "i32", "\""+pincode+"\"", "i32", "0")
	case 12:
		cmdExec = exec.Command(ADBPATH, "shell", "service", "call", "package", "136", "s16", pkg, "i32", "\""+pincode+"\"", "i32", "0")
	case 11:
		cmdExec = exec.Command(ADBPATH, "shell", "service", "call", "package", "136", "s16", pkg, "i32", "\""+pincode+"\"", "i32", "0")
	case 10:
		cmdExec = exec.Command(ADBPATH, "shell", "service", "call", "package", "156", "s16", pkg, "i32", "\""+pincode+"\"", "i32", "0")
	case 9:
		cmdExec = exec.Command(ADBPATH, "shell", "service", "call", "package", "151", "s16", pkg, "i32", "\""+pincode+"\"", "i32", "0")

	default:
		log.Fatal("--android-version= should be in 9/10/11/12/13")
	}

	out, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Printf("%s:\n%s\n", pkg, string(out))
		fmt.Printf("ERROR: %s:\n%s\n", pkg, err.Error())
		return err
	}
	fmt.Printf("%s:\n%s\n", pkg, string(out))

	return nil
}

func AdbDetectAndroidVersion() (ver int) {
	cmdExec := exec.Command(ADBPATH, "shell", "getprop", "ro.build.version.release")
	out, err := cmdExec.CombinedOutput()
	s_out := strings.Trim(string(out), "\r\n")

	if err != nil {
		fmt.Printf("GetAndroidVersion:\n%s\n", s_out)
		fmt.Printf("ERROR: GetAndroidVersion:\n%s\n", err.Error())
		return -1
	}

	var i_out int = -1
	if s_out != "" {
		i_out, err = strconv.Atoi(s_out)
		if err != nil {
			//fmt.Printf("AdbDetectAndroidVersion: %s\n", s_out)
			fmt.Println(err)
			return -1
		} else {
			return i_out
		}
	}
	return i_out
}

func AdbList() (err error) {
	cmdExec := exec.Command(ADBPATH, "shell", "pm", "list", "packages", "-3")
	out, err := cmdExec.CombinedOutput()

	if err != nil {
		fmt.Printf("AdbList:\n%s\n", string(out))
		fmt.Printf("ERROR: AdbList:\n%s\n", err.Error())
		return err
	}

	err = ioutil.WriteFile("packages_3.txt", out, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
