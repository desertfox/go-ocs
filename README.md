# Description
OCS is a wrapper for openshift cli-client *oc* logins to facilitate switching between multiple clusters easily.
---

## Install

- Visit Releases page and download binary for your systems arch

## Commands
`ocs command --opt_name=opt_value`
- login
- list
- clear
- cycle ( No command word executes cycle )
- del
- prune
- help
---

## Examples

### Login 
Adds server host and token hash to local config and executes oc login passthrough
- `ocs login --server=https://somecluster.com:6667 --token=sha256_blahblah`
- `ocs oc login --server=https://somecluster.com:6667 --token=sha256_blahblah`

### List 
Lists Hosts in config, no-op
- `ocs list`

### Clear
Clear config entries
- `ocs clear`

### Cycle
Cycles through Hosts and executes oc login
- `ocs`

### Del
Removes Host from config
- `ocs del 1`

### Prune
Remove credentials that are older than 24 hours
- `ocs prune`

### Help
Prints this file
- `ocs help` or `ocs h` (only command with a "short" flag currently.)
