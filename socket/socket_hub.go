// Copyright 2017 HenryLee. All Rights Reserved.
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

package socket

import (
	"github.com/henrylee2cn/goutil"
)

// SocketHub sockets hub
type SocketHub struct {
	// key: socket id (ip, name and so on)
	// value: Socket
	sockets goutil.Map
}

// NewSocketHub creates a new sockets hub.
func NewSocketHub() *SocketHub {
	chub := &SocketHub{
		sockets: goutil.AtomicMap(),
	}
	return chub
}

// Set sets a Socket.
func (sh *SocketHub) Set(id string, socket Socket) {
	_socket, loaded := sh.sockets.LoadOrStore(id, socket)
	if !loaded {
		return
	}
	if oldSocket := _socket.(Socket); socket != oldSocket {
		oldSocket.Close()
	}
}

// Get gets Socket by id.
// If second returned arg is false, mean the Socket is not found.
func (sh *SocketHub) Get(id string) (Socket, bool) {
	_socket, ok := sh.sockets.Load(id)
	if !ok {
		return nil, false
	}
	return _socket.(Socket), true
}

// Range calls f sequentially for each id and Socket present in the socket hub.
// If f returns false, range stops the iteration.
func (sh *SocketHub) Range(f func(string, Socket) bool) {
	sh.sockets.Range(func(key, value interface{}) bool {
		return f(key.(string), value.(Socket))
	})
}

// Random gets a Socket randomly.
// If third returned arg is false, mean no Socket is exist.
func (sh *SocketHub) Random() (string, Socket, bool) {
	id, socket, exist := sh.sockets.Random()
	if !exist {
		return "", nil, false
	}
	return id.(string), socket.(Socket), true
}

// Len returns the length of the socket hub.
// Note: the count implemented using sync.Map may be inaccurate.
func (sh *SocketHub) Len() int {
	return sh.sockets.Len()
}

// Delete deletes the Socket for a id.
func (sh *SocketHub) Delete(id string) {
	sh.sockets.Delete(id)
}