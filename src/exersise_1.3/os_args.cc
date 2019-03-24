/*
 * same behaviour like os_args.go
 *
 * BSD 3-Clause
 * (c) 2015, thorsten.johannvorderbrueggen@t-online.de
 */

#include <iostream>
#include <chrono>
#include <vector>

using namespace std;
using namespace std::chrono;

int main(int argc, char *argv[])
{
	string res;

	cout << ("\n\nFourth version -> lang==cpp\n");

	auto t0 = high_resolution_clock::now();
	for (int i = 1; i < argc; i++)
		res = res + argv[i] + " ";
	auto t1 = high_resolution_clock::now();

	cout << res << endl;
	cout << "Result in micro seconds: "
	     << duration_cast<microseconds> (t1 - t0).count()
	     << endl;
	// Result -> in micro seconds: 120

	exit(0);
}
