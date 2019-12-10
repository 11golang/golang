package main
type Snowpea interface {
	shoot()
	retard()
}
type Repeatshoot interface {
	repeat()
}
type Repeater struct{
	Snowpea
}
func(stu Repeater)shoot(){
}
func(stu Repeater)repeat(){
}
func main(){
	var a Snowpea
	var b Repeatshoot
	var rep Repeater
	a=rep
	b=rep
	a.shoot()
	b.repeat()
}
