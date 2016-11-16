#include <iostream>
using namespace std;

typedef struct {
    char ch;
    unsigned int freq;
} CharFrequency;

class Heap {
public:
    Heap(size_t max_size) {
        this->_size = 0;
        this->_capacity = max_size;
        this->_heap = new CharFrequency*[max_size + 1];
    }
    ~Heap() {
        if (this->_heap != NULL) {
            size_t max_size = sizeof(this->_heap);
            for (size_t i = 0; i < max_size; i++) {
                if (this->_heap[i] != NULL) {
                    delete this->_heap[i];
                }
            }
            delete this->_heap;
        }
    }
    void insert(CharFrequency* node) {
        if (this->_capacity <= 0) {
            return;
        }

        this->_heap[++this->_size] = node;
        unsigned int index = this->_size;
        while (index > 1 &&
               this->_heap[index]->freq > this->_heap[int(index/2)]->freq) {
            CharFrequency* temp = this->_heap[index];
            this->_heap[index] = this->_heap[int(index/2)];
            this->_heap[int(index/2)] = temp;
            index /= 2;
        }
        this->_capacity--;
    }

    bool is_empty() {
        return this->_size <= 0;
    }

    CharFrequency* extract() {
        if (this->_size <= 0) {
            return NULL;
        }
        CharFrequency* result = new CharFrequency();
        result->ch = this->_heap[1]->ch;
        result->freq = this->_heap[1]->freq;

        this->_swap(&(this->_heap[1]), &(this->_heap[this->_size]));

        this->_size--;
        this->_capacity++;

        this->max_heapify(1);
        return result;
    }

    void max_heapify(unsigned int i) {
        unsigned int left = 2*i;
        unsigned int right = 2*i + 1;
        unsigned int largest = i;

        if (left <= this->_size &&
            this->_heap[left]->freq > this->_heap[largest]->freq) {
            largest = left;
        }
        if (right <= this->_size &&
            this->_heap[right]->freq > this->_heap[largest]->freq) {
            largest = right;
        }
        if (largest != i) {
            this->_swap(&(this->_heap[i]), &(this->_heap[largest]));
            max_heapify(largest);
        }
    }
private:
    CharFrequency** _heap;
    unsigned int _size;
    unsigned int _capacity;
    void _swap(CharFrequency** a, CharFrequency** b) {
        CharFrequency* temp = *b;
        *b = *a;
        *a = temp;
    }
};

class Solution {
public:
    string frequencySort(string s) {
        size_t strlen = s.size();
        for (size_t i = 0; i < strlen; i++) {
            this->_map[(unsigned int)s[i]]++;
        }

        Heap* heap = new Heap(256);
        size_t size = sizeof(this->_map)/sizeof(unsigned int);
        for (size_t i= 0; i < size; i++) {
            if (this->_map[i] <= 0) {
                continue;
            }

            CharFrequency* node = new CharFrequency();
            node->ch = (char)i;
            node->freq = this->_map[i];
            heap->insert(node);
        }

        string result = "";
        while (!heap->is_empty()) {
            CharFrequency* cf = heap->extract();
            unsigned int times = cf->freq;
            while (times-- > 0) {
                result += (char)cf->ch;
            }
        }
        if (heap != NULL) {
            delete heap;
        }
        return result;
    }
private:
    unsigned int _map[256];
};

int main(int argc, char** argv) {
    Solution* solution = new Solution();
    cout << solution->frequencySort(string(argv[1])) << endl;
    if (solution != NULL) {
        delete solution;
    }
    return 0;
}
