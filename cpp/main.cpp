#include<iostream>

int main()
{
    int x = 0;
    for (int n=0 ; n < 1000000000; n++) {
        x = x + 1;
    }
    std::cout << x << std::endl;
}