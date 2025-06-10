package services

type Order struct {
	OrderId  int
	products []string
}

/*
	Distribution transaction management.

	2.
	Products service - availabe or unavailable  --> size decrement
	Payment Service - Amount card validity
	Order Service.

	1. Two phase commit
	2. Sagas
		Chreograpy
		Orchestration
*/

/*
 1. Two phase commit (2pc) - In first phases checks availabilty and locking acros all services. There after commit in second phase.
    Locking leads to performance bottlenecks.
*/
func creatOrderNew() {

}
