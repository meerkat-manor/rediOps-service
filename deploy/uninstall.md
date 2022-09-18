# Uninstall

The **uninstall** section of the **devops.json** is to enable the discovery of the functions
to use when a decommisison of the component is required.

The definition does not include details about the server or infrastructure
as these can vary.  This is definition works similar to the Ansible 
inventory and playbooks.

Like all sections, there is guide reference which can give further information
on how the section can be used.

## Engine

The engine identifies the processor for the refresh.

For Ansible, the parameters are defined in the "playbook".

For Command, the commands or script to execute are defined in the "commands"

## Playbook


## Commands

The commands allows for commands to be defined either in the definition or the commands
as found in a file, where such a file can (should) be stored on the source version control
system (e.g. Git)

If the decommission commands are stored in a file in Git, then they can be amended there and tested 
without changing the contents of the devops.json file.

## Example

```json
    "install": {
        "guide": "https://github.com/meerkat-manor/rediOps/blob/main/guide/install.md",
        "engine": "COMMAND",
        "playbook": "",
        "commands": {
            "pre": [],
            "script": "/deploy/decommission.sh",
            "post": []
        }
    }
```