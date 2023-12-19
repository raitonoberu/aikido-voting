import os
import requests


class Service(object):
    API_BASE = "https://web-portfolio.tech/api"

    def __init__(self):
        self.email = os.getenv("ADMIN_EMAIL")
        self.password = os.getenv("ADMIN_PASSWORD")
        self.token: str = None

    def login(self):
        resp = requests.post(
            self.API_BASE + "/user/login",
            json={"email": self.email, "password": self.password},
        ).json()
        if "error" in resp:
            raise Exception("Couldn't login: " + resp["error"])
        self.token = resp["token"]

    def get_pools(self) -> list:
        if self.token == None:
            self.login()

        resp = requests.get(
            self.API_BASE + "/pool",
            headers={"Authorization": "Bearer " + self.token},
        ).json()
        if "error" in resp:
            raise Exception("Couldn't get pools: " + resp["error"])
        return resp
