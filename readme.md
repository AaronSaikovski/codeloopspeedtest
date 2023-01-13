## Test Results and Setup
Loop over 1,000,000,000 (1 Billion) integers with no third party libraries using only native code constructs.
The results are based in seconds.
The results for compiled languages were run using compiled versions of the code and not interpreted versions and not optimised for multicore/multiprocessor execution - set to single core only.


### Psuedocode
```
int n;
for(n=0;n < 1000000000;n++){
    //do nothing
}
printf("%d ", n);
return 0;
```

### Result Set Run 1 - run on MacBook Pro 2021 (512Gb HDD, 23GB Memory, M1 Pro Apple silicon)

#### Go (v1.19.5)
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 0.32 | 0.32 | 0.32 |


#### Java (v11.0.12 2021-07-20)
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 0.43 | 0.44 | 0.44 |


#### C# (v7.0.102)
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 0.66 | 0.66 | 0.66 |


#### Ansi-C (Apple clang version 14.0.0 (clang-1400.0.29.202))
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 0.96 | 0.96 | 0.96 |


#### C++ (Apple clang version 14.0.0 (clang-1400.0.29.202)
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 0.96 | 0.96 | 0.96 |


#### Rust (v1.66.1)
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 7.99 | 8 | 8 |


#### Python (v3.10.9)
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 21 | 17 | 17 |


#### PowerShell (v7.3.1)
| Run 1 | Run 2 | Run 3 |
|---|---|---|
| 53 | 52 | 51 |