
## COP5615 - Distributed Operating System Principles
## Project 2

| Member | UFID |
| ------ | ------ |
| jagan mohan reddy Dwarampudi | jdwarampudi@ufl.edu |
| Praveen Ravula | praveenravula@ufl.edu |

## Overview
The main goal of this project is to implement Gossip algorithms using four topologies.

Steps to run:

1.	erl -name <user>@<ip-address>
enter the erlang command line using a username and working machine ip address. 
2.	c(main).
Compile the main code.
3.	c(gossip).
Compile the gossip erl file.
4.	c(pushsum).
Compile the pushsum erl file.
5.	c(topology).
Compile the topology erl file.
6.	main:start(N, "Topology","algorithm").

#### Variables
     Algorithm : 
      > gossip
      > push-sum
     N: Number of Nodes
     Topology :
      > full
      > 2D
      > line
      > Imp2D

## Working Simulation
#### What is working?
The gossip and pushsum algorithms with Line, 2D, Full and Imp2D are fully functional with these topologies and the output shows the time and status took for the converegence.

For all the topologies we scaled up to 4096 (212).

## Simulation and results.
Here the program is executed on all 4 topologies with scaling an exponential in powers of 2 from 8 to 4096 and the graphs are interpreted as results with respective to the overall run time. From these observations we can see that number of workers increases with the time for convergence also increased.

In gossip algorithm full topology is found to be relatively faster than other three topologies whereas in pushsum algorithm Imp2D topology is found to be faster than the other topologies.
  
Gossip algorithm simulations:
  
<img width="613" alt="Gossip_graph" src="https://user-images.githubusercontent.com/46889521/194989962-384573b0-88ed-452b-9bff-ef95732ef4c1.png">


Pushsum algorithm simulations:
<img width="613" alt="pushsum_graph" src="https://user-images.githubusercontent.com/46889521/194989985-b0958fbf-ad31-4cc9-b9f9-7cda55f13412.png">
