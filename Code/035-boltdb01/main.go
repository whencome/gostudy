package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"

	"github.com/boltdb/bolt"
)

func main() {
	// 打开数据库连接
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Panicf("open my.db fail: %v", err)
	}
	defer db.Close()

	// define data
	var userHobbies = map[string]string{
		"eric":  "basketball & football",
		"john":  "drawing",
		"smith": "dancing && singing",
	}
	var userScores = map[string]float64{
		"eric":  97.5,
		"john":  83,
		"smith": 59.6,
	}

	// create buckets
	err = createBucket01(db, "userHobbies")
	if err != nil {
		log.Panic("create bucket userHobbies fail")
	}
	err = createBucket02(db, "userScores")
	if err != nil {
		log.Panic("create bucket userScores fail")
	}

	// store data
	// store user hobbies
	for userName, hobby := range userHobbies {
		saveValue(db, "userHobbies", userName, []byte(hobby))
	}
	// store user scores
	for userName, score := range userScores {
		// convert float64 to []byte
		bits := math.Float64bits(score)
		scoreBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(scoreBytes, bits)
		saveValue(db, "userScores", userName, scoreBytes)
	}

	// fetch data
	for userName, hobby := range userHobbies {
		hobbyBytes, err := getValue(db, "userHobbies", userName)
		if err != nil {
			log.Panicf("get user %s's hobby fail: %s", userName, err)
		}
		fmt.Printf("%s's hobby: got %s; expect %s\n", userName, string(hobbyBytes), hobby)
	}
	// store user scores
	for userName, score := range userScores {
		scoreBytes, err := getValue(db, "userScores", userName)
		if err != nil {
			log.Panicf("get user %s's score fail: %s", userName, err)
		}
		// convert []byte to float64
		bits := binary.LittleEndian.Uint64(scoreBytes)
		scoreVal := math.Float64frombits(bits)
		fmt.Printf("%s's score: got %.2f; expect %.2f\n", userName, scoreVal, score)
	}

}

func createBucket01(db *bolt.DB, name string) error {
	// create a bucket
	return db.Update(func(tx *bolt.Tx) error {
		// 先尝试直接获取bucket，如果为nil，表示bucket不存在
		// 也可以直接使用CreateBucketIfNotExists方法
		b := tx.Bucket([]byte(name))
		if b == nil {
			_, err := tx.CreateBucket([]byte(name))
			return err
		}
		// 返回nil表示创建成功
		return nil
	})
}

func createBucket02(db *bolt.DB, name string) error {
	// create a bucket
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})
}

func saveValue(db *bolt.DB, bucketNmae, key string, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketNmae))
		if b == nil {
			return fmt.Errorf("bucket %s not exists", bucketNmae)
		}
		// 设置值
		return b.Put([]byte(key), value)
	})
}

func getValue(db *bolt.DB, bucketNmae, key string) ([]byte, error) {
	var val []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketNmae))
		if b == nil {
			return fmt.Errorf("bucket %s not exists", bucketNmae)
		}
		val = b.Get([]byte(key))
		return nil
	})
	return val, err
}
