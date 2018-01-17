// Go Threading Example
// @author Casey Charlesworth

package main

import "fmt"
import "sync"
import "math/rand"
import "strings"
import "time"

var wg sync.WaitGroup


// Customer is safe to use concurrently if Mutex is used.
type Customer struct {
  name string
  id int
  balance   int
	mux sync.Mutex // Locks thread
}

// Withdraw amount from account balance.
func (c *Customer) Withdraw(amount int){
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the balance
  fmt.Println("Withdrawing...\nBalance:",c.balance)
  if c.balance > amount {
    c.balance = c.balance - amount;
  }
  //fmt.Println("Balance Withdraw", c.balance)
	c.mux.Unlock()
  defer wg.Done()
}


func (c *Customer) WithdrawBad(n int){


    c.balance = c.balance - n;
time.Sleep(1000 * time.Millisecond)
  fmt.Println(c.balance)
  defer wg.Done()
}

func (c *Customer) WithdrawBad2(n int){
  fmt.Println(c.balance)

    c.balance = c.balance - n;

  //time.Sleep(100 * time.Millisecond)
  defer wg.Done()
}

// Return multiple results
func SplitString(str string)(a, b string){
  var array []string
  array = strings.SplitAfter(str, ",")
  array[0] = strings.Replace(array[0],",", "", 1)
  return array[0], array[1]
}



func main() {

  var x Customer
  x.name = "John,Doe"
  x.balance = 500
  x.id = rand.Intn(1000)

  fmt.Println("*** Starting ***")
  fmt.Print("Name: ")
  fmt.Println(SplitString(x.name))
  fmt.Println("Balance:", x.balance, "\nAccount Number:", x.id)
  //x.Withdraw(100)

  wg.Add(5)
  for i := 0; i < 5; i++ {
    //go Withdraw(&x, 100)
    go x.Withdraw(100)

  }

  wg.Wait()
  fmt.Println("New Balance:", x.balance)
  /*
  var y Customer
  y.name = "Jane Doe"
  y.balance = 500
  y.id = rand.Intn(1000)
  wg.Add(20)

  for i:= 0; i < 5; i++ {
  if(y.balance > 5){  go y.WithdrawBad(5)}
  if(y.balance > 495){  go y.WithdrawBad(155)}
  if(y.balance > 50){  go y.WithdrawBad(50)}
  if(y.balance > 20) { go y.WithdrawBad(20)}
  }
  wg.Wait()
  fmt.Println("New Balance:", y.balance)
  */
  fmt.Println("*** Done ***")

}
