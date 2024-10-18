#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os

if str(__name__).upper() not in ("__MAIN__",):
    raise SystemExit(1)

class ChangeDir:
    def __init__(self, new_path):
        self.new_path = os.path.abspath(new_path)
        self.saved_path = None

    def __enter__(self):
        self.saved_path = os.getcwd()
        os.chdir(self.new_path)

    def __exit__(self, e_type, value, traceback):
        os.chdir(self.saved_path)


script_dir = os.path.dirname(os.path.realpath(__file__))

def bit_range(s, n):
    started = False
    for i in range(n):
        v = 2 ** i
        if not started:
            if v < s:
                continue
            started = True
        if n <= v:
            break
        yield v

constraint = " | ".join([ f"~[{i}]T" for i in bit_range(2, 8192) ])
cases = "\n".join([ f"    case [{i}]T:\n        return GetArraySliceSize{i}(v)" for i in bit_range(2, 8192) ])
functions = "\n\n".join([ f"func GetArraySliceSize{i}[T any](data [{i}]T) []T \x7b\n    return data[:]\n\x7d" for i in bit_range(2, 8192) ])

template = """
package cores

type ArraySliceSizeNImpl[T any] interface {
    ~[]T | ~[1]T | """ + constraint + """ | ~[8192]T
}

func GetBoolArraySliceSizeN[V ArraySliceSizeNImpl[bool]](data V) []bool {
    return GetArraySliceSizeN[bool, V](data)
}

func GetIntArraySliceSizeN[V ArraySliceSizeNImpl[int]](data V) []int {
    return GetArraySliceSizeN[int, V](data)
}

func GetUintArraySliceSizeN[V ArraySliceSizeNImpl[uint]](data V) []uint {
    return GetArraySliceSizeN[uint, V](data)
}

func GetInt8ArraySliceSizeN[V ArraySliceSizeNImpl[int8]](data V) []int8 {
    return GetArraySliceSizeN[int8, V](data)
}

func GetUint8ArraySliceSizeN[V ArraySliceSizeNImpl[uint8]](data V) []uint8 {
    return GetArraySliceSizeN[uint8, V](data)
}

func GetInt16ArraySliceSizeN[V ArraySliceSizeNImpl[int16]](data V) []int16 {
    return GetArraySliceSizeN[int16, V](data)
}

func GetUint16ArraySliceSizeN[V ArraySliceSizeNImpl[uint16]](data V) []uint16 {
    return GetArraySliceSizeN[uint16, V](data)
}

func GetInt32ArraySliceSizeN[V ArraySliceSizeNImpl[int32]](data V) []int32 {
    return GetArraySliceSizeN[int32, V](data)
}

func GetUint32ArraySliceSizeN[V ArraySliceSizeNImpl[uint32]](data V) []uint32 {
    return GetArraySliceSizeN[uint32, V](data)
}

func GetInt64ArraySliceSizeN[V ArraySliceSizeNImpl[int64]](data V) []int64 {
    return GetArraySliceSizeN[int64, V](data)
}

func GetUint64ArraySliceSizeN[V ArraySliceSizeNImpl[uint64]](data V) []uint64 {
    return GetArraySliceSizeN[uint64, V](data)
}

func GetUintPtrArraySliceSizeN[V ArraySliceSizeNImpl[uintptr]](data V) []uintptr {
    return GetArraySliceSizeN[uintptr, V](data)
}

func GetBytesArraySliceSizeN[V ArraySliceSizeNImpl[byte]](data V) []byte {
    return GetArraySliceSizeN[byte, V](data)
}

func GetStringArraySliceSizeN[V ArraySliceSizeNImpl[string]](data V) []string {
    return GetArraySliceSizeN[string, V](data)
}

func GetArraySliceSizeN[T any, V ArraySliceSizeNImpl[T]](data V) []T {
    temp := Unwrap(CastAny(data))
    switch v := temp.(type) {
    case []T:
        return v
    case [1]T:
        return GetArraySliceSize1(v)\n""" + cases + """
    case [8192]T:
        return GetArraySliceSize8192(v)
    default:
        panic("Exception: Invalid data type")
    }
}

func GetArraySliceSize1[T any](data [1]T) []T {
    return data[:]
}\n""" + functions + """

func GetArraySliceSize8192[T any](data [8192]T) []T {
    return data[:]
}
"""

script_path_file = "app/cores/slices.go"

def ensure_file(path_file):
    if os.path.exists(path_file):
        if os.path.isdir(path_file):
            raise IsADirectoryError(f"Path '{path_file}' is a directory, not a file.")
        elif not os.access(path_file, os.W_OK):
            raise PermissionError(f"No write access to '{path_file}' file.")
    else:
        with open(path_file, "w"):
            pass

with ChangeDir(script_dir):
    os.chdir("..")

    ensure_file(script_path_file)

    with open(script_path_file, "w") as file:
        file.write(template)

    print(f"Template has been written to '{script_path_file}' file")
