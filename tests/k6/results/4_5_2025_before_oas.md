 ~/Work/Own/oas-proxy/tests/k6/ [main*] k6 run main.js
```
          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
          \   |  |\  \ |  (‾)  |
 __________ \  |__| \__\ \_____/ .io

     execution: local
        script: main.js
        output: -

     scenarios: (100.00%) 1 scenario, 100000 max VUs, 7m30s max duration (incl. graceful stop):
              * contacts: Up to 1666.67 iterations/s for 7m0s over 3 stages (maxVUs: 50-100000, gracefulStop: 30s)


     ✗ status is 200
      ↳  99% — ✓ 264168 / ✗ 1

     checks.........................: 99.99% ✓ 264168     ✗ 1     
     data_received..................: 772 MB 1.8 MB/s
     data_sent......................: 21 MB  50 kB/s
     dropped_iterations.............: 4330   10.284897/s
     http_req_blocked...............: avg=11.54µs  min=0s       med=1µs      max=7.8ms  p(90)=5µs      p(95)=7µs  
     http_req_connecting............: avg=8.44µs   min=0s       med=0s       max=7.74ms p(90)=0s       p(95)=0s   
     http_req_duration..............: avg=458.76ms min=148.79ms med=179.49ms max=10.15s p(90)=995.82ms p(95)=2.35s
       { expected_response:true }...: avg=458.72ms min=148.79ms med=179.48ms max=7.96s  p(90)=995.8ms  p(95)=2.35s
     http_req_failed................: 0.00%  ✓ 1          ✗ 264168
     http_req_receiving.............: avg=29.24µs  min=5µs      med=18µs     max=4.79ms p(90)=63µs     p(95)=83µs 
     http_req_sending...............: avg=9.13µs   min=2µs      med=5µs      max=9.12ms p(90)=19µs     p(95)=28µs 
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s     p(90)=0s       p(95)=0s   
     http_req_waiting...............: avg=458.72ms min=148.76ms med=179.46ms max=10.15s p(90)=995.75ms p(95)=2.35s
     http_reqs......................: 264169 627.471362/s
     iteration_duration.............: avg=458.85ms min=148.87ms med=179.54ms max=10.15s p(90)=995.97ms p(95)=2.35s
     iterations.....................: 264169 627.471362/s
     vus............................: 409    min=3        max=3856
     vus_max........................: 3927   min=50       max=3927


running (7m01.0s), 000000/003927 VUs, 264169 complete and 0 interrupted iterations
contacts ✓ [======================================] 000000/003927 VUs  7m0s  1666.66 iters/s
```