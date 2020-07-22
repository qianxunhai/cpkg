package template

var Sh_publish_script =`#!/bin/bash

HOME=/home/CI
#source ~/.profile
Delopy_User='root'
Deploy_Port='52222'
Deploy_Result='True'
Exec_Name=$1
for i in `+"`"+`echo ${Servers} | awk -F',' '{for(i=1;i<=NF;i++)print $i}'`+"`"+`;do
if [ -n "$i" ];then
Old_IFS=${IFS}
IFS=':'
#取出服务器变量信息172.18.15.23:/app/www/test-project => $1 $2
set -- $i
Delopy_Ip=$1
Deploy_Path=$2
IFS=${Old_IFS}
ssh -p ${Deploy_Port} ${Delopy_User}@${Delopy_Ip}  "/bin/ps -ef |grep "${Exec_Name}"| egrep -v 'grep|monitor'  | /bin/awk '{print \$2}' |xargs /bin/kill -9"

fi
done`
