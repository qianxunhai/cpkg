package template

var Sh_check_monitor =`#!/bin/bash
Exec_Name=$1
Env=$2
count=`+ "`" + `ps -ef | grep ${Exec_Name} | grep monitor_exec |  grep -v grep | wc -l` + "`"+
`echo $count
if [ $count -eq 0 ];then
   /bin/bash /data/go/src/{{.Name}}/shell/monitor_exec.sh ${Exec_Name} ${Env}
fi
`
