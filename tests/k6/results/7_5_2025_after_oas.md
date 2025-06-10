```

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: tests/k6/main.js
        output: -

     scenarios: (100.00%) 1 scenario, 100000 max VUs, 7m30s max duration (incl. graceful stop):
              * contacts: Up to 1666.67 iterations/s for 7m0s over 3 stages (maxVUs: 50-100000, gracefulStop: 30s)


     ✗ status is 200
      ↳  99% — ✓ 263119 / ✗ 199

     checks.........................: 99.92% ✓ 263119     ✗ 199   
     data_received..................: 291 MB 691 kB/s
     data_sent......................: 24 MB  58 kB/s
     dropped_iterations.............: 5181   12.328094/s
     http_req_blocked...............: avg=7.44µs   min=0s       med=1µs      max=13.42ms p(90)=3µs      p(95)=5µs     
     http_req_connecting............: avg=4.88µs   min=0s       med=0s       max=4.5ms   p(90)=0s       p(95)=0s      
     http_req_duration..............: avg=333.92ms min=1.26ms   med=130.57ms max=8.74s   p(90)=223.81ms p(95)=657.63ms
       { expected_response:true }...: avg=334.1ms  min=119.02ms med=130.58ms max=8.74s   p(90)=223.85ms p(95)=657.6ms 
     http_req_failed................: 0.07%  ✓ 199        ✗ 263119
     http_req_receiving.............: avg=18.65µs  min=5µs      med=12µs     max=2.45ms  p(90)=38µs     p(95)=56µs    
     http_req_sending...............: avg=6.66µs   min=1µs      med=4µs      max=14.26ms p(90)=11µs     p(95)=19µs    
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s      p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=333.9ms  min=1.24ms   med=130.54ms max=8.74s   p(90)=223.79ms p(95)=657.59ms
     http_reqs......................: 263318 626.560351/s
     iteration_duration.............: avg=333.99ms min=1.29ms   med=130.61ms max=8.74s   p(90)=223.87ms p(95)=657.92ms
     iterations.....................: 263318 626.560351/s
     vus............................: 317    min=2        max=4923
     vus_max........................: 5024   min=50       max=5024


running (7m00.3s), 000000/005024 VUs, 263318 complete and 0 interrupted iterations

```