#!/bin/bash
set -e

echo "Kubernetes マニフェストの適用と Pod 起動時間（μ秒単位）の計測を開始します"

# 開始時刻をナノ秒単位で記録
# start_time=$(date +%s%N)
start_time=$(gdate +%s%N)

echo "Kubernetes マニフェストを適用中..."
kubectl apply -f k8s/

kubectl rollout status deployment/gateway --timeout=300s
kubectl rollout status deployment/hello --timeout=300s

# end_time=$(date +%s%N)
end_time=$(gdate +%s%N)

elapsed_ns=$(( end_time - start_time ))
elapsed_us=$(( elapsed_ns / 1000 ))

echo "全 Deployment の起動完了までにかかった時間: ${elapsed_us} μ秒"
