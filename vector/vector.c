#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#include "vector.h"

static void grow_vector(Vector *v) {
	int newCapacity = v->capacity == 0 ? 1 : v->capacity * 2;
	int *newData = calloc(newCapacity, sizeof(int));

	for (int i = 0; i < v->capacity; i++) {
		newData[i] = v->data[i];
	}

	free(v->data);

	v->capacity = newCapacity;
	v->data = newData;
}

void delete_vector(Vector *v) {
	free(v -> data);

	v -> capacity = 0;
	v -> size = 0;
	v -> data = NULL;
}

Vector new_vector(int capacity) {
	if (capacity < 0) {
		capacity = 0;
	}

	// TODO: What if calloc() fails?
	Vector v = {
		    .data = calloc(capacity, sizeof(int)),
		    .capacity = capacity,
		    .size = 0,
	};

	return v;
}

void print_vector(Vector *v) {
	printf("capacity: %d\n", v -> capacity);
	printf("size: %d\n", v -> size);

	if (v -> data == NULL) {
		printf("data: NULL\n");
		return;
	};

	printf("data: \n");
	for (int i = 0; i < v -> capacity; i++) {
		printf("[%d]: %d\n", i, v -> data[i]);
	}
}
