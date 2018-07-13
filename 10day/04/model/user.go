package model

const (
	UserStatusOnline  = 1
	UserStatusOffline = iota
)


// 这个将来要被 json 的，所以最好就是 int 和 string
type User struct {
	Id            int
	Passwd        string
	Name          string
	Age           int
	Status        int
	LastLoginTime string
}