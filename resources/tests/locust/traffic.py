import time
from locust import HttpUser, task, between
import json
from random import seed
from random import randint

class QuickstartUser(HttpUser):
    wait_time=between(1,4)

    @task
    def on_start(self):
        reg=""
        with open("traffic.json") as file:
            data=json.load(file)
            value=randint(0,19)
            reg=data[value]
            print(reg)

        self.client.get("http://localhost/")
        #self.client.post("http://localhost/", json=reg)