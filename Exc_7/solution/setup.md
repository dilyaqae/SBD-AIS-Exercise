# todo note commands

## Ex 7. Swarm Setup

## 1) Initialize Docker Swarm (manager)

```bash
ipconfig
```

On the manager node we ran:

```bash
docker swarm init --advertise-addr <manager-ip>
```
not simply 

```bash
docker swarm init
```

to avoid docker advertising wrong IP to other computers and specify it explicitly.

This returns a worker join command and returns output like this:

Swarm initialized: current node (9glle87f9lf03dhzf4v8u3qz9) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-53f2qo6nrdnr7n6gdcjdsxm173mzc5tp1nqdigofqo9mu5wl08-bdrls3y785g1is8plygh5lf83 192.168.65.3:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.

## 2) Joining swarm (worker nodes)

On each worker (including me) we ran 

```bash 
docker swarm join --token SWMTKN-1-5eualyrjfi6n49jwk1d8zhdte5pntynak0zthyybxfhijlr8r4-323gnl2jw4yt4wzhz26bzyu9y 159.69.183.68:2377
```

The manager can check all computers currently part of the swarm and confirm if nodes successfully joined.

```bash 
docker node ls
```

