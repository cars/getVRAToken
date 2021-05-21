#getRefresh Token 

Simple program to get a refresh token from vRealize Automation 8


## Usage

There are five params

* -username(required) : Name of the user logging in to vRA
* -password(required): the user's password
* -domain(optional) : the authentication domain
* -server(optional): defaults to vRA Cloud API Endpoint if not specified
* -type(optional default=refresh): Token Type to return [access|refresh]

Simply returns the refresh_token value and nothing else.


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