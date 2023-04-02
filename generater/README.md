# study-hybrid-programming

## 介绍

Hybrid Programming 混合编程 golang c c++

## 定义

- 函数

  - 调用函数

  - 响应函数

- 方法

  - 类的函数

- 接口封装
  将方法转换为函数的过程。

## 过程

- 调用函数

  go ➡️cgo➡️c➡️c++

- 响应函数

  c++➡️c➡️cgo➡️go

## 类型对应

| C 语言类型             | CGO 类型    | Go 语言类型    |
| ---------------------- | ----------- | -------------- |
| char                   | C.char      | byte           |
| singed char            | C.schar     | int8           |
| unsigned char          | C.uchar     | uint8          |
| short                  | C.short     | int16          |
| unsigned short         | C.ushort    | uint16         |
| int                    | C.int       | int32          |
| unsigned int           | C.uint      | uint32         |
| long                   | C.long      | int32          |
| unsigned long          | C.ulong     | uint32         |
| long long int          | C.longlong  | int64          |
| unsigned long long int | C.ulonglong | uint64         |
| float                  | C.float     | float32        |
| double                 | C.double    | float64        |
| size_t                 | C.size_t    | uint           |
| struct xx              | C.struct_xx | type xx struct |
