## Table of Contents

- [Overview](#overview)
- [How BDR Works](#how-bdr-works)
- [How safe is BDr?](#how-safe-is-BDR)

## Overview

BDR (Backups Done Right) is a backup program that allows trading local disk space for remote disk space to allow offsite backups with minimal cost.  So this allows trading of local disk space (disks can be less than $30 per TB and warrantied for 3-5 years) for valuable offsite backups.

The default configuration should work for most users and theres just a few simple commands to get started.  More advanced users can adjust things as needed.

Encryption is done on the client, and no password recovery mechanism is provided.  Loss of the password results in loss of data, please act accordingly.

## How BDR Works

BDR is composed of 3 pieces: the client, the proxy, and the blob server.

The client tracks all changes (additions, deletions and updates) in the configured directories.  Any new/uploaded files are offered to the proxy as encrypted blobs.   The proxy records the client<->blob mapping and allows uploading.  If the encrypted blob checksum matches, the proxy can skip uploading thus enabling deduplications on a per host or across host basis.

Some trust is required of the proxy because it can see the size of encrypted files (blobs) uploaded, this can allow things like seeing which hosts have the same files and OS fingerprinting.   For the most security the proxy can be run on the same host as client, at the cost of reduced deduplication.

The proxy server queues the encrypted blobs into fixed sized chunks then uses Reed Solomon Erasure coding to add redundancy.  Then the fixed size chunks are uploaded to available peers.  The proxy server monitors the entire life cycle of those chunks and actively will garbage collect, add replication, and rebalance chunks as needed.  The proxy will also challange peers to prove they are storing the chunks that they claim to.  Peers that aren't reliable enough will be removed.

The blob server tracks what blobs it already has and when a client offers a blob it's replies with "upload" or "I already have that".  This saves time, disk space, and network bandwidth.  It's the blob servers job to manage the blobs, and ensure that the desired replication is met, or to inform the admin otherwise.

The blob server can be configured to require manual introduction of peers, or to search for new peers with a DHT.  This allows automatically finding new peers to trade blobs with.

The blob server:
- accepts blob uploads from trusted clients, if blobs hasn't been seen
- notifies trusted clients if a blob has already been uploaded
- coalesses those blobs into chunks that default to 256MB
- applies a reed solomon code to add redundancy, default to 12 data + 6 parity
- finds 16 or more peers to trade blobs with
- sets up unique encyption key per peer
- actively monitors blobs for desired replication
- actively monitors peers for quality
- actively challenges peers to prove they have the blobs they claim

## How safe is BDR?

Plain text never leaves the client.  AES256 encryption is the default and is a well respected encryption.  However to allow duplication the NONCE is set to the SHA256 checksum.  This allows the blob server to:
* know the size of the encrypted files
* know which clients have an encrypted blob in common
* likely be able to fingerprint the client OS by watching the timing and size of updates and comparing it to automatic patches for various operation systems.

Note that while reusing a NONCE is generally bad, the NONCE should only be used with precisely the same plaintext, so the attacker can't use repeated NONCEs to attack the encryption key.

The blob server receives the variable sized AES256 encrypted blobs then queues them into 256MB chunks, then a reed solomon erasure code is used to add redundancy.  Peers will then exchange 256MB encrypted blobs.  The blob server will monitor peers and blob redundancy to ensure the desired level of redundancy is maintained.   For those who want the best protection can trade less deduplication for additional security by running the backup client and blob server on the same host.

