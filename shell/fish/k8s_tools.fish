#!/bin/fish

# force delete namespace
function kfdelns
    set -l namespace $argv
    kubectl get namespace $namespace -o json \
        | tr -d "\n" | sed "s/\"finalizers\": \[[^]]\+\]/\"finalizers\": []/" \
        | kubectl replace --raw /api/v1/namespaces/$namespace/finalize -f -
end
