#include<iostream>

using namespace std;


int main(void) {

    unsigned long long x = 10;
    cout << "Starting loop wait test" << endl;
    for (unsigned long long i = 0; i < 1000000000000000000000000; i++) {
        if(i % x == 0) {
            cout << i << endl;
            x*=10;
        }
    }

    cout << "done loop wait test" << endl;
}

