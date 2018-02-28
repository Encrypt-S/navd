rpcclient
=========

[![Build Status](http://img.shields.io/travis/aguycalled/navd.svg)](https://travis-ci.org/aguycalled/navd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/aguycalled/navd/rpcclient)

rpcclient implements a Websocket-enabled Navcoin JSON-RPC client package written
in [Go](http://golang.org/).  It provides a robust and easy to use client for
interfacing with a Navcoin RPC server that uses a navd/navcoin core compatible
Navcoin JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](http://godoc.org/github.com/aguycalled/navd/rpcclient)
* [navd Websockets Example](https://github.com/aguycalled/navd/rpcclient/blob/master/examples/navdwebsockets)  
  Connects to a navd RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [navwallet Websockets Example](https://github.com/aguycalled/navd/rpcclient/blob/master/examples/navwalletwebsockets)  
  Connects to a navwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign
* [Navcoin Core HTTP POST Example](https://github.com/aguycalled/navd/rpcclient/blob/master/examples/navcoincorehttp)  
  Connects to a navcoin core RPC server using HTTP POST mode with TLS disabled
  and gets the current block count

## Major Features

* Supports Websockets (navd/navwallet) and HTTP POST mode (navcoin core)
* Provides callback and registration functions for navd/navwallet notifications
* Supports navd extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get -u github.com/aguycalled/navd/rpcclient
```

## License

Package rpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
