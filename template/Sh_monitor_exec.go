package template

var Sh_monitor_exec =`#!/bin/bash
Exec_Name=$1
Env=$2
while [ 1 ]
do
	ct=`+"`"+`ps -ef | grep "\.\/"${Exec_Name} | egrep -v "grep|monitor" | wc -l`+"`"+
`if [ $ct -eq 0 ];then
cd /data/go/src/{{.Name}}/
export GIN_MODE=${Env}
nohup ./${Exec_Name} >> /data/logs/go/${Exec_Name}.log 2>&1 &
fi
sleep 1
done`
