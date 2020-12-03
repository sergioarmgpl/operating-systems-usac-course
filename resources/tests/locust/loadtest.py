import time
from locust import HttpUser, task

class QuickstartUser(HttpUser):
    @task
    def access_model(self):
        self.client.post("/URL1"\
	,json={"data":1})

    def on_start(self):
        self.client.get("/URL2")
