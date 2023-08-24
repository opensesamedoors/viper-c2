# Viper | Command & Control Server

## Setup
Download the `server-windows.exe` or the `server-linux` binary
from the releases tab, create a directory for your server, and then you can
just run it

## Configuration
The viper C2 server can by creating `config.json` to file, here are all the
configuration options that you can use:
- ### `address`
Specify the server address, for example `localhost:8087`
- ### `password`
Specify the server password, default is `changeme`
- ### `domain`
Specify the server domain, for exmaple `c2.epicserver.com`
- ### `page404`
Specify a custom 404 page file, default is `404.html`, if the file is not found,
it will just return `<h1>404 Page Not Found</h1>`
- ### `rootpath`
Specify a root path, all the endpoints will be under this path, default is `/`
- ### `database_dir`
Specify the database directory, default is `db`
- ### `header`
Specify a custom server header for the responses, default is `nginx`
- ### `cookie_httponly`
Set to `true` if you want to use HTTP Only cookies, defaukt is `false`
- ### `cookie_secure`
Set to `true` if you want to use secure cookies, default is `true`
- ### `allow_cors`
Set to `false` if you don't want to allow CORS, default is `true`

## The Database
Database is a just directory with JSON files in it, if you want to transfer your
server etc. you need to transfer this directory as well to keep all the data

## How Does the Sever Work?
The C2 server has 2 different types of clients that connects to it.
One of them is the operator and the other one is the agent.

### The Operator
Operator can control the C2 server, it can create/delete tokens, set jobs,
view agents etc.

Here is how an operator looks like in JSON form:
```json
{
    "username": "ngn",
    "lastcon": "01:01:01 01/01/01"
}
```
`lastcon` is the last connection time, each request will update
this time

### The Agent
Agent is the malicious client that can register to the C2 server and can
ask for jobs

Here is how an agent object looks like in JSON form:
```json
{
    "uid": "123-123-123",
    "desc": "Very cool agent 1.0",
    "token": "validtoken321",
    "message": "Hello world!",
    "lastcon": "01:01:01 01/01/01"
}
```
Description is just a small field to store agent information such as 
agent version, name etc. But you can use it for really anything

Message is a one time thing that will be stored when the agent 
registers, message may contain stolen information, system information etc. 

### How Does the Operator Authenticate?
The C2 server has a password that is specified in the configuration file, 
operators send a login request with this password to receive a cookie that will
be valid till the operator logs out

### How Does the Agent Authenticate?
Agents use tokens to authenticate. Tokens are just passwords that can be created/deleted
by the operators, agents need to register using a valid token and all their requests should contains this token

### What is a Job?
- A Job is just a JSON object that tells it's target agent what it needs to do. Here is an example job object:
```json
{
    "action": "cmd",
    "data": "whomai",
    "target": "123-123-123",
    "done": false
}
```
- Jobs can be added by the operators
- A job can target only one agent
- Server provides an avaliable job to the agent when it asks for it
- Agent returns results for a job when it's done
- Jobs are not saved to the database, all the jobs and their results will be lost when the server restarts 

<br>

## Endpoints 
Here are all the endpoints for the C2 web server 

#### Operator Endpoints (`/operator`)
* #### GET `/version`
Returns server version, requires auth

* #### POST `/login`
Example body:
```json
{
    "username": "ngn",
    "password": "changeme"
}
```
Returns a cookie (token) for operator to use in other reqeusts

* #### GET `/logout`
Logs out the operator, requires auth 

* #### GET `/operators`
Returns a list of operators, requires auth

* #### POST `/job`
Example body: 
```json
{
    "action": "cmd",
    "data": "dir",
    "target": "123-123-123"
}
```
Add a new job, requires auth 

* #### GET `/results?uid=<uid>`
Returns the job result for agent with the `<uid>`, requires auth

* #### GET `/jobs`
Returns a list of jobs, requires auth 

* #### GET `/tokens` 
Returns a list of tokens, requires auth

* #### POST `/tokens`
Creates a new token, requires auth 

* #### DELETE `/tokens?token=<token>`
Removes the `<token>`, requires auth 

* #### GET `/agents` 
Returns a list of agents, requires auth 

#### Agent Endpoints (`/agent`)
* #### POST `/register`
Example body:
```json
{
    "uid": "123-123-123",
    "desc": "Very cool agent 1.0",
    "token": "validtoken321",
    "message": "Hello world!",
}
```
Registers a new client if it doesn't already exists


* #### POST `/result?uid=<uid>&token=<token>`
Example body: 
```json
{
    "result": "hey world!"
}
```
Set the job result of the agent with the `<uid>`, `<token>` needs 
to be same with the agent registeration token

* #### GET `/job?uid=<uid>&token=<token>`
Returns an avaliable job for the agent with the `<uid>`, `<token>` needs to be same with the agent registeration token
