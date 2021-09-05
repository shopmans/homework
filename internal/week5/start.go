package week5

import "fmt"

func Start() {
	///////////////////////////////////////////////
	// Redis 3.0.504 (00000000/0) 64 bit
	///////////////////////////////////////////////
	//
	// 并发连接50，10000请求，keep alive = true
	//

	fmt.Println("============================= GET =============================")
	// 10字节，GET
	pyload_10byte_get := "GET 并发连接50，100000请求，keep alive = true，payload = 10byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"134228.19 requests per second\n"
	fmt.Println(pyload_10byte_get)

	// 20字节，GET
	pyload_20byte_get := "GET 并发连接50，100000请求，keep alive = true，payload = 20byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"143884.89 requests per second\n"
	fmt.Println(pyload_20byte_get)

	// 50字节，GET
	pyload_50byte_get := "GET 并发连接50，100000请求，keep alive = true，payload = 50byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"141043.72 requests per second\n"
	fmt.Println(pyload_50byte_get)

	// 100字节，GET
	pyload_100byte_get := "GET 并发连接50，100000请求，keep alive = true，payload = 100byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"140646.97 requests per second\n"
	fmt.Println(pyload_100byte_get)

	// 200字节，GET
	pyload_200byte_get := "GET 并发连接50，100000请求，keep alive = true，payload = 200byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"141442.72 requests per second\n"
	fmt.Println(pyload_200byte_get)

	// 1024字节，GET
	pyload_1024byte_get := "GET 并发连接50，100000请求，keep alive = true，payload = 1024byte\n" +
		"-----------------------------------------------------------\n" +
		"99.93% <= 1 milliseconds\n" +
		"100.00% <= 1 milliseconds\n" +
		"142247.52 requests per second\n"
	fmt.Println(pyload_1024byte_get)

	// 5120字节，GET
	pyload_5120byte_get := "GET 并发连接50，100000请求，keep alive = true，payload = 5120byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"143061.52 requests per second\n"
	fmt.Println(pyload_5120byte_get)

	fmt.Println("============================= SET =============================")
	// 10字节，SET
	pyload_10byte_set := "SET 并发连接50，100000请求，keep alive = true，payload = 10byte\n" +
		"-----------------------------------------------------------\n" +
		"99.89% <= 1 milliseconds\n" +
		"99.95% <= 3 milliseconds\n" +
		"100.00% <= 3 milliseconds\n" +
		"130890.05 requests per second\n"
	fmt.Println(pyload_10byte_set)

	// 20字节，SET
	pyload_20byte_set := "SET 并发连接50，100000请求，keep alive = true，payload = 20byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"137551.58 requests per second\n"
	fmt.Println(pyload_20byte_set)

	// 50字节，SET
	pyload_50byte_set := "SET 并发连接50，100000请求，keep alive = true，payload = 50byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"139664.80 requests per second\n"
	fmt.Println(pyload_50byte_set)

	// 100字节，SET
	pyload_100byte_set := "SET 并发连接50，100000请求，keep alive = true，payload = 100byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"139082.06 requests per second\n"
	fmt.Println(pyload_100byte_set)

	// 200字节，SET
	pyload_200byte_set := "SET 并发连接50，100000请求，keep alive = true，payload = 200byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"136425.66 requests per second\n"
	fmt.Println(pyload_200byte_set)

	// 1024字节，SET
	pyload_1024byte_set := "SET 并发连接50，100000请求，keep alive = true，payload = 1024byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"137362.64 requests per second\n"
	fmt.Println(pyload_1024byte_set)

	// 5120字节，SET
	pyload_5120byte_set := "SET 并发连接50，100000请求，keep alive = true，payload = 5120byte\n" +
		"-----------------------------------------------------------\n" +
		"100.00% <= 0 milliseconds\n" +
		"128205.13 requests per second\n"
	fmt.Println(pyload_5120byte_set)

	///////////////////////////////////////////////
	// Redis 3.0.504 (00000000/0) 64 bit
	///////////////////////////////////////////////
	//
	// 并发连接50，10000请求，keep alive = true
	// 计算每个KV平均内存点用，分别使用 10，20，50，100，200，1024，5120写入
	// 计算方法：flushall清空数据库。（写入前used_memory - 写入后used_memory） / 写入数量

	/////////////////////////////////////////// set 10 byte
	fmt.Println("============================= SET 10000 10 byte KV =============================")
	fmt.Println("============================= Before used_memory 712488 byte =============================")
	fmt.Println("============================= After used_memory 1942680 byte =============================")
	fmt.Println("============================= Avg item 123 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 80 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 250000 10 byte KV =============================")
	fmt.Println("============================= Before used_memory 691608 byte =============================")
	fmt.Println("============================= After used_memory 30789328 byte =============================")
	fmt.Println("============================= Avg item 120 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 80 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 500000 10 byte KV =============================")
	fmt.Println("============================= Before used_memory 692312 byte =============================")
	fmt.Println("============================= After used_memory 60886616 byte =============================")
	fmt.Println("============================= Avg item 120 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 80 byte =============================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	/////////////////////////////////////////// set 20 byte
	fmt.Println("============================= SET 10000 20 byte KV =============================")
	fmt.Println("============================= Before used_memory 692448 byte =============================")
	fmt.Println("============================= After used_memory 1943520 byte =============================")
	fmt.Println("============================= Avg item 125 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 88 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 250000 20 byte KV =============================")
	fmt.Println("============================= Before used_memory 692448 byte =============================")
	fmt.Println("============================= After used_memory 30789600 byte =============================")
	fmt.Println("============================= Avg item 120 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 88 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 500000 20 byte KV =============================")
	fmt.Println("============================= Before used_memory 692584 byte =============================")
	fmt.Println("============================= After used_memory 60886888 byte =============================")
	fmt.Println("============================= Avg item 120 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 88 byte =============================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	/////////////////////////////////////////// set 50 byte
	fmt.Println("============================= SET 10000 50 byte KV =============================")
	fmt.Println("============================= Before used_memory 692720 byte =============================")
	fmt.Println("============================= After used_memory 2263792 byte =============================")
	fmt.Println("============================= Avg item 157 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 120 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 250000 50 byte KV =============================")
	fmt.Println("============================= Before used_memory 692720 byte =============================")
	fmt.Println("============================= After used_memory 38789872 byte =============================")
	fmt.Println("============================= Avg item 152 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 120 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 500000 50 byte KV =============================")
	fmt.Println("============================= Before used_memory 692720 byte =============================")
	fmt.Println("============================= After used_memory 76887160 byte =============================")
	fmt.Println("============================= Avg item 152 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 120 byte =============================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	/////////////////////////////////////////// set 100 byte
	fmt.Println("============================= SET 10000 100 byte KV =============================")
	fmt.Println("============================= Before used_memory 692992 byte =============================")
	fmt.Println("============================= After used_memory 2744064 byte =============================")
	fmt.Println("============================= Avg item 205 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 176 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 250000 100 byte KV =============================")
	fmt.Println("============================= Before used_memory 692992 byte =============================")
	fmt.Println("============================= After used_memory 50790144 byte =============================")
	fmt.Println("============================= Avg item 200 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 176 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 500000 100 byte KV =============================")
	fmt.Println("============================= Before used_memory 693128 byte =============================")
	fmt.Println("============================= After used_memory 100887432 byte =============================")
	fmt.Println("============================= Avg item 200 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 176 byte =============================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	/////////////////////////////////////////// set 200 byte
	fmt.Println("============================= SET 10000 200 byte KV =============================")
	fmt.Println("============================= Before used_memory 693264 byte =============================")
	fmt.Println("============================= After used_memory 3864336 byte =============================")
	fmt.Println("============================= Avg item 317 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 288 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 250000 200 byte KV =============================")
	fmt.Println("============================= Before used_memory 693264 byte =============================")
	fmt.Println("============================= After used_memory 78790416 byte =============================")
	fmt.Println("============================= Avg item 312 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 288 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 500000 200 byte KV =============================")
	fmt.Println("============================= Before used_memory 693400 byte =============================")
	fmt.Println("============================= After used_memory 156887704 byte =============================")
	fmt.Println("============================= Avg item 312 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 288 byte =============================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	/////////////////////////////////////////// set 1024 byte
	fmt.Println("============================= SET 10000 1024 byte KV =============================")
	fmt.Println("============================= Before used_memory 693536 byte =============================")
	fmt.Println("============================= After used_memory 14424608 byte =============================")
	fmt.Println("============================= Avg item 1373 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 1344 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 250000 1024 byte KV =============================")
	fmt.Println("============================= Before used_memory 693536 byte =============================")
	fmt.Println("============================= After used_memory 342790688 byte =============================")
	fmt.Println("============================= Avg item 1368 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 1344 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 500000 1024 byte KV =============================")
	fmt.Println("============================= Before used_memory 693672 byte =============================")
	fmt.Println("============================= After used_memory 684887976 byte =============================")
	fmt.Println("============================= Avg item 1368 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 1344 byte =============================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	/////////////////////////////////////////// set 5120 byte
	fmt.Println("============================= SET 10000 5120 byte KV =============================")
	fmt.Println("============================= Before used_memory 693808 byte =============================")
	fmt.Println("============================= After used_memory 83544880 byte =============================")
	fmt.Println("============================= Avg item 8285 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 6208 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 250000 5120 byte KV =============================")
	fmt.Println("============================= Before used_memory 693944 byte =============================")
	fmt.Println("============================= After used_memory 2070791096 byte =============================")
	fmt.Println("============================= Avg item 8280 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 6208 byte =============================")
	fmt.Println("")
	fmt.Println("============================= SET 500000 5120 byte KV =============================")
	fmt.Println("============================= Before used_memory 694080 byte =============================")
	fmt.Println("============================= After used_memory 4140888384 byte =============================")
	fmt.Println("============================= Avg item 8280 byte =============================")
	fmt.Println("============================= redis-memory-for-key: 6208 byte =============================")
}
