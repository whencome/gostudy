package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Sex   string `db:"sex"`
	Email string `db:"email"`
}

// 获取数据库连接
func getConn() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ddl_test")
	if err != nil {
		log.Panicf("connect db failed: %s\n", err)
		return nil
	}
	return db
}

func main() {
	db := getConn()
	defer db.Close()
	// 初始化待插入数据
	p := Person{
		Name:  "eric01",
		Sex:   "Male",
		Email: "eric01@domain.com",
	}
	// 插入数据
	rs, err := db.Exec("insert into person(name,sex,email) values(?,?,?)", p.Name, p.Sex, p.Email)
	if err != nil {
		log.Printf("insert data failed: %s\n", err)
		return
	}
	p.Id, _ = rs.LastInsertId()
	log.Println("insert person success, id = ", p.Id)

	// 查询数据
	var newPerson Person
	err = db.Get(&newPerson, "select * from person where id = ?", p.Id)
	if err != nil {
		log.Printf("query person failed: %s\n", err)
		return
	}
	log.Printf("query person success, person = %+v\n", newPerson)

	// 更新数据
	rs, err = db.Exec("update person set email = 'eric01@domain.cn' where id = ?", p.Id)
	if err != nil {
		log.Printf("update data failed: %s\n", err)
		return
	}
	affectedRows, _ := rs.RowsAffected()
	log.Println("update data success, affected rows = ", affectedRows)

	// 使用事务批量插入数据
	tx, err := db.Begin()
	if err != nil {
		log.Printf("begin tx failed: %s\n", err)
		return
	}
	err = execTx(tx)
	if err != nil {
		tx.Rollback()
		log.Printf("exec tx failed: %s\n", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Printf("commit tx failed: %s\n", err)
		return
	}

	// query all data
	showAllPersons(db)

	// delete data
	rs, err = db.Exec("delete from person where name in ('jack01', 'smith')")
	if err != nil {
		log.Printf("delete data failed: %s\n", err)
		return
	}
	affectedRows, _ = rs.RowsAffected()
	log.Printf("deleted rows: %d\n", affectedRows)

	showAllPersons(db)
}

func execTx(tx *sql.Tx) error {
	data := [][3]string{
		{"jack01", "male", "jack01@yahoo.com"},
		{"brown02", "female", "brown02@yahoo.com"},
		{"smith", "male", "smith@yahoo.com"},
		{"elly04", "female", "elly04@yahoo.com"},
	}
	for _, row := range data {
		_, err := tx.Exec("insert into person(name,sex,email) values(?,?,?)", row[0], row[1], row[2])
		if err != nil {
			log.Printf("insert data failed: %s\n", err)
			return err
		}
	}
	return nil
}

func showAllPersons(db *sqlx.DB) {
	var persons []Person
	err := db.Select(&persons, "select * from person order by id asc")
	if err != nil {
		log.Printf("query person list failed: %s\n", err)
		return
	}
	if len(persons) == 0 || persons == nil {
		log.Println("found no person at all")
		return
	}
	log.Println("person list:\nid,  name,  sex,  email")
	for _, p := range persons {
		log.Printf("%d, %s, %s, %s\n", p.Id, p.Name, p.Sex, p.Email)
	}
	return
}
