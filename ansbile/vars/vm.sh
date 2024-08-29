#!/bin/bash
sudo pacman -S qemu libvirt virt-manager dnsmasq
sudo systemctl start libvirtd
sudo systemctl enable libvirtd
sudo usermod -aG libvirt $USER
qemu-img create -f qcow2 ubuntu-vm.qcow2 20G

sudo virt-install --name=ubuntu-vm \
--ram=3072 \
--disk path=/var/lib/libvirt/images/ubuntu-vm.qcow2,size=20,device=disk,bus=virtio \
--vcpus=2 \
--os-variant=ubuntu24.04 \
--location '/home/'$USER'/ubuntu-24.04-live-server-amd64.iso',kernel=casper/vmlinuz,initrd=casper/initrd \
--extra-args console=ttyS0 \
--graphics none -w bridge=br0,model=virtio
