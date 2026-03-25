ops pkg list


ops volume list

rm -rf ~/.ops/packages/

ops pkg load eyberg/node:v14.2.0 -p 8083 -n -a hi.js

ops image list

ps aux | grep qemu-system-x86_64

ops image create --package eyberg/node:v14.2.0 -a hi.js -c config.json -i my-node-app

ops image list

ops instance create my-node-app -c config.json -p 8083 -t onprem

ops instance list -t onprem

ops instance delete XXXX -t onprem

cloud platform [gcp, aws, onprem, vultr, vsphere, azure, openstack, upcloud, hyper-v, oci, vbox, hetzner, scaleway] (default "onprem")