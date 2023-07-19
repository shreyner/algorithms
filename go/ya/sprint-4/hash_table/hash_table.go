package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BucketValue struct {
	Key   int
	Value *string
}

type HashTable struct {
	size    int
	buckets []*BucketValue
}

func NewHash(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*BucketValue, size),
	}
}

func (t *HashTable) Put(key int, value string) error {
	hash, err := t.findHashEmptyBucket(key)

	if err != nil {
		return err
	}

	t.buckets[hash] = &BucketValue{
		Key:   key,
		Value: &value,
	}

	return nil
}

func (t *HashTable) Get(key int) (string, error) {
	bucket := t.findByKey(key)

	if bucket == nil {
		return "", errors.New("None")
	}

	return *bucket.Value, nil
}

func (t *HashTable) Delete(key int) (string, error) {
	bucket := t.findByKey(key)

	if bucket == nil {
		return "", errors.New("None")
	}

	value := *bucket.Value
	bucket.Value = nil

	return value, nil
}

func (t *HashTable) findHashEmptyBucket(key int) (int, error) {
	hash := 0

	for i := 0; i <= t.size; i++ {
		hash = t.hashByIteration(key, hash, i)

		if t.buckets[hash] != nil && t.buckets[hash].Value != nil && t.buckets[hash].Key != key {
			continue
		}

		return hash, nil
	}

	return 0, errors.New("out of range")
}

func (t *HashTable) findByKey(key int) *BucketValue {
	hash := 0

	for i := 0; i <= t.size; i++ {
		hash = t.hashByIteration(key, hash, i)

		if t.buckets[hash] == nil || (t.buckets[hash].Key == key && t.buckets[hash].Value == nil) {
			return nil
		}

		if t.buckets[hash].Key != key {
			continue
		}

		return t.buckets[hash]
	}

	return nil
}

func (t *HashTable) hashByIteration(key, hash, i int) int {
	return t.bucketMod(hash+t.hash(key)+i, t.size)
}

func (t *HashTable) hash(key int) int {
	return key
}

func (t *HashTable) bucketMod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	hash := NewHash(104723)

	commandMap := map[string]func(key int, value string) (string, error){
		"put": func(key int, value string) (string, error) {
			return "", hash.Put(key, value)
		},
		"delete": func(key int, _ string) (string, error) {
			return hash.Delete(key)
		},
		"get": func(key int, _ string) (string, error) {
			return hash.Get(key)
		},
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		values := strings.Split(scanner.Text(), " ")

		commandFunc := commandMap[values[0]]

		key, _ := strconv.Atoi(values[1])
		value := ""
		if len(values) > 2 {
			value = values[2]
		}

		result, err := commandFunc(key, value)

		if err != nil {
			fmt.Fprintln(writer, err)

			continue
		}

		if result == "" {
			continue
		}

		fmt.Fprintln(writer, result)
	}

	writer.Flush()
}
