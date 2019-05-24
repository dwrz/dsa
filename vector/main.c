#include <assert.h>
#include <stdio.h>
#include "vector.h"

void test_delete_vector(void) {
	Vector v = new_vector(2);

	delete_vector(&v);

	assert(v.capacity == 0);
	assert(v.size == 0);
	assert(v.data == NULL);
}

int main(void) {
	Vector v = new_vector(2);

	print_vector(&v);

	test_delete_vector();

}
