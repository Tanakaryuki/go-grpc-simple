for i in {1..5}; do
  echo "Run #$i"
  k6 run k6/script.js
done
