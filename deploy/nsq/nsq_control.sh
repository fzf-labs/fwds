#!/usr/bin/env bash

NSQ_VERSION=`./nsqd --version`
NSQ_ADDRESS=""
NSQ_BROADCAST_ADDRESS=""
NSQ_CURRENT_DIR=$(pwd)
NSQ_LOG_DIR=$NSQ_CURRENT_DIR/log

function start(){
    #create nsqlookupd
    for p in {1..2};
    do
            $NSQ_CURRENT_DIR/nsqlookupd \
            -broadcast-address="$NSQ_BROADCAST_ADDRESS" \
            -tcp-address="$NSQ_ADDRESS:900$p" \
            -http-address="$NSQ_ADDRESS:901$p" >> "$NSQ_LOG_DIR/nsqlookupd.log" 2>&1 &
    done

    #create nsqd
    for p in {1..4};
    do
        if [ ! -d $NSQ_CURRENT_DIR"/data/"$p ];then
            mkdir -p $NSQ_CURRENT_DIR"/data/"$p
            if [ ! $? -eq 0 ];then
                echo 'create data directory failture.'
                stop
                exit
            fi
        fi
            $NSQ_CURRENT_DIR/nsqd \
            -data-path "data/$p" \
            -broadcast-address="$NSQ_BROADCAST_ADDRESS" \
            -tcp-address="$NSQ_ADDRESS:1001$p" \
            -http-address="$NSQ_ADDRESS:1002$p" \
            -lookupd-tcp-address="$NSQ_ADDRESS:9001" \
            -lookupd-tcp-address="$NSQ_ADDRESS:9002" >> "$NSQ_LOG_DIR/nsqd.log" 2>&1 &
    done

    #nsqadmin
    $NSQ_CURRENT_DIR/nsqadmin \
        -http-address="0.0.0.0:9010" \
        -lookupd-http-address="$NSQ_ADDRESS:9011" \
        -lookupd-http-address="$NSQ_ADDRESS:9012" >> "$NSQ_LOG_DIR/nsqadmin.log" 2>&1 &
}

function restart(){
    stop
    start
}

#stop process
function stop(){
    for process in nsqlookupd nsqd nsqadmin;
    do
        pkill "$process"
    done
    echo 'ok'
}

#help info
function help(){
    echo $NSQ_VERSION
    echo 'nsq_control start|stop|restart|help'
}

#status
function status(){
    local count=0
    for process in nsqlookupd nsqd nsqadmin;
    do
        local result=`ps -ef|grep "$NSQ_CURRENT_DIR/$process"|grep -v grep|awk '$1="";$3="";$4="";$5="";$6="";$7="";{print $0}'`
        if [ -n "$result" ];then
            echo $result
            let 'count=count+1'
        fi
    done
    if [ $count -ne 0 ];then
        echo -e '\033[32m running \033[0m'
    else
        echo 'stoped.'
    fi
}

if [ $# -eq 0 ];then
    echo 'param can not null.'
else
    if [ ! -d $NSQ_LOG_DIR ];then
        mkdir -p $NSQ_LOG_DIR
        if [ ! $? -eq 0 ];then
            exit
        fi
    fi
    case $1 in
        start)
            start
            ;;
        stop)
            stop
            ;;
        restart)
            restart
            ;;
        status)
            status
            ;;
        help)
            help
            ;;
    esac
fi