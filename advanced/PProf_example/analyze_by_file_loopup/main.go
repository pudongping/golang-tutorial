package main

import (
	"io"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
)

func init() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}

type LookupType int8

const (
	LookupGoroutine LookupType = iota
	LookupThreadcreate
	LookupHeap
	LookupAllocs
	LookupBlock
	LookupMutex
)

func pprofLookup(lookupType LookupType, w io.Writer) error {
	var err error
	switch lookupType {
	case LookupGoroutine:
		p := pprof.Lookup("goroutine")
		err = p.WriteTo(w, 2)
	case LookupThreadcreate:
		p := pprof.Lookup("threadcreate")
		err = p.WriteTo(w, 2)
	case LookupHeap:
		p := pprof.Lookup("heap")
		err = p.WriteTo(w, 2)
	case LookupAllocs:
		p := pprof.Lookup("allocs")
		err = p.WriteTo(w, 2)
	case LookupBlock:
		p := pprof.Lookup("block")
		err = p.WriteTo(w, 2)
	case LookupMutex:
		p := pprof.Lookup("mutex")
		err = p.WriteTo(w, 2)
	}

	return err
}

func main() {
	http.HandleFunc("/lookup/heap", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupHeap, os.Stdout)
	})
	http.HandleFunc("/lookup/threadcreate", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupThreadcreate, os.Stdout)
	})
	http.HandleFunc("/lookup/block", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupBlock, os.Stdout)
	})
	http.HandleFunc("/lookup/goroutine", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupGoroutine, os.Stdout)
	})
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
