package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func genEnv() (envKey,envValue string){
	letterandnumber := [] rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.Seed(time.Now().Unix())
	key :=[]rune("")
	for i:=0;i<4;i++ {
		number := rand.Intn(len(letterandnumber))
		index := letterandnumber[number : number+1]
		key = append(key, index...)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(letterandnumber), func(i, j int) { letterandnumber[i], letterandnumber[j] = letterandnumber[j], letterandnumber[i] })
	if strings.Contains("1234567890",string(key[0])){
			key[0]='a'
	}
	fmt.Println("export "+string(key)+"="+string(letterandnumber))
	return string(key),string(letterandnumber)

}
// ${LOGNAME:0:1}${LOGNAME:1:1}这种形式进行混淆
func cmdObfuscator(cmd ,envKey,envValue string){
	for i:=0;i<len(cmd);i++{
		pos1 :=strings.Index(envValue,string(cmd[i]))
		if pos1!=-1{
			print("${"+envKey+":"+strconv.Itoa(pos1)+":1}")
		}else{
			print(string(cmd[i]))
		}
	}
}

func main(){
	envKey,envValue := genEnv()
	var cmdStr = flag.String("cmd","","commmand to be obscured")
	var fileStr = flag.String("file","","file cotains commands to be obscured")
	flag.Parse()
	if (*cmdStr != ""){
		cmdObfuscator(*cmdStr,envKey,envValue)
	} else if(*fileStr != ""){
		body, err := ioutil.ReadFile(*fileStr)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		cmdObfuscator(string(body),envKey,envValue)
	}

}
