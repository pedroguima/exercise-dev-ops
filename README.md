# dev-ops-exercise
A simple exercise for our future dev-ops to promote their skills.

## Introduction

   Your mission, should you choose to accept it, involves the packaging and deployment of 3 μServices.

   Please read the following instructions before starting to implement your mission, you don't want to miss any important instruction, especially those in General Guidelines

   Get your environment ready.

   Just know that we develop on Osx/Sierra and wish to deploy an efficient `image` in production on Google Cloud Platform.

   Ready for action?

   Great.
   Your project is simple, as a DevOps you need to have the ability to package μServices and create a mechanism for deploying them.
   Below, you can find the description of your tasks.

## Packaging

   We have 3 μServices: written in Go, using GRPC and protocol buffers, a Key/Value DB (LMDB).

   ![Our little system](./dev-ops-exercise.png)

   We want to containerize them: the strategy is yours.
   We look forward to hearing your rationale.

## Deployment

   We want to deploy them on __Google Cloud Platform__ and run them under __Kubernetes__ (Google Container Engine).

   You could use __gcloud SDK__ and __minikube__ (local Kubernetes) to do it locally.

   We expect to have at least 2 replicas (resiliency), and expect you to demonstrate to us how to scale our cluster to 5 replicas.

## Automation

   This mission would not be complete without automation.

   We are keen to rerun automatically your packaging and deployment whenever a change occurs.

## Expected Deliverables

   A GitHub Pull-Request to YOUR FORKED REPO, containing:

   1. What your consider is necessary for any developer to use your automated solution/service.
   2. What you feel should be required (within the context of this exercise)

##  General Guidelines

   Your implementation should be as simple as possible, yet well documented and robust (easy to use and maintain/enhance).

   Spend some time on designing your solution. Think about operational use cases from the real world. Few examples:

       Can you run your implementation multiple times without any problem?
       What happens if a service crashes?
       How would a new Version be deployed to replace the previous one?
       How would you organized a rollback?
       How much effort will it take to create a new service? D.R.Y!
       etc...

   And tell us how you decided on the solution you are coming up with.

   We sincerely look forward to your approach and solution to this task. And if satisfactory, we will discuss it together in the next meeting.

   Good luck!

## If you are curious

To run the backend, the server and the client, open 3 different terminals.

In the first, go to the backend/ folder and from there start the backend, serving the DB:

    1> cd backend
    1> go run backend.go

Second, start the server (our `front-end`) from the server/ folder:

    2> cd server
    2> go run server.go

Last, in the third terminal,  run the client:

    3> cd client
    3> go run client.go

In this third terminal you should receive a list of responses from the `server`.

In the server terminal a response from the backend.

In the backend a list of saved data.

