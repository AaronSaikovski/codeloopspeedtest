#include<iostream>

int main()
{
    // size_t n=0;
    // while(n < 1000000000) {
    //     n++;
    // }
    // std::cout << n << std::endl;

    //optimised code - for loop instead of a while loop - ChatGPT!
    uint32_t n=0;
    for (; n < 1000000000; n++) {
        //do nothing
    }
    std::cout << n << std::endl;
}