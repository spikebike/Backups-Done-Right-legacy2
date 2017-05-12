## Table of Contents

- [Overview](#overview)
- [How BDR Works](#how-bdr-works)
- [How safe is BDr?](#how-safe-is-BDR)

## Overview

BDR (Backups Done Right) is a backup program that allows trading local disk space for remote disk space to allow offsite backups with minimal cost.  So this allows trading of local disk space (disks can be less than $50 per TB) for valuable offsite backups.

The default configuration should work for most users and theres just a few simple commands to get started.  More advanced users can adjust things as needed.

Encryption is done on the client, and no recovery mechanisms are provided for.  Loss of the password results in loss of data.  Please act accordingly.

## How BDR Works

Each client finds any new, changed, or deleted files in any of the configured directories.  The updated files are checksummed, encrypted into blobs, and checksummed again.  Blobs are queued for upload to a blob server.

The blob server tracks what blobs it already has and when a client offers a blob it's replies with "upload and subscribe" or "I have that, I'll subscribe you to that blob".  It's the blob servers job to manage the blobs, and ensure that the desired replication is met, or to inform the admin otherwise.

The blob server can be configured to require manual introduction of peers, or to search for new peers with a DHT.  This allows automatically finding new peers to trade blobs with.

The blob server:
- accepts blob uploads from trusted clients, if blobs hasn't been seen
- notifies trusted clients if a blob has already been uploaded
- coalesses those blobs into chunks that default to 256MB
- applies a reed solomon code to add redundancy, default to 12 data + 4 parity
- finds 16 or more peers to trade blobs with
- sets up unique per peer keys
- actively monitors blobs for desired replication
- actively monitors peers for quality
- actively challenges peers to prove they have the blobs they claim

## How safe is BDR?

Plain text never leaves the client.  AES256 encryption is the default and is a well respected encryption.  However to allow duplication the NONCE is set to the SHA256 checksum.  This allows the blob server to:
* know the size of the encrypted files
* know which clients share the same encrypted files
* likely be able to fingerprint the client by watching the timing and size of updates and comparing it to automatic patches for various operation systems.

Note that while reusing a NONCE is generally bad, the NONCE should only be used with precisely the same plaintext, so the attacker can't use it to start calculating the password.

The blob server receives the encrypted blobs then builds them into larger chunks.  Then applies a Reed Solomon/Erasure code to them.  Then encrypts them again.  This provides a much higher degress of protection since attackers can no longer tell which client has which blob or the original size of the blob.  More paranoid clients might decide to run the BDR client and blob server on the same machine, which works fine.  The main downside is you lose any deduplication across clients.




