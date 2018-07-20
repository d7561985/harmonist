// +build js

package main

func init() {
	TileImgs = map[string][]byte{}
	TileImgs["letter-m"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAAAAKgA/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW2HnUKtAAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAASGNyTmcAAAACMD8A
AAACQE8AAAACUF8AAAACYG8AAAACcH8AAAACgI8AAAACkJ8AAAACoK8AAAACsL8AAAACwM8AAAAC
0N8AAAAC4O+6mu6bAAAAH0lEQVQYlWNgGHLAHgSQmFCAzsclgCAGSsAemwAaAAD1LRqVM6d6rAAA
AABJRU5ErkJggg==
`)
	TileImgs["map-B"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAASElEQVQYlZWPCQoA
IAgE3f9/ukNMVy1ICJxBFxP5KoAaGAO7ZZ6gT+eDWsKWgJN1SloR1U0Uk1PQpfjBxmb488yteA4k
LvteA4nYAKSoSlXOAAAAAElFTkSuQmCC
`)
	TileImgs["map-D"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8AAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AABXYnq0tLT///9dIegXAAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAATElEQVQYlZWQQRIA
IAgC5f+frsYIdDqUNzbAxoiPQdcQxBSpQb2IvycADcdSGhIqgV7BPQYsEo0U0yO41FYA38NDXMCO
KsGzCPhPfQaIcADIYNcerAAAAABJRU5ErkJggg==
`)
	TileImgs["map-F"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAARElEQVQYlWNgoBZg
ZGRE56MJgMQgGKoAogiujhHOwyHAiATQ+WARcgWQHYxVBdR7SNai+BlVKcI05DCAmYIkgCXUYAAA
eMkAsXon9z0AAAAASUVORK5CYII=
`)
	TileImgs["map-G"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAASElEQVQYlY3PWwoA
IAhE0bn733Rgpvai/LA4iE3SR1GueMPKwSbwIwGOoNJQAYitMzCmd+hLD7A+mxCxE0bUJ+gO8ZtP
qKm2annSAJwQQSXhAAAAAElFTkSuQmCC
`)
	TileImgs["map-O"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8AAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AABXYnq0tLT///9dIegXAAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAASGNyTmcAAAACMD8A
AAACQE8AAAACUF8AAAACYG8AAAACcH8AAAACgI8AAAACkJ8AAAACoK8AAAACsL8AAAACwM8AAAAC
0N8AAAAC4O+6mu6bAAAARklEQVQYlY2Q0QoAIAwCz///6WCDsVnEhCAv7UHYSpL7gXQDxDcRr3FG
hN7KgixR36hkNoF7svsFGGAB9ABjtHaxNX3X1AF2yAClqcqzIQAAAABJRU5ErkJggg==
`)
	TileImgs["map-S"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8AAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AABXYnq0tLT///9dIegXAAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAASGNyTmcAAAACMD8A
AAACQE8AAAACUF8AAAACYG8AAAACcH8AAAACgI8AAAACkJ8AAAACoK8AAAACsL8AAAACwM8AAAAC
0N8AAAAC4O+6mu6bAAAAVUlEQVQYlYVRgQ0AIAiCy/r/q1pOhWrLVhuEmAasGPgHScd79e0GFAwX
yJl+GZYUldhFA9JtKisVKBvKrrc5cSvCQtuBE3gTlgEd0d3wQYTYpir/MAGyxgDbvjAukgAAAABJ
RU5ErkJggg==
`)
	TileImgs["map-a"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAKElEQVQYlWNgGJGA
EY3DyIgkwggFOAUYGVFFGBnRRBjRRfALoLgGAgAZ8ABUOok77gAAAABJRU5ErkJggg==
`)
	TileImgs["map-door"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAJUlEQVQYlWNgoApg
RAY4BMAkEpNmAoxQErc7UFXQx2GEAggVAABiPgCMYmpbawAAAABJRU5ErkJggg==
`)
	TileImgs["map-dreaming"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAALklEQVQYlWNgGHaA
EQTQuMhC6AJgFiOUxiWAqgdIQ3jIWmAYyRAGPNZiuhQZAAAmvwBXc8jX/AAAAABJRU5ErkJggg==
`)
	TileImgs["map-fog"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAEUlEQVQYlWNgGE6A
EQUMZTMAFoMAK2vWenQAAAAASUVORK5CYII=
`)
	TileImgs["map-foliage"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAARElEQVQYlc2O2xEA
MAQEbf9NJ2E4VBDjw57HMfs32CWRqMcj5xJuiTiXrDE2B/xIv9wzG+5D+6Os5Qr1RoCEQk1sZcQB
KnQAXoyunkIAAAAASUVORK5CYII=
`)
	TileImgs["map-footsteps"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAALElEQVQYlWNgGLyA
EQjQaRALRpMjwIAmgTCdAYkGq8DlHGQzGNAEkBRT5ncAQvAAYmbxeFUAAAAASUVORK5CYII=
`)
	TileImgs["map-frontier"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAHklEQVQYlWNgGNyA
EQjR+YwofASJECGkApupgwsAAA3YABH9HCbJAAAAAElFTkSuQmCC
`)
	TileImgs["map-g"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAATklEQVQYlZ2QSQ4A
IAgDmf9/2ogLrXqyJgZHoGLEt4AK+5ZLAQGWLCUjVc4dkCoARuAGCXefafIAy7tAiNG6MCOf9QT+
+Dfw+W6Af4moAVkPAHV0XLSkAAAAAElFTkSuQmCC
`)
	TileImgs["map-ground"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAEklEQVQYlWNgGImA
kZGQwEADAAP8AAV3EZSrAAAAAElFTkSuQmCC
`)
	TileImgs["map-h"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAASklEQVQYlbWP0QoA
IAgDd///01GxWkmPCYKeU6b0KeDuC3j0ONkTDH3LYFZdsQYMPQfwRhIViU6FPWMWTlFu2EH8dHt2
nW9TyIoGT1YAfBSNWpAAAAAASUVORK5CYII=
`)
	TileImgs["map-kill"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAPklEQVQYlWNgGGKA
EQLQ+XAREAOCGeEKGMAILsAAEWBAEoBoQghAlSPMQKXhLIQ7GDAFGFF0wNyK7hvs3gQAK9kAS2UV
wmgAAAAASUVORK5CYII=
`)
	TileImgs["map-notile"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAALklEQVQYlWNgIBIw
AgE6H1kEzEESgTIRIjAGI4JBrAya1STzGdD5dBLAAMRrAQBBhwBFq/1ziwAAAABJRU5ErkJggg==
`)
	TileImgs["map-player"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAM0lEQVQYlWNgoBNg
RALofLAIRBimmHgBqH64AMxEJAEIgzIBdJei+gXDs3h8j+l/LMEFAFrWAI/HkedyAAAAAElFTkSu
QmCC
`)
	TileImgs["map-potion"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAP0lEQVQYlcWOQQ4A
IQwCmf9/eqNmE0DjVS60k9JWeimGrkDdb6AjkASSQBIo8pcFtJBHFBPTlVslu+t2+tT1AUTpAH7I
/GBjAAAAAElFTkSuQmCC
`)
	TileImgs["map-rock"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAK0lEQVQYlWNgGNyA
EQjQ+cgijIxoImAOmgADCKFqYUA1hIEB0xqq+YEcAAAjrQAv/A1zzAAAAABJRU5ErkJggg==
`)
	TileImgs["map-s"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAT0lEQVQYlY2QQQ4A
IAjD6P8/bdDIQKKRi1DJNjX7KFrrB6QrfHYiwKwbIAOkIBUQmF5ayTmqbdKo5JxFzqQsZFu1Beu2
a49o40Fv0L83agByFgCNKWZ45gAAAABJRU5ErkJggg==
`)
	TileImgs["map-simella"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAPklEQVQYlc2OuQ0A
MAwCffsvnQL/ipQ2dDwGm30OYPOmkAKi4IJCsjwxlGpOL3p1TpuO1ja8+H7s/fqto3AANIcAX8b7
W48AAAAASUVORK5CYII=
`)
	TileImgs["map-space"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAADklEQVQYlWNgGAWD
EQAAAZgAAYLRFi4AAAAASUVORK5CYII=
`)
	TileImgs["map-stairs"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAALUlEQVQYlWNgIAwY
UQBWATCCqcYlAFGLogJTAKwIVQWmAIq1WFUMNgGCQYgGAHjfAJ9Wi0gEAAAAAElFTkSuQmCC
`)
	TileImgs["map-wall"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAVElEQVQYlY2QUQoA
IAhD5/0vHanp1IL8kHy4tQI+SnZZl5gPEAe5yjLteg4ZjoSM2jbI1ARpCnfLIB6F5gvAAEDN7Kax
S4+jyzgpgySvDypg/GmrBZ1IANWOroY0AAAAAElFTkSuQmCC
`)
	TileImgs["map-wall2"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAAPElEQVQYlWNgIAIw
IgF0PlgEKs4AZwIJqAgjA0wAIsLIgBBggAoiBJDJQSWA7lJ0v6D7FiM88IUhliAHAHczAJvEmIc9
AAAAAElFTkSuQmCC
`)
	TileImgs["map-wall3"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAASUlEQVQYlY2PUQ4A
IAhC5f6Xbs0ywLbio+StFCP+hCn3hAAlgBIpzo2s1pei1jbfcd+WA+47QHkOtCfuo0oJ8pji297X
Z9KAaQCQBADMjKNixQAAAABJRU5ErkJggg==
`)
}
