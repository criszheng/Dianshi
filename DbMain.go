package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//func main() {
//	o := models.O
//
//	//插入
//	user := new(models.TbUser)
//	user.RtxName = "test";
//	user.Password = "123456";
//	id,_ := o.Insert(user)
//	fmt.Println(id)
//
//	//查询
//	qs := o.QueryTable("tb_user")
//	var users []*models.TbUser
//	num,_ := qs.Filter("rtx_name", "test").All(&users)// WHERE rtx_name = 'test'
//	fmt.Println(users[0])
//	fmt.Println(num)
//
//	//update
//	user2 := new(models.TbUser)
//	user2.Uid = int(id);
//	user2.RtxName = "test2";
//	id2,_ := o.Update(user2)
//	fmt.Println(id2)
//
//	//delete
//	o.Delete(user2)
//}
