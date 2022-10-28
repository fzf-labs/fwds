#!/bin/bash
# Go代码部署流程
#1.堡垒机->代码拉取
#2.堡垒机->代码编译
#3.堡垒机->推送(二进制文件)->主机中(多台)
#4.主机代码更新

BASE_DIR=$PWD
#服务名称
Name="cloud_business_pay"
#代码目录
CodePath="/Users/fuzhifei/Code/go/src/kmzx/cloud_business_pay/"
#分支
GitBranch="$1"
#编译时信息注入go代码位置
versionDir=$Name/"pkg/version"
#编译后的代码目录
AppPath=$BASE_DIR/$Name/$GitBranch/"app"/
#备份目录
BackUpPath=$BASE_DIR/$Name/$GitBranch/"backup"/
#部署日志
LogFile=$BASE_DIR/$Name"_deploy.log"
#目标主机代码目录
TargetPath="/www"/$Name/$GitBranch/

#获取git tag 标记
GitTag() {
  if [ "$(git describe --tags --abbrev=0 2>/dev/null)" != "" ]; then
    git describe --tags --abbrev=0
  else
    return
  fi
}
#获取当前最新的commitID
GitCommitID() {
  git rev-parse HEAD
}
#获取当前时间
BuildTime() {
  tz=Asia/Shanghai date +%Y%m%d%H%M%S
  return
}
#Git 代码分支检查
GitBranchCheck() {
  if [ "$GitBranch" == "" ]; then
    Log "请传入参数:代码分支 usage: $0 {branch}" "red"
    exit
  fi
  cd "$CodePath" || exit
  if [ "$(git rev-parse --verify "$GitBranch" 2>/dev/null)" == "" ]; then
    Log "代码分支不存在" "red"
    exit
  fi
}

#代码部署初始化
Init() {
  GitBranchCheck
  [[ ! -d ${AppPath} ]] && mkdir -p "${AppPath}"
  [[ ! -d ${BackUpPath} ]] && mkdir -p "${BackUpPath}"
  [[ ! -f "$LogFile" ]] && touch "$LogFile"
  Log "初始化成功" green
}

#代码备份
BackUp() {
  Log "代码备份开始" green
  cp -rf "$AppPath" "$BackUpPath"/"$(GitCommitID)"
  Log "代码备份成功" green
}

#拉取代码
Pull() {
  Log "代码拉取开始" green
  cd "$CodePath" || exit
  git pull
  git checkout "$GitBranch"
  Log "代码拉取完毕" green
}
#代码构建
Build() {
  cd "$CodePath" || exit
  local ldflags="-X $versionDir.gitTag=$(GitTag) -X $versionDir.buildTime=$(BuildTime) -X $versionDir.gitCommit=$(GitCommitID) -X $versionDir.gitBranch=$GitBranch"
  if (go mod tidy && go mod download && go build -v -ldflags "$ldflags" -o "$AppPath$Name"_api main.go && go build -v -ldflags "$ldflags" -o "$AppPath$Name"_cron cron.go && go build -v -ldflags "$ldflags" -o "$AppPath$Name"_mq_consumer mq_consumer.go); then
    Log "代码编译成功" green
  else
    Log "代码编译失败" red
    exit
  fi
}

Push() {
  Log "代码推送开始" green
  rsync -avzP --delete "$AppPath" root@"$1":"$TargetPath" >>"$LogFile"
  Log "代码推送成功" green
}

Update() {
  Log "主机代码更新开始" green
  local cmd="cd $TargetPath && sh ./admin.sh $2 restart  > /dev/null 2>&1"
  ssh root@"$1" "$cmd" >>"$LogFile"
  Log "主机代码更新成功" green
}

Log() {
  local RED_COLOR='\033[31m'    #红
  local GREEN_COLOR='\033[32m'  #绿
  local YELLOW_COLOR='\033[33m' #黄
  local BLUE_COLOR='\033[34m'   #蓝
  local RES_COLOR='\033[0m'

  case "$2" in
  'red')
    echo -e "$RED_COLOR*** $1 ***$RES_COLOR"
    echo -e "$RED_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  'green')
    echo -e "$GREEN_COLOR*** $1 ***$RES_COLOR"
    echo -e "$GREEN_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  'yellow')
    echo -e "$YELLOW_COLOR*** $1 ***$RES_COLOR"
    echo -e "$YELLOW_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  'blue')
    echo -e "$BLUE_COLOR*** $1 ***$RES_COLOR"
    echo -e "$BLUE_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  *)
    echo -e "*** $1 ***" >>"$LogFile"
    ;;
  esac

}

main() {
  Init
  BackUp
  Pull
  Build
  case $GitBranch in
  'dev')
    Push 122.51.176.45
    Update 122.51.176.45 "$Name"_api
    Update 122.51.176.45 "$Name"_corn
    Update 122.51.176.45 "$Name"_mq_consumer
    ;;
  'master') ;;

  *)
    Log "未指定部署服务器" red
    ;;
  esac

}
main
exit
