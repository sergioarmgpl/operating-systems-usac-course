# Projects
Here you can find ideas of projects to use with your students

## Realtime COVID-19 Display 2021/April-May [view](project1v3/project1.md)
This is the third version of the Covid-19 project, its focused on generating a high concurrent distributed system using Kubernetes and containers. This project uses Linkerd to get golden metrics while a traffic simulator send traffic from different sources and different technologies, in the final part of this project its used Linkerd to generate faulty traffic and Chaos Mesh in order to explore Chaos Engineering in a high concurrent systems. The goal is to obtain metrics, implement basic observability and use the features of the services meshes to get into Chaos Engineering. The goal is to control the vaccined people around the world.

## Realtime COVID-19 Display 2020/December [view](project1v2/project1.md)
This is a second version of the Covid-19 project, focused more on golden metrics, using Linkerd, Prometheus and Grafana. All the microservices and apps will be deployend using managed services on the cloud provider and Kubernetes. The traffic in this version is generated using Locust, and all task are processed using queues with Redis or high performance RPC library in this case gRPC.

## Realtime COVID-19 Display 2020/February-November - [view](project1/project1.md)
Build a generic distributed system architecture that shows statistics in realtime using Kubernetes and Cloud Native technologies. This project will be applied to the current infected cases of COVID-19 around the world.