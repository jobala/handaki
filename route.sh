#! /bin/bash

ip link set dev utun1 mtu 1300
ip addr add 10.0.0.1/24 dev utun1
ip set dev tun1 up
