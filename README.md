# Get an access or refresh token from  vRA8/vRAC

Simple program to get either a refresh  or acces token from vRealize Automation 8 that can be used
by something else. 

_This has been tested with vRA8 but not vRA Cloud (though it should work)_




## Usage

There are five params

* -user : Name of the user logging in to vRA(defaults to administrator)
* -password: the user's password
* -domain: (optional) : the authentication domain
* -server: Name/FQDN/IP of vra server (defaults to vRA Cloud Endpoint if not specified)
* -type: Token Type to return  access|refresh (Defaults to refresh if not specified)

Simply returns either the refresh or access token value and nothing else.


## Examples

Get a refresh token from vra.example.com for the user 'administrator' (default) using the password 'MyP@ssw0rd'
```
> getvratoken -server vra.example.com -password 'MyP@ssw0rd'
X6OLxwWUGtf1Xg9mDADyyWXxuUgCQ23u
```

Get a refresh token from vra.example.com for the user 'jdoe' (default) using the password 'MyP@ssw0rd'
```
> getvratoken -server vra.example.com -user jdoe -password 'MyP@ssw0rd'
TVFpSNKO6pQ6lDY9bFNMyl4Eq5MQcZZO
```

Get an access token from vra.example.com for the user 'jdoe' (default) using the password 'MyP@ssw0rd'
```
> getvratoken -server vra.example.com -user jdoe -password 'MyP@ssw0rd' -type access
eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjU2NjYwNTA5OTEwMjMwNTM2MDIifQeyJpc3MiOiJDTj1QcmVsdWRlIElkZW50aXR5IFNlcnZpY2UsT1U9Q01CVSxPPVZNd2FyZSxMPVNvZmlhLFNUPVNvZmlhLEM9QkciLCJpYXQ...
```
Get an refresh token and save it to an environment variable for use with the vRA8 Terraform provider
```powershell
> $ENV:VRA_REFRESH_TOKEN=getvratoken -server vra.example.com -user jdoe -password 'MyP@ssw0rd'
❯ $ENV:VRA_REFRESH_TOKEN
lfr6DQWsd7ALbjEZ6QcPcnze45b9GMAi
```
