 #!/bin/sh

# 使用tcpdump进行抓包
# -i 指定网卡, lo0是本地回环网卡
# -w 指定输出文件
# port 指定端口
sudo tcpdump -i lo0 -w capture.pcap port 8443