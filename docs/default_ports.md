While navd is highly configurable when it comes to the network configuration,
the following is intended to be a quick reference for the default ports used so
port forwarding can be configured as required.

navd provides a `--upnp` flag which can be used to automatically map the navcoin
peer-to-peer listening port if your router supports UPnP.  If your router does
not support UPnP, or you don't wish to use it, please note that only the navcoin
peer-to-peer port should be forwarded unless you specifically want to allow RPC
access to your navd from external sources such as in more advanced network
configurations.

|Name|Port|
|----|----|
|Default NavCoin peer-to-peer port|TCP 8333|
|Default RPC port|TCP 8334|
