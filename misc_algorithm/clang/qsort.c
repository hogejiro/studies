#include <stdio.h>

int arrange_array(int a[], int head, int tail) {
    int pivot, left, right, temp;

    pivot = a[head];
    left  = head + 1;
    right = tail;

    while (-1) {
        while (left < tail && pivot > a[left]) left++;
        while (pivot < a[right]) right--;

        if (left >= right) break;

        temp = a[left];
        a[left] = a[right];
        a[right] = temp;

        left++;
        right--;
    }

    // update pivot
    a[head] = a[right];
    a[right] = pivot;

    return right;
}

void quick_sort(int a[], int start, int end) {
    int p;

    if (start < end) {
        p = arrange_array(a, start, end);
        quick_sort(a, start, p - 1);
        quick_sort(a, p + 1, end);
    }
}

// test
int main(void) {
    int i;
    int sample[10] = {4, 15, 3, 2, 500, 6, 0, 15, 80, 1};
    quick_sort(sample, 0, 9);
    for (i = 0; i < 10; i++) {
        printf("%d\n", sample[i]);
    }
}
