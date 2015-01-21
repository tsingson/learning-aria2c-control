# 1847test

some golang tools for testing
- aria2jsonrpc client
- nsq webui

## 1.aria2jsonrpc control client

ria2 jsonrpc control client in golang
fork from https://github.com/matzoe/argo and modify for my propose

### aria2 install and running parameter
    use aria2 1.18+ in linux and mac OS
    testting pass in mac OS 10.10

### aria2 install in centOS 6.x 64

##### 1 install libnettle.so first for aria2 dependcy

 downloadn from http://pkgs.repoforge.org/nettle/

CentOS 6.x 64bit install libnettle

wget -c http://pkgs.repoforge.org/nettle/nettle-2.2-1.el6.rf.x86_64.rpm

wget -c http://pkgs.repoforge.org/nettle/nettle-devel-2.2-1.el6.rf.x86_64.rpm

rpm -ivh nettle-2.2-1.el6.rf.x86_64.rpm

rpm -ivh nettle-devel-2.2-1.el6.rf.x86_64.rpm

##### 2 CentOS 6.x 64bit install aria2

wget -c http://pkgs.repoforge.org/aria2/aria2-1.16.4-1.el6.rf.x86_64.rpm

rpm -ivh aria2-1.16.4-1.el6.rf.x86_64.rpm


### run aria2c as:

aria2c  -daemon --enable-rpc=true --check-certificate=false --disable-ipv6 --on-download-complete=~/git/golang/aria2jsonrpc/bin/aria2notify --quiet=true --disk-cache=64M --rpc-listen-all --dir=/temp/ --rpc-allow-origin-all=true --auto-file-renaming=true

### dependency

 
