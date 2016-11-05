#include <dispatch/dispatch.h>
#include <mach/mach.h>

int f(int x) {
	int result = (x / 42);

	dispatch_main();

	return result;
}
