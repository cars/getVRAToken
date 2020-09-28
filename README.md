#getRefresh Token 

Simple program to get a refresh token from vRealize Automation 8


## Usage

There are five params

* -username : Name of the user logging in to vRA
* -password: the user's password
* -domain(optional) : the authentication domain
* -server: defaults to vRA Cloud Endpoint if not specified
* -type: Token Type to return [access|refresh]

Simply returns the refresh_token value and nothing else.
