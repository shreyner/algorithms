package hash_table_old

import (
	"errors"
)

type BucketValue struct {
	Key   int
	Value *int
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

func (t *HashTable) Put(key, value int) error {
	hashIndex := 0

	for i := 0; i <= t.size; i++ {
		hashIndex = t.bucketMod(hashIndex+t.hash(key)+i, t.size)

		if t.buckets[hashIndex] != nil && t.buckets[hashIndex].Value != nil && t.buckets[hashIndex].Key != key {
			continue
		}

		t.buckets[hashIndex] = &BucketValue{
			Key:   key,
			Value: &value,
		}

		return nil
	}

	return errors.New("out of range")
}

func (t *HashTable) Get(key int) (int, error) {
	bucket := t.findByKey(key)

	if bucket == nil {
		return 0, errors.New("not found")
	}

	return *bucket.Value, nil
}

func (t *HashTable) Delete(key int) (int, error) {
	bucket := t.findByKey(key)

	if bucket == nil {
		return 0, errors.New("not found")
	}

	value := *bucket.Value
	bucket.Value = nil

	return value, nil
}

func (t *HashTable) findByKey(key int) *BucketValue {
	hashIndex := 0

	for i := 0; i <= t.size; i++ {
		hashIndex = t.bucketMod(hashIndex+t.hash(key)+i, t.size)

		if t.buckets[hashIndex] == nil || (t.buckets[hashIndex].Key == key && t.buckets[hashIndex].Value == nil) {
			return nil
		}

		if t.buckets[hashIndex].Key != key {
			continue
		}

		return t.buckets[hashIndex]
	}

	return nil
}

func (t *HashTable) hash(key int) int {
	return key
}

func (t *HashTable) bucketMod(a, b int) int {
	return (a%b + b) % b
}

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	writer := bufio.NewWriter(os.Stdout)
//	scanner := bufio.NewScanner(reader)
//	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
//	scanner.Split(bufio.ScanLines)
//
//	scanner.Scan()
//	n, _ := strconv.Atoi(scanner.Text())
//
//	hash := NewHash(104723)
//
//	commandMap := map[string]func(key int) (int, error){
//		"delete": hash.Delete,
//		"get":    hash.Get,
//	}
//
//	for i := 0; i < n; i++ {
//		scanner.Scan()
//
//		values := strings.Split(scanner.Text(), " ")
//
//		command := values[0]
//
//		if command == "put" {
//			key, _ := strconv.Atoi(values[1])
//			value, _ := strconv.Atoi(values[2])
//
//			if err := hash.Put(key, value); err != nil {
//				fmt.Fprintln(writer, "out of range")
//			}
//
//			continue
//		}
//
//		commandFunc, ok := commandMap[command]
//
//		if ok {
//			key, _ := strconv.Atoi(values[1])
//
//			value, err := commandFunc(key)
//
//			if err != nil {
//				fmt.Fprintln(writer, "None")
//
//				continue
//			}
//
//			fmt.Fprintln(writer, value)
//		}
//	}
//
//	writer.Flush()
//}

//func mod(a, b int) int {
//	return (a%b + b) % b
//}
//
//type BucketValue struct {
//	Key   int
//	Value *int
//}
//
//type HashTable struct {
//	size    int
//	buckets []*BucketValue
//}
//
//func NewHash(size int) *HashTable {
//	return &HashTable{
//		size:    size,
//		buckets: make([]*BucketValue, size),
//	}
//}
//
//func (t *HashTable) Put(key, value int) error {
//	h := 0
//
//	for i := 1; i <= t.size; i++ {
//		h = t.hash(key)
//
//		if t.buckets[h] != nil && t.buckets[h].Key != key {
//			continue
//		}
//
//		t.buckets[h] = &BucketValue{
//			Key:   key,
//			Value: &value,
//		}
//
//		return nil
//	}
//
//	return errors.New("sorry!")
//}
//
//func (t *HashTable) Get(key int) (int, error) {
//	h := 0
//
//	for i := 1; i <= t.size; i++ {
//		h = t.hash(key)
//
//		if t.buckets[h] == nil {
//			return 0, errors.New("not found")
//		}
//
//		if t.buckets[h] != nil && t.buckets[h].Value == nil {
//			continue
//		}
//
//		if t.buckets[h].Key != key {
//			continue
//		}
//
//		return *t.buckets[h].Value, nil
//	}
//
//	return 0, errors.New("not found")
//}
//
//func (t *HashTable) Delete(key int) (int, error) {
//	h := 0
//
//	for i := 1; i <= t.size; i++ {
//		h = t.hash(key)
//
//		if t.buckets[h] == nil {
//			return 0, errors.New("not found")
//		}
//
//		if t.buckets[h] != nil && t.buckets[h].Value == nil {
//			continue
//		}
//
//		if t.buckets[h].Key != key {
//			continue
//		}
//
//		value := *t.buckets[h].Value
//		t.buckets[h].Value = nil
//
//		return value, nil
//	}
//
//	return 0, errors.New("not found")
//}
//
//func (t *HashTable) hash(key int) int {
//	return mod(key, t.size)
//}

//func mod(a, b int) int {
//	return (a%b + b) % b
//}
//
//type BucketValue struct {
//	Key   int
//	Value *int
//}
//
//type HashTable struct {
//	size    int
//	buckets []*BucketValue
//}
//
//func NewHash(size int) *HashTable {
//	return &HashTable{
//		size:    size,
//		buckets: make([]*BucketValue, size),
//	}
//}
//
//func (t *HashTable) Put(key, value int) error {
//	h := 0
//
//	for i := 1; i <= t.size; i++ {
//		h = t.hash(key)
//
//		if t.buckets[h] != nil && t.buckets[h].Key != key {
//			continue
//		}
//
//		t.buckets[h] = &BucketValue{
//			Key:   key,
//			Value: &value,
//		}
//
//		return nil
//	}
//
//	return errors.New("sorry!")
//}
//
//func (t *HashTable) Get(key int) (int, error) {
//	h := 0
//
//	for i := 1; i <= t.size; i++ {
//		h += t.hash(key) * i
//
//		if t.buckets[h] == nil {
//			return 0, errors.New("not found")
//		}
//
//		if t.buckets[h] != nil && t.buckets[h].Value == nil {
//			continue
//		}
//
//		if t.buckets[h].Key != key {
//			continue
//		}
//
//		return *t.buckets[h].Value, nil
//	}
//
//	return 0, errors.New("not found")
//}
//
//func (t *HashTable) Delete(key int) (int, error) {
//	h := 0
//
//	for i := 1; i <= t.size; i++ {
//		h += t.hash(key) * i
//
//		if t.buckets[h] == nil {
//			return 0, errors.New("not found")
//		}
//
//		if t.buckets[h] != nil && t.buckets[h].Value == nil {
//			continue
//		}
//
//		if t.buckets[h].Key != key {
//			continue
//		}
//
//		value := *t.buckets[h].Value
//		t.buckets[h].Value = nil
//
//		return value, nil
//	}
//
//	return 0, errors.New("not found")
//}
//
//func (t *HashTable) hash(key int) int {
//	return mod(key, t.size)
//}
