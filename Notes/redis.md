# Redis

`see: https://www.redis.net.cn/tutorial/3502.html`

## 特点

- Redis支持数据的持久化，可以将内存中的数据保持在磁盘中，重启的时候可以再次加载进行使用。
- Redis不仅仅支持简单的key-value类型的数据，同时还提供list，set，zset，hash等数据结构的存储。
- Redis支持数据的备份，即master-slave模式的数据备份。

## 优势

- 性能极高 – Redis能读的速度是110000次/s,写的速度是81000次/s 。
- 丰富的数据类型 – Redis支持二进制案例的 String, List, Hash, Set 及 Ordered Set 数据类型操作。
- 原子 – Redis的所有操作都是原子性的，同时Redis还支持对几个操作合并后的原子性执行。
- 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性

## 数据类型

Redis支持五种数据类型：string（字符串），hash（哈希），list（列表），set（集合）及zset(sorted set：有序集合)。

### String（字符串）

string是redis最基本的类型，一个key对应一个value，一个键最大能存储512MB。string类型是二进制安全的，它可以包含任何数据。比如jpg图片或者序列化的对象 。

* 常用操作
```
# 设置值
SET [key] [value]
# 读取值
GET [key]
```

### Hash（哈希）

Redis hash是一个string类型的field和value的映射表，每个hash可以存储 2^32 - 1 个键值对（40多亿）。hash特别适合用于存储对象。

* 常用操作
```
# 设置值
HMSET [key] [field1] [value1] [field2] [value2] ...
# 读取值
HGETALL [key]
HGET [key] [field]
```

### List（列表）

Redis 列表是简单的字符串列表，按照插入顺序排序。可以添加一个元素到列表的头部（左边）或者尾部（右边）。列表最多可存储 2^32 - 1 个元素。

### Set（集合）

Redis的Set是string类型的无序集合。集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是O(1)。

* 常用操作
```
# 添加元素到set
SADD [key] [member]
# 取出集合的元素列表
SMEMEBERS [key]
```

### ZSet(Sorted Set：有序集合)

Redis zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。 不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。zset的成员是唯一的,但分数(score)却可以重复。

* 常用操作
```
# 添加元素
ZADD [key] [score] [member]
# 取出指定分数范围的元素
ZRANGEBYSCORE [key] [minScore] [maxScore]
```



