A brief introduction to K8S
===========================

The repository contains the code samples for a short (20 minutes) introduction
to Kubernetes talk.

It contains

 a. a nano-service written in Go in the `./src` folder with
    - an `/healthz` endpoint (to introduce Kubernetes readiness/liveness probes)
    - an `/hello` endpoint that leverages a `PERSON_TO_GREET` env. var. (to
      demonstrate passing env vars to Kubernetes pods, as well as live
      `kubectl edit deployment`)
    - an `/countprimesuntil/:target` endpoint that counts the number of prime
      numbers from 1 to `:target` - which can be pretty CPU intensive and
      therefore, a good way to show Kubernetes autoscaling capabilities, as well
      as resource constraints.
    - a basic request logger (to show off with `kubectl logs`)
    - graceful termination (to show no-downtime rolling upgrades)
    - a Dockerfile that wraps all this in an `alpine` container (so that we can
      show `kubectl exec` capabilites).

 b. The Kubernetes +Istio manifests needed to deploy this service in the
    `./deploy` folder
    * a working [Cert-Manager]() and [External-DNS]() setup is
      needed for automatic TLS and DNS setup (not provided here)
    * manifests are commented and can serve as a reference when writing your
      first manifests



1. Introduction
---------------
* An introduction to the state of infrastructure as code
  - once upon a time, shell-wizards were ruling the infrastructure world
  - the more programmable infrastructure got (Virtualization, Public Clouds...)
    the more engineers turn to code to manage it, yet some problems were still
    to be solved
    * provisioning tools (Chef, Puppett...) VS process supervisors (SystemD,
      Foreman, SupervisorD) VS cloud-specific solutions for autoscaling, service
      discovery
    * many platforms (OpenStack, VMWare, AWS, Google, Microsoft...) and as many
      pieces of code to write for each
    * as many target operating systems to handle as well (tons of Linux
      distributions, Windows...)
    * limited programmability, in particular due to a strict separation between
      provisioners and supervisors... still a lot of shell scripts under the
      hood... supposedly idempotents, wrapped in declarative code
    * therefore, limited shareability of infrastructure code

* One leap forward was still possible... and expected (more and more
  applications with huge infrastructures in the cloud: Uber, Spotify, Lyft,
  DataDog, insert your favorite Start-UP name here...)

  At the same time, AWS was leading the public cloud market, with Google and
  Microsoft struggling to get the remaining market shares... This leap forward
  had to be theirs...

* Introducing Kubernetes: one platform one top of any cloud provider... or bare
  metal. One platform to rule them all ?
  * takes care of both provisoning and supervision (self-healing capabilities),
    auto-scaling
  * one single platform that abstract vendor specific details (to a large
    extent), while making
  * a fully programmable REST-API
  * makes sharing infrastructure code possible (though not always a good
    idea...) with Helm (now run by Microsoft) and operators

2. A hands-on overview of Kubernetes
 - The API we want to deploy: quick walkthrough
 - Deploying containers with Kubernetes... Deployments ! (+ reading logs,
   jumping into containers, port-forwarding...)
 - Exposing them internally with services
 - Exposing them externally with...
 - Zero Downtime rolling updates (+ kubectl edit deployment)

3. Uncovered topics
 - Stateful applications and volumes (and operators <3)
 - Security (or lack, thereof...)
 - Observability on Kubernetes
 - Kubernetes clusters architecture
 - Application Packaging (Helm)
 - GitOps (& tools like ArgoCD)
