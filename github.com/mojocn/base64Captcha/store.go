// Copyright 2014 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base64Captcha

import (
	"container/list"
	"sync"
	"time"
)

// Store An object implementing Store interface can be registered with SetCustomStore
// function to handle storage and retrieval of captcha ids and solutions for
// them, replacing the default memory store.
//
// It is the responsibility of an object to delete expired and used captchas
// when necessary (for example, the default memory store collects them in Set
// method after the certain amount of captchas has been stored.)
type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, digits []byte)

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) (digits []byte)
}

// expValue stores timestamp and id of captchas. It is used in the list inside
// memoryStore for indexing generated captchas by timestamp to enable garbage
// collection of expired captchas.
type idByTimeValue struct {
	timestamp time.Time
	id        string
}

// memoryStore is an internal store for captcha ids and their values.
type memoryStore struct {
	sync.RWMutex
	digitsById map[string][]byte
	idByTime   *list.List
	// Number of items stored since last collection.
	numStored int
	// Number of saved items that triggers collection.
	collectNum int
	// Expiration time of captchas.
	expiration time.Duration
}

// NewMemoryStore returns a new standard memory store for captchas with the
// given collection threshold and expiration time (duration). The returned
// store must be registered with SetCustomStore to replace the default one.
func NewMemoryStore(collectNum int, expiration time.Duration) Store {
	s := new(memoryStore)
	s.digitsById = make(map[string][]byte)
	s.idByTime = list.New()
	s.collectNum = collectNum
	s.expiration = expiration
	return s
}

func (s *memoryStore) Set(id string, digits []byte) {
	s.Lock()
	s.digitsById[id] = digits
	s.idByTime.PushBack(idByTimeValue{time.Now(), id})
	s.numStored++
	s.Unlock()
	if s.numStored > s.collectNum {
		go s.collect()
	}
}

func (s *memoryStore) Get(id string, clear bool) (digits []byte) {
	if !clear {
		// When we don't need to clear captcha, acquire read lock.
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}
	digits, ok := s.digitsById[id]
	if !ok {
		return
	}
	if clear {
		delete(s.digitsById, id)
		// XXX(dchest) Index (s.idByTime) will be cleaned when
		// collecting expired captchas.  Can't clean it here, because
		// we don't store reference to expValue in the map.
		// Maybe store it?
	}
	return
}

func (s *memoryStore) collect() {
	now := time.Now()
	s.Lock()
	defer s.Unlock()
	s.numStored = 0
	for e := s.idByTime.Front(); e != nil; {
		ev, ok := e.Value.(idByTimeValue)
		if !ok {
			return
		}
		if ev.timestamp.Add(s.expiration).Before(now) {
			delete(s.digitsById, ev.id)
			next := e.Next()
			s.idByTime.Remove(e)
			e = next
		} else {
			return
		}
	}
}
