package singleton

import "sync"

/*
==============================
Singleton Design Pattern
==============================

Definition:
Ensures a class has only one instance in the entire program and provides a global point of access to it.

Use Case:
Useful when exactly one object is needed to coordinate actions across the system.

Examples:
- Configuration Manager
- Logging System
- Caching Mechanism
- Database Connection Pool (in some cases)

Why use it?
It avoids redundant object creation and ensures a shared resource is consistent.
*/

type singleton struct {
	Message string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{Message: "I am the only one!"}
	})
	return instance
}
