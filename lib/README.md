# Viper | Python API Libary 
To interact with the viper web server, you can use this python API libary,
so you can write your own agents (malware) and operator clients in python 

## Install
- Go to the releases tab 
- Copy the URL to the latest lib.zip file 
- Open a terminal and run:
```bash
pip install <url>
```

## Examples
### Example Agent
```python
from viper import Agent 

SERVER = "http://localhost:8080"
TOKEN  = "here_goes_your_token"
UID    = "123-123-123"

# create an agent object
agent = Agent(SERVER, UID, TOKEN)

# register the agent 
try:
    agent.register("example agent", "hello world!")
except:
    print("Register failed!")
    exit(1)

# get a job
try:
    job = agent.jobs()
except:
    print("No job found!")
    exit(1)

print("New job!")
print(job)

# return a result
agent.result("I did the job!")
```

### Example Operator Client 
```python
from viper import Operator 

SERVER   = "http://localhost:8080"
USERNAME = "ngn"
PASSWORD = "changeme"

# create an operator object
op = Operator(SERVER, USERNAME, PASSWORD)

# try to login
try:
    op.login()
except:
    print("Login failed!")
    exit(1)

# get the server version
print(f"Server version: {op.get_version()}")

# list all connected operators
operators = op.get_operators()
for o in operators:
    print(f"Operator name: {o.split(",")[0]}")
    print(f"Operator last connection: {o.split(",")[1]}")

# list all connected/saved agents
agents = op.get_agents()
if len(agents) == 0:
    print("No agent found!")
    exit(1)

# get a random agent
agent = agents[0]
print(f"UID: {agent['uid']}")
print(f"Desc: {agent['desc']}")
print(f"Token: {agent['token']}")
print(f"Message: {agent['message']}")

# send a random job to the agent 
op.add_job("some action", "some data", agent["uid"])

# list all tokens
tokens = op.get_tokens()
for t in tokens:
    print(f"Token: {t}")

# add a token 
op.create_token()

# logout
op.logout()
```
