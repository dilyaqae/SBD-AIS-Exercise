package mapred

// imports
import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
// input []string
// output map[word]frequency
func (mr MapReduce) Run(input []string) map[string]int {
	kvCh := make(chan []KeyValue, len(input))
	// wait until all mapper goroutines are finished
	var wg sync.WaitGroup

	// map phase
	for _, line := range input {
		wg.Add(1)
		// start one goroutine per input string
		go func(t string) {
			defer wg.Done()
			kvCh <- mr.wordCountMapper(t)
		}(line)
	}

	go func() {
		wg.Wait()
		close(kvCh)
	}()
	// shuffle phase
	grouped := make(map[string][]int)
	for kvs := range kvCh {
		for _, kv := range kvs {
			grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
		}
	}

	// reduce phase
	out := make(map[string]int)
	var rwg sync.WaitGroup
	var mu sync.Mutex

	for key, values := range grouped {
		rwg.Add(1)
		go func(k string, vals []int) {
			defer rwg.Done()
			res := mr.wordCountReducer(k, vals)
			mu.Lock()
			out[res.Key] = res.Value
			mu.Unlock()
		}(key, values)
	}

	rwg.Wait()
	return out
}

// emit (word, 1) for every word occurrence in the input text
func (mr MapReduce) wordCountMapper(text string) []KeyValue {
	s := strings.ToLower(text)

	// remove everything that is not either a-z or whitespace
	re := regexp.MustCompile(`[^a-z\s]+`)
	s = re.ReplaceAllString(s, " ")
	// split into words by whitespace
	words := strings.Fields(s)
	kvs := make([]KeyValue, 0, len(words))
	for _, w := range words {
		kvs = append(kvs, KeyValue{Key: w, Value: 1})
	}
	return kvs
}

// word count reducer

func (mr MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	// sum all counts for the word
	for _, v := range values {
		sum += v
	}
	return KeyValue{Key: key, Value: sum}
}
