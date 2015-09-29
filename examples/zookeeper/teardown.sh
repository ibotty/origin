#!/bin/bash

oc delete scc allowhostports
oc delete sa zookeeper
oc delete service zookeeper zookeeper-1 zookeeper-2 zookeeper-3
oc delete rc zookeeper-1 zookeeper-2 zookeeper-3
oc delete imageStream zookeeper-346-jdk7
