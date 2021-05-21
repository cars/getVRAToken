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
