# Description
OCS is a wrapper for openshift cli-client *oc* logins to facilitate switching between multiple clusters easily.
---

## Install

- CP binary from repo/bin/ocs-$arch that matches your system into your $PATH

## Commands
```
 ocs command --opt_name=opt_value
 ```
- login
    token
    server
- list
- swap
    \#
- clear
- cycle ( No command word executes cycle )
---

## Examples

### Login 
Adds server host and token hash to local config and logs you into cluster
- `ocs login --server=https://somecluster.com:6667 --token=sha256_blahblah`

### List 
Lists Hosts in config, no-op
- `ocs list`

### Swap 
Swap between Hosts in config and logs in
- `ocs swap 2`

### Clear
Clear config entries
- `ocs clear`

### Cycle through clusters
Cycles through Hosts and logs in
- `ocs`