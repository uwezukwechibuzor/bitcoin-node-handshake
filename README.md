# BTCHANDSHAKE

Local Bitcoin Network

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











