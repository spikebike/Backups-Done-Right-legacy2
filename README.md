# Backups-Done-Right
Backups-Done-Right (BDR) is a p2p backup system allowing the trading of local storage for offsite backups.  The general idea is that you tell BDR which directories to backup, set a few basic settings, and tell BDR how much space it can use to trade.  From then on it should be self maintaining and will notify the user of problems.


## Table of Contents
- [Client](#Client)
- [Server](#Server)
- [DHT](#DHT)

##Client

The client walks the configured directories tracking any deletions, updates, or additions.  All filesystem related metadata is kept locally and any changes or additions are encrypted and uploaded to the blob server.  Deduplication is implemented, but only among clients sharing the same encryption key.

##Server

The server receives SHA256 checksums from the client and accepts uploads for blobs it hasn't seen before.  Then applies a reed-solomon error correction code to generate chunks, which are traded with other peers.  The server works to ensure the specified replication is maintained and periodically challenges other peers to ensure they are storing the agreed upon chunks.

##DHT
The DHT allows (if desired) the finding of other peers to trade chunks with.
