Locust is an open source testing tool, after installation we can use it as normal Python module and import the code from other files or packages.


Installation
> pip install locust

Validate the installation
> locust --v

To check what other options we have you can use
> locust --help

Start Locust:
In the path of your "traffic.py" file run:
> locust

If your file is in any other path run:
> locust -f path/traffic.py



This will run your test and enable an IP to access the graphical view of the tool.
If the page doesn't load access using "localhost:port" with the given port and you will see a graphical interface
that require 3 things defined in the traffic.py file:
1. Number of total users to simulate:  the number of elements in your Json file.
2. Spawn rate: the amount of users per second to simulate you can choose a number let's say 10.
3. And the IP/domain of the destination server




https://docs.locust.io/en/stable/quickstart.html