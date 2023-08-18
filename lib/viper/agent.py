import requests as req
from viper.errors import *
from viper.util import verify_url

class Agent:
    def __init__(self, server_url: str, uid: str, token: str) -> None:
        if not verify_url(server_url): 
            raise req.models.InvalidURL 
    
        self.url    = server_url+"/agent"
        self.uid    = uid
        self.token  = token

    def auth(self, url: str) -> str:
        return url+f"?uid={self.uid}&token={self.token}"

    def register(self, desc: str, message: str) -> None:
        data = {
            "uid":      self.uid,
            "desc":     desc,
            "token":    self.token,
            "message":  message
        }

        res = req.post(self.url+"/register", json=data).json()
        if res["error"] != "":
            raise ServerError(res["error"]) 

    def result(self, result: str, done=True) -> None:
        data = {
            "result":   result,
            "done":     done
        }

        res = req.post(self.auth(self.url+"/result"), json=data).json()
        if res["error"] != "":
            raise ServerError(res["error"])

    def jobs(self) -> dict: 
        res = req.get(self.auth(self.url+"/job")).json()

        if res["error"] != "":
            raise ServerError(res["error"])

        return res["result"]
