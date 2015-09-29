#!/bin/bash

#oc delete scc allowhostports
oc delete sa kafka
oc delete service kafka kafka-1 kafka-2 kafka-3
oc delete rc kafka-1 kafka-2 kafka-3
oc delete imageStream kafka-0.8.2
