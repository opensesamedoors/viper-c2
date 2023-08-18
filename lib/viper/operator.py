import requests as req 
from viper.errors import *
from viper.util import verify_url 

class Operator:
    def __init__(self, server_url: str, username: str, password: str) -> None:
        if not verify_url(server_url):
            raise req.models.InvalidURL

        self.url      = server_url+"/operator"
        self.username = username
        self.password = password
        self.token    = ""

    def logged(self) -> bool:
        if self.token != "":
            return True 
        return False

    def cookies(self) -> dict:
        if not self.logged():
            raise NotAuth()
        return {"token":self.token}

    def login(self) -> None:
        if self.password == "" or self.username == "":
            raise ServerError("Bad username or password")

        data = {
            "username": self.username,
            "password": self.password
        }

        res = req.post(self.url+"/login", json=data)
        if res.json()["error"] != "":
           raise ServerError(res.json()["error"]) 

        self.token = res.cookies.get("token")

        if self.token == None:
            self.token = ""
            raise CookieNotFound()

    def get_version(self) -> str:
        res = req.get(self.url+"/version", cookies=self.cookies()).json()
        if res["error"] != "":
            raise ServerError(res["error"])

        return res["version"]

    def logout(self) -> None:
        res = req.get(self.url+"/logout", cookies=self.cookies()).json()
        if res["error"] != "":
            raise ServerError(res["error"])

        self.token = ""

    def get_operators(self) -> list:
        res = req.get(self.url+"/operators", cookies=self.cookies()).json()
        if res["error"] != "":
            raise ServerError(res["error"])

        return res["result"]

    def add_job(self, job_action: str, job_data: str, job_target: str) -> None:
        data = {
            "action":   job_action,
            "data":     job_data,
            "target":   job_target
        }

        res = req.post(self.url+"/job", cookies=self.cookies(), json=data).json()
        if res["error"] != "":
            raise ServerError(res["error"])

    def get_result(self, uid: str) -> str:
        res = req.get(self.url+f"/results?uid={uid}", cookies=self.cookies()).json()
        if res["error"] != "":
            raise ServerError(res["error"])
    
        return res["result"]

    def create_token(self) -> str:
        res = req.post(self.url+"/tokens", cookies=self.cookies()).json()
        
        if res["error"] != "":
            raise ServerError(res["error"])

        return res["result"]

    def get_tokens(self) -> list: 
        res = req.get(self.url+"/tokens", cookies=self.cookies()).json()

        if res["error"] != "":
            raise ServerError(res["error"])

        return res["result"]

    def get_agents(self) -> list:
        res = req.get(self.url+"/agents", cookies=self.cookies()).json() 

        if res["error"] != "":
            raise ServerError(res["error"])

        return res["result"]

    def get_jobs(self) -> list:
        res = req.get(self.url+"/jobs", cookies=self.cookies()).json()

        if res["error"] != "":
            raise ServerError(res["error"])

        return res["result"]

    def del_token(self, token: str) -> None:
        res = req.delete(self.url+f"/tokens?token={token}", cookies=self.cookies()).json()

        if res["error"] != "":
            raise ServerError(res["error"])
