# BTCHANDSHAKE

Local Bitcoin Network

Bitcoin is a freely available computer program that operates on an open port on your computer, enabling anyone to connect to it and communicate over the Internet.
During development, we'll opt for a local network instead of the main Bitcoin network. For this purpose, I used [btcd](https://github.com/btcsuite/btcd), an alternative Bitcoin full node implementation written in Golang.

![bitcoin1](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/fdb282c9-7b09-4842-bbd9-f5628bcdf10b)

When you run Bitcoin, it utilizes a specific port, 8333 by default, to establish connections with other computers running the same program. This creates a network of interconnected computers communicating with each other


![bitcoin2](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/1bb70916-c061-4724-8292-8b6189334604)
![bitcoin3](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/124a8579-29b5-44a6-9143-f159ca5abc92)

## Messages
A "message" is just a structured piece of data that Bitcoin nodes send each other over the network. They all have the same format:

![bitcoin4](https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/b32958af-5540-417d-aed2-690a96b20dbb)


Version

The “version” message provides information about the transmitting node to the receiving node at the beginning of a connection. Until both peers have exchanged “version” messages, no other messages will be accepted.

Header 
The header contains a summary of the message, and its structure is the same for every message in the Bitcoin protocol.
<img width="1012" alt="Screenshot 2024-04-11 at 05 29 51" src="https://github.com/uwezukwechibuzor/bitcoin-node-handshake/assets/66339097/3c09e965-e77a-41b0-b19f-cda9adf0fbbe">

Magic Bytes: Serve as unique markers at the beginning of a message, aiding in identifying message boundaries within a byte stream.

Command: Specifies the type of message being transmitted, encoded in ASCII characters within a 12-byte field.

Size: Indicates the size of the upcoming payload, specifying the number of bytes needed to read to obtain the complete message.

Checksum: Provides a compact fingerprint for payload integrity verification, computed through double-hashing and extracting the first 4 bytes of the result.

    

