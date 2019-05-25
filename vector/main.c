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
	Vector v = new_vector(0);

	print_vector(&v);

	for (int i = 0; i < 20; i++) {
		append_vector(&v, i);
	}

	print_vector(&v);

	test_delete_vector();


}
