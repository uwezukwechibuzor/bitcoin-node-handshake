# BTCHANDSHAKE
**BTCHANDSHAKE** is a network handshake implementation using a publicly available [bitcoin node - btcd](https://github.com/btcsuite/btcd).

## Prerequisites
Before proceeding with the installation, ensure that you have the following prerequisites installed:

- [Golang](https://go.dev/dl/): you can download it from the linked page or:
  - Linux: Use your distribution's package manager.
  - Mac: Use `macports` or `brew`.
- Ensure that `$GOPATH` and `$PATH` have been set properly. On a Mac that uses the Z shell, you may have to run the following:

```zsh
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.zprofile
echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.zprofile
echo "export GO111MODULE=on" >> ~/.zprofile
source ~/.zprofile
```

Clone the Repository
```
git clone https://github.com/uwezukwechibuzor/bitcoin-node-handshake.git
```

### [make](https://www.gnu.org/software/make/)

We use GNU Make to help us built, lint, fmt, and etc for this project.

- Linux:
  - Your distro likely already comes with this. You can check by running `which make`.
  - Otherwise please see your distro specific installation tool(i.e `apt`) and use that to install it.
- Macos:
  - You can check if it's already installed by `which make`.
  - Otherwise use [brew](https://brew.sh/) or [macports](https://www.macports.org/) to install it.

## Building using Make

To build, run:

```bash
make build
```

To install (builds and moves the executable to `$GOPATH/bin`, which should be in `$PATH`), run:

```bash
make install
```

## Linting & Formatting

Run format and run linter for go sided:

```bash
make lint
```

## Unit Testing

To run all unit tests:

```bash
make test
```

To see test coverage:

```bash
make test-cover
```

## TODO - Docker

### [docker](https://www.docker.com/)

Docker is used to help make release and static builds locally.

# btcd & btcwallet Installation Guide

Installation Steps

1. Clone btcd and btcwallet repositories
   
First, clone the btcd and btcwallet repositories from GitHub:


```
git clone https://github.com/btcsuite/btcd.git
git clone https://github.com/btcsuite/btcwallet.git
```

2. Build btcd
Navigate to the btcd directory and build the btcd executable:

```
cd btcd
go install . ./cmd/...
```

3. Build btcwallet
Navigate to the btcwallet directory and build the btcwallet executable:

```
cd ../btcwallet
go install . ./cmd/...
```

Setting up wallets

We’ll need two wallets: one for the miner and one for the user

Before running btcwallet we must create a default wallet
```
btcwallet -C ./btcwallet.conf --create
```
Start btcd server

```
btcd --configfile ./btcd.conf
```
Start btcwallet server on another terminal

```
 btcwallet -C ./btcwallet.conf
```

Start btchandshake Node on another terminal

![Screenshot 2024-04-11 at 09 02 29](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/227d3122-b7e4-475b-a6d8-85811d08ec7f)

```
btchandshake
```

Generate two address for the miner and the user
```
btcctl -C ./btcctl-wallet.conf getnewaddress
```

Next step is to setup miner’s address. Stop btcd and start:

```
 btcd --configfile ./btcd.conf --miningaddr=MINER_ADDRESS
```

Send a Transaction using wallet and check logs on btcd and btchandshake node terminal
```
btcctl -C ./btcctl.conf generate 100
```

## screenshot sample

![Screenshot 2024-04-11 at 08 53 02](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/284f1e4f-06a2-405c-a691-522fbdfc572c)

<img width="755" alt="Screenshot 2024-04-11 at 05 01 39" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/aaeaf8c9-f120-4935-8a84-ed5c97fbfcf6">

<img width="610" alt="Screenshot 2024-04-11 at 05 00 57" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/6f7f477a-52b8-48fa-928c-b6592bc0ce09">

<img width="731" alt="Screenshot 2024-04-11 at 05 00 30" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/7f40c27c-fe7d-44df-a667-f128029666cf">





# Local Bitcoin Network

Bitcoin is a freely available computer program that operates on an open port on your computer, enabling anyone to connect to it and communicate over the Internet.
During development, we'll opt for a local network instead of the main Bitcoin network. For this purpose, I used [btcd](https://github.com/btcsuite/btcd), an alternative Bitcoin full node implementation written in Golang.

![bitcoin1](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/fdb282c9-7b09-4842-bbd9-f5628bcdf10b)

When you run Bitcoin, it utilizes a specific port, 8333 by default, to establish connections with other computers running the same program. This creates a network of interconnected computers communicating with each other


![bitcoin2](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/1bb70916-c061-4724-8292-8b6189334604)
![bitcoin3](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/124a8579-29b5-44a6-9143-f159ca5abc92)



## [BITCOIN PROTOCOL DOCUMENTATION](https://en.bitcoin.it/wiki/Protocol_documentation)

Messages

A "message" is just a structured piece of data that Bitcoin nodes send each other over the network. They all have the same format:

![bitcoin4](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/b32958af-5540-417d-aed2-690a96b20dbb)


Version

The “version” message provides information about the transmitting node to the receiving node at the beginning of a connection. Until both peers have exchanged “version” messages, no other messages will be accepted.

Header 
The header contains a summary of the message, and its structure is the same for every message in the Bitcoin protocol.

<img width="1012" alt="Screenshot 2024-04-11 at 05 35 22" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/1d5527e5-8a17-4ad6-b18d-c710c3409336">

Payload 

The payload contains the main content of the message. Different message types have different structures for their payloads.
When a node creates an outgoing connection, it will immediately advertise its version. The remote node will respond with its version. No further communication is possible until both peers have exchanged their version.

<img width="1218" alt="Screenshot 2024-04-11 at 05 41 31" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/5af94433-407f-4e72-a80e-b5b373e77901">

Handshaking is the initial process that establishes communication between two networking devices. Before data exchange can begin, a "handshake" is necessary. This handshake involves exchanging a sequence of messages to initiate communication.

In the Bitcoin protocol, the handshake follows this sequence:

![bitcoin6](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/46268645-523d-49e4-9796-cacf864ec05e)

The handshake process entails two steps:

We kick off the communication by sending our "version" message, to which they respond with their own "version" message.
Subsequently, they acknowledge receipt of our version message by sending a "verack" message, and we conclude the handshake by reciprocating with a "verack" message back to them.


Message preparation 

We need to send two messages to perform the handshake:

Version Message

Verack Message - The verack message is sent in reply to version. This message consists of only a message header with the command string "verack".
Once we've prepared our messages, sending them to the connected node and receiving messages from them becomes straightforward.

To "send" messages, we write bytes to our TCP socket connection.
To "receive" messages, we read bytes from the same socket.

Once the "verack" message exchange is complete, indicating the conclusion of the handshake, the node will begin sending additional message types. To continue receiving these messages, we simply need to maintain a loop that reads from the socket.

Here's what the incoming messages will resemble:

![bitcoin7](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/82270365-0604-4855-9811-3cd88b72ed52)

When interacting with a node, it won't indiscriminately send all newly received transactions and blocks. Instead, to conserve bandwidth, it will transmit a list of hashes of the latest transactions and blocks via "inv" (inventory) messages.

In response to these "inv" messages, you can specify the specific transactions and blocks you desire using "getdata" messages.
Then, after you've sent your "getdata" message, the node will send you the full transactions and blocks you've requested in subsequent "tx" and "block" messages:

![bitcoin8](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/059df7c6-2075-4b30-86f3-17f0fe24a2af)

<img width="1218" alt="Screenshot 2024-04-11 at 08 03 06" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/b736bf4e-7896-46f2-8d45-6d8755a9d963">

<img width="1218" alt="Screenshot 2024-04-11 at 08 05 15" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/89356553-e2ea-4fcb-b3cd-8a2dd06cf525">

<img width="1218" alt="Screenshot 2024-04-11 at 08 04 30" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/4b067582-fe77-407d-973e-612dfc4576b4">


### Stay Connected

To maintain the connection, the node you've established a connection with will intermittently send "ping" messages to verify your presence. To ensure the connection remains active, you must promptly respond with "pong" messages.


![bitcoin9](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/9b43c159-99cc-4443-b07f-677dee74fe36)


<img width="1218" alt="Screenshot 2024-04-11 at 08 10 39" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/0adcf1ad-0b73-4cb2-bad4-ddb445a5722c">




