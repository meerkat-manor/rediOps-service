# Refresh

The **refresh.sh** bash script refreshes the Docker image
on Linux and restarts the instance.

The **refresh.ps1** PowerShell script refreshes the Docker image
on Windows and restarts the instance. 

In both cases there is no pre-requisite checks made on
dependencies.  The particular dependencies you should have
are:

* Docker installed
* Docker access rights
* Docker composer
* Access to the Docker repository - private or public
* Firewalls open
* No other applicatons bound to the defaul port of 8080

For a more complete solution, including dependency checks, you 
would use something like an Ansible playbook - which can 
be linked in the **devops.json** definition. 
