package main

import (
	"context"
	"fmt"
	"time"
)

/*
- создать 3 контекста со следующей схемой наследования:
 1. parentCtx <- context.Background
 2. middleCtx <- parentCtx
 3. childCtx <- middleCtx
    - создать 3 группы горутин
    - завязать каждую группу горутин на определённый контекст
    - попробовать разные комбинации порядков отмены контекста,
    оценить что происходит при этом с каждой отдельной группой горутин
*/

/*
Первый вариант комбинации порядков отмены контекста:
	childCtx -> middleCtx -> parentCtx
Вывод:
	grtC2
	Значение №3: 2
	grtM2
	Значение №2: 8
	Значение №1: 6
	grtP2
	main завершился!
*/
/*
// группа горутин №1
func grtP1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value += n
	fmt.Println("Значение №1:", value)
}

func grtP2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtP2"
	fmt.Println(s)
}

// группа горутин №2
func grtM1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value *= n
	fmt.Println("Значение №2:", value)
}

func grtM2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtM2"
	fmt.Println(s)
}

// группа горутин №3
func grtC1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value -= n
	fmt.Println("Значение №3:", value)
}

func grtC2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtC2"
	fmt.Println(s)
}

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)

	go grtP1(parentCtx, 2)
	go grtP2(parentCtx, "grtP2")
	go grtM1(middleCtx, 2)
	go grtM2(middleCtx, "grtM2")
	go grtC1(childCtx, 2)
	go grtC2(childCtx, "grtC2")

	// в таком случае средний контекст нужно отменять после отмены дочернего,
	// иначе дочерний контекст будет отменён раньше среднего
	time.Sleep(500 * time.Millisecond)
	childCancel() // отменяем дочерний контекст

	time.Sleep(500 * time.Millisecond)
	middleCancel() // отменяем средний контекст

	time.Sleep(500 * time.Millisecond)
	parentCancel() // отмена родительского контекста

	time.Sleep(250 * time.Millisecond)
	fmt.Println("main завершился!")
}
*/

/*
Второй вариант комбинации порядков отмены контекста:
	childCtx (комментируем); middleCtx -> parentCtx
Вывод:
	Значение №3: 2
	Значение №2: 8
	grtC2
	grtM2
	Значение №1: 6
	grtP2
	main завершился!
*/
/*
// группа горутин №1
func grtP1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value += n
	fmt.Println("Значение №1:", value)
}

func grtP2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtP2"
	fmt.Println(s)
}

// группа горутин №2
func grtM1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value *= n
	fmt.Println("Значение №2:", value)
}

func grtM2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtM2"
	fmt.Println(s)
}

// группа горутин №3
func grtC1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value -= n
	fmt.Println("Значение №3:", value)
}

func grtC2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtC2"
	fmt.Println(s)
}

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)
	defer childCancel() // отменяем дочерний контекст в конце main()

	go grtP1(parentCtx, 2)
	go grtP2(parentCtx, "grtP2")
	go grtM1(middleCtx, 2)
	go grtM2(middleCtx, "grtM2")
	go grtC1(childCtx, 2)
	go grtC2(childCtx, "grtC2")

	// средний контекст сам отменит childCtx и разбудит горутины, которые его ждут,
	// поэтому childCancel() можно не вызывать
	time.Sleep(500 * time.Millisecond)
	middleCancel() // отменяем средний контекст

	time.Sleep(500 * time.Millisecond)
	parentCancel() // отмена родительского контекста

	time.Sleep(250 * time.Millisecond)
	fmt.Println("main завершился!")
}
*/

/*
Третий вариант комбинации порядков отмены контекста:
	parentCtx; defer выполняет в конце middleCtx и childCtx
В таком варианте горутины выполняются параллельно,
и после отмены родительского контекста они все завершаются.
Вывод:
	Значение №2: 8
	grtP2
	Значение №1: 6
	grtM2
	grtC2
	Значение №3: 2
	main завершился!
*/
// группа горутин №1
func grtP1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value += n
	fmt.Println("Значение №1:", value)
}

func grtP2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtP2"
	fmt.Println(s)
}

// группа горутин №2
func grtM1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value *= n
	fmt.Println("Значение №2:", value)
}

func grtM2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtM2"
	fmt.Println(s)
}

// группа горутин №3
func grtC1(ctx context.Context, n int) {
	<-ctx.Done() // ждем отмены контекста
	value := 4
	value -= n
	fmt.Println("Значение №3:", value)
}

func grtC2(ctx context.Context, s string) {
	<-ctx.Done() // ждем отмены контекста
	s = "grtC2"
	fmt.Println(s)
}

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)
	defer middleCancel() // отменяем средний контекст в конце main()
	defer childCancel()  // отменяем дочерний контекст в конце main()

	go grtP1(parentCtx, 2)
	go grtP2(parentCtx, "grtP2")
	go grtM1(middleCtx, 2)
	go grtM2(middleCtx, "grtM2")
	go grtC1(childCtx, 2)
	go grtC2(childCtx, "grtC2")

	// родительский контекст отменит childCtx и middleCtx и разбудит горутины, которые их ждут
	time.Sleep(500 * time.Millisecond)
	parentCancel() // отмена родительского контекста

	time.Sleep(250 * time.Millisecond)
	fmt.Println("main завершился!")
}

/*
Вывод:
	Контексты образуют "семейное дерево". Когда отменяешь контекст, отменяются и все его потомки,
	но не его родительский контекст. Поэтому отмена дочернего контекста затрагивает только
	дочернюю группу, среднего - среднюю и дочернюю, родительского - все три группы.
	Контекст сам по себе не останавливает горутину, горутина должна сама дождаться отмены,
	например через <-ctx.Done() и решить, что делать дальше.
	Горутины работают параллельно, поэтому порядок их сообщений может меняться.
*/
