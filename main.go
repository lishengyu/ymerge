package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"modify/filescan"
	"modify/logtar"
	"modify/upp"
)

func FileExist(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		if !fi.IsDir() {
			return true
		}
	}
	return false
}

var (
	TplFile string
	UseFile string
	OutFile string
	Module  string
	InKey   string
	InValue string
)

func mergeYamlFile(tpl, use, module string) (data []byte, err error) {
	switch module {
	case "upp":
		data, err = upp.MergeUppCfg(tpl, use)
	case "logtar":
		data, err = logtar.MergeLogtarCfg(tpl, use)
	case "fs":
		data, err = filescan.MergeFilescanCfg(tpl, use)
	default:
		fmt.Printf("Not support Module, please check!!!\n")
		return data, err
	}

	if err != nil {
		fmt.Printf("merge '%s' cfg failed: %v\n", err)
	}

	return data, err
}

func modifyYamlFile(key, value, module, filename string) (data []byte, err error) {
	switch module {
	case "upp":
		data, err = upp.ModifyUppCfg(key, value, module, filename)
	case "logtar":
		data, err = logtar.ModifyLogtarCfg(key, value, module, filename)
	case "fs":
		data, err = filescan.ModifyFilescanCfg(key, value, module, filename)
	default:
		fmt.Printf("Not support Module, please check!!!\n")
		return data, err
	}

	if err != nil {
		fmt.Printf("modify '%s' to '%s' failed: %v\n", key, value, err)
	}

	return data, err
}

func main() {
	flag.StringVar(&TplFile, "t", "anvs_config.yaml", "版本最新的配置模板文件")
	flag.StringVar(&UseFile, "u", "config.yaml", "现网使用的配置模板文件")
	flag.StringVar(&OutFile, "o", "new_config.yaml", "合并后的配置模板文件")
	flag.StringVar(&Module, "m", "", "需要更新的项目名称：upp/logtar/fs")
	flag.StringVar(&InKey, "k", "", "需要修改的key")
	flag.StringVar(&InValue, "v", "", "需要修改的value")
	flag.Parse()

	if Module != "upp" && Module != "logtar" && Module != "fs" {
		fmt.Printf("必须指定项目名称：upp/logtar/fs\n")
		flag.Usage()
	}

	if InKey != "" && InValue != "" {
		data, err := modifyYamlFile(InKey, InValue, Module, UseFile)
		if err != nil {
			fmt.Printf("update [%s] file [%s] key [%s] value to [%s] failed!\n", Module, UseFile, InKey, InValue)
			return
		}

		err = ioutil.WriteFile(OutFile, data, 0644)
		if err != nil {
			fmt.Printf("write yaml file failed: %v\n", err)
			return
		}
		//fmt.Printf("update [%s] file [%s] key [%s] value to [%s] succ!\n", Module, UseFile, InKey, InValue)
	} else if FileExist(TplFile) && FileExist(UseFile) {
		data, err := mergeYamlFile(TplFile, UseFile, Module)
		if err != nil {
			fmt.Printf("merge yaml file failed: %v\n", err)
			return
		}
		err = ioutil.WriteFile(OutFile, data, 0644)
		if err != nil {
			fmt.Printf("write yaml file failed: %v\n", err)
			return
		}
		//fmt.Printf("merge [%s] yaml file [%s] to [%s], generate file [%s] succ!\n", Module, TplFile, UseFile, OutFile)
	} else {
		flag.Usage()
		os.Exit(-1)
	}

	return
}
