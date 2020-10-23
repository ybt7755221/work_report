package gutil

import (
	"testing"
	"work_report/dao"
	"work_report/entities"
)

func TestBeanUtil(t *testing.T) {
	usersDao := new(dao.WrUsersDao)
	users := new(entities.WrUsers)
	users.Id = 100
	users.Mobile = "11111111111"
	BeanUtil(usersDao, users)
	t.Logf("usersDao : %v", usersDao)
}

func TestFirstToLower(t *testing.T) {
	s := FirstToLower("Hello World")
	t.Log(s)
}

func TestFirstToUpper(t *testing.T) {
	s := FirstToUpper("helloWorld")
	t.Log(s)
}

func BenchmarkBeanUtil(b *testing.B) {
	usersDao := new(dao.WrUsersDao)
	users := new(entities.WrUsers)
	users.Id = 100
	users.Mobile = "11111111111"
	users.Email = "asdfoas@dsfs.com"
	users.Username = "sadfas"
	users.Password = "sdfasdfds"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BeanUtil(usersDao, users)
	}
}

func BenchmarkTwoJson(b *testing.B) {
	usersDao := new(dao.WrUsersDao)
	users := new(entities.WrUsers)
	users.Id = 100
	users.Mobile = "11111111111"
	users.Email = "asdfoas@dsfs.com"
	users.Username = "sadfas"
	users.Password = "sdfasdfds"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoJson(usersDao, users)
	}
}
