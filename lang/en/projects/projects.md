# Projects
Here you can find ideas of projects to use with your students
## Weather Tweets 2024/Jun [view](project1v6/project1.md)
This is project consists in show realtime system that shows Weather messages using Grafana, for this we are using Kafka Topics, Redis, Mongo and Kubernetes. Also we included Kepler to add a part of environmental sustainability feature to measure the power and carbon emissions for the infrastructure..

## Usactar 2022/Oct-Nov [view](project1v5/project1.md)
This is project consists in show realtime people predictions about soccer matches, for this we are using Kafka Topics, Kubernetes, Linkerd to interconnect a Google and an Azure cluster, and some Google service to show data in realtime. This technologies are used to create our distributed system, and tested using Chaos Mesh.
## USAC Squid Game 2021/Oct-Nov [view](project1v4/project1.md)
This is the First version of the USAC Squid Game, where the students create a distributed systems to generate random data for games, in this case will be based on the Netflix series Squid Game, this project uses Kubernetes, Redis, MongoDB, Linkerd and Chaos Mesh. We are testing, service meshes, traffic splitting, concurrency/parallelism and chaos engineering.


## Realtime COVID-19 Vaccinated People Display 2021/April-May [view](project1v3/project1.md)
This is the third version of the Covid-19 project, its focused on generating a high concurrent distributed system using Kubernetes and containers. This project uses Linkerd to get golden metrics while a traffic simulator send traffic from different sources and different technologies, in the final part of this project its used Linkerd to generate faulty traffic and Chaos Mesh in order to explore Chaos Engineering in a high concurrent systems. The goal is to obtain metrics, implement basic observability and use the features of the services meshes to get into Chaos Engineering. The goal is to control the vaccined people around the world.

## Realtime COVID-19 Display 2020/December [view](project1v2/project1.md)
This is a second version of the Covid-19 project, focused more on golden metrics, using Linkerd, Prometheus and Grafana. All the microservices and apps will be deployend using managed services on the cloud provider and Kubernetes. The traffic in this version is generated using Locust, and all task are processed using queues with Redis or high performance RPC library in this case gRPC.

## Realtime COVID-19 Display 2020/February-November - [view](project1/project1.md)
Build a generic distributed system architecture that shows statistics in realtime using Kubernetes and Cloud Native technologies. This project will be applied to the current infected cases of COVID-19 around the world.
