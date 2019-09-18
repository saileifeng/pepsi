#!/bin/bash

dir=`pwd`

clear() {
	for d in $(ls ./$1); do
		echo "rm $1/$d filename:$d"
		# pushd命令常用于将目录加入到栈中，加入记录到目录栈顶部，并切换到该目录；若pushd命令不加任何参数，则会将位于记录栈最上面的2个目录对换位置
		pushd $dir/$1/$d >/dev/null
		rm $dir/$1/$d/$d
		# opd用于删除目录栈中的记录；如果popd命令不加任何参数，则会先删除目录栈最上面的记录，然后切换到删除过后的目录栈中的最上面的目录
		popd >/dev/null
	done
}

clear api
clear srv
clear gateway



