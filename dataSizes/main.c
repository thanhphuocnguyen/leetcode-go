#include <stdio.h>

int main()
{
    printf("Size of bool: %zu bytes\n", sizeof(_Bool));
    printf("Size of int: %zu bytes\n", sizeof(int));
    printf("Size of short: %zu bytes\n", sizeof(short));
    printf("Size of char: %zu bytes\n", sizeof(char));
    printf("Size of long: %zu bytes\n", sizeof(long));
    printf("Size of long long: %zu bytes\n", sizeof(long long));
    printf("Size of float: %zu bytes\n", sizeof(float));
    printf("Size of double: %zu bytes\n", sizeof(double));
    printf("Size of long double: %zu bytes\n", sizeof(long double));
    int d = 4;
    _Bool b = 0;

    return 0;
}