#!/bin/bash
# /*
# 参数说明
# --codedir 要分析的代码目录
# --gopath GOPATH环境变量目录
# --outputfile 分析结果保存到该文件
# --ignoredir 不需要进行代码分析的目录（可以不用设置）
# */

./go-package-plantuml --codedir /home/hubing/mygo/src/taro \
--gopath ~/mygo \
--outputfile  /home/hubing/tools/uml/txt/uml.txt \
--ignoredir /home/hubing/mygo/src/taro/vendor
java -jar plantuml.jar /home/hubing/tools/uml/txt/uml.txt
