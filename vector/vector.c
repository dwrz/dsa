#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#include "vector.h"

static void grow_vector(Vector *v) {
	v->capacity = v->capacity == 0 ? 1 : v->capacity * 2;
	v->data = (int *)realloc(v->data, sizeof(int) * v->capacity);
}

void append_vector(Vector *v, int elem) {
	if (v->size == v->capacity) grow_vector(v);

	v->size++;
	v->data[v->size - 1] = elem;
}

void delete_vector(Vector *v) {
	free(v->data);

	v->capacity = 0;
	v->size = 0;
	v->data = NULL;
}

Vector new_vector(int capacity) {
	if (capacity < 0) capacity = 0;

	// TODO: What if calloc() fails?
	Vector v = {
		    .data = (int *)calloc(capacity, sizeof(int)),
		    .capacity = capacity,
		    .size = 0,
	};

	return v;
}

void print_vector(Vector *v) {
	printf("capacity: %d\n", v->capacity);
	printf("size: %d\n", v->size);

	if (v->data == NULL) {
		printf("data: NULL\n");
		return;
	};

	printf("data: \n");
	for (int i = 0; i < v->capacity; i++) {
		printf("[%d]: %d\n", i, v->data[i]);
	}
}
