# dev-ops-exercise
A simple exercise for our futures dev-ops to promote their skills

## INTRO

   Your mission, should you choose to accept it, involves the packaging and deployment of 3 μServices.

   Please read the following instructions before starting to implement your mission, you don't want to miss any important instruction, especially those in General Guidelines

   Get your environment ready.

   Just know that we develop on Mac and deploy in production on Google Cloud Platform.

   Ready for action?

   Great.
   Your project is simple, as a DevOps you need to have the ability to package μServices and create a mechanism for deploying them.
   Below, you can find the description of your tasks.

## Packaging

    We have 3 μServices: written in Go, using GRPC and protocol buffers, a Key/Value DB (LMDB).

   We want to containerize them: the strategy is yours. We look forward to hearing your rationale.

## Deployment

   We want to deploy them on Google Cloud Platform and run them under Kubernetes (Google Container Engine).

   You could use gcloud SDK and minikube (local Kubernetes) to do it locally.

   We expect to have least 2 replicas, and expect you to demonstrate us how to scale to 5.

## Automation

   This mission would not be complete without automation.

   We are keen to rerun automatically your packaging and deployment whenever a change occurs.

## Expected Deliverables

   A GitHub Pull-Request to YOUR DUPLICATED REPO, containing:

   What your consider is necessary for any developer to use you automated solution.

##  General Guidelines

   Your implementation should be as simple as possible, yet well documented and robust (easy to use and maintain).

   Spend some time on designing your solution. Think about operational use cases from the real world. Few examples:

   Can you run your implementation multiple times without any problem?
   What happens if a service crashes?
   How would a new Version be deploy to replace the previous one?
   How much effort will it take to create a new service? D.R.Y!

   And tell us how you decided on the solution you are coming up with.

   We sincerely look forward to have you on our team and discuss further your approach to this task.

   Good luck!
