package main

import "log"

var deposits = make(chan int) // send amount to deposit
var draws = make(chan draw)
var balances = make(chan int) // receive balance

type draw struct {
	int
	res chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	res := make(chan bool)
	draws <- draw{amount, res}
	return <-res
}

func init() {
	go teller() // start the monitor goroutine
}

// 练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
// 这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。
func main() {
	log.Println(Balance())

	Deposit(100)
	log.Println(Balance())

	log.Println(Withdraw(50))
	log.Println(Balance())

	log.Println(Withdraw(50))
	log.Println(Balance())

	log.Println(Withdraw(50))
	log.Println(Balance())
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case d := <-draws:
			if d.int > balance {
				d.res <- false
				continue
			}

			balance -= d.int
			d.res <- true
		}
	}
}
