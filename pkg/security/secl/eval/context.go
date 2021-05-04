// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package eval

import (
	"time"
	"unsafe"
)

// Context describes the context used during a rule evaluation
type Context struct {
	Object unsafe.Pointer

	Registers Registers

	// cache available across all the evaluations
	Cache map[string]unsafe.Pointer

	now time.Time
}

// Now return and cache the `now` timestamp
func (c *Context) Now() time.Time {
	if c.now.IsZero() {
		c.now = time.Now()
	}
	return c.now
}

// SetObject set the given object to the context
func (c *Context) SetObject(obj unsafe.Pointer) {
	c.Object = obj
}

// NewContext return a new Context
func NewContext(obj unsafe.Pointer) *Context {
	return &Context{
		Object: obj,
		Cache:  make(map[string]unsafe.Pointer),
	}
}
