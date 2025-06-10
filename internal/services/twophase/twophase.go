package twophase

type Order2Pc struct {
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
type InventoryService struct{}

func (is *InventoryService) prepare() bool {
	return true
}
func (is *InventoryService) commit() bool {

	return true
}
func (is *InventoryService) rollback() bool {

	return true
}

type PaymentService struct{}

func (is *PaymentService) prepare() bool {
	return true
}
func (is *PaymentService) commit() bool {
	return true
}
func (is *PaymentService) rollback() bool {
	return true
}

type OrderService struct{}

func (is *OrderService) prepare() bool {
	return true
}
func (is *OrderService) commit() bool {
	return true
}
func (is *OrderService) rollback() bool {
	return true
}

func CreatOrder() {
	iservice := new(InventoryService)
	pservice := new(PaymentService)
	oservice := new(OrderService)
	invent_ready := iservice.prepare()
	if invent_ready {
		p_ready := pservice.prepare()
		if p_ready {
			o_ready := oservice.prepare()
			if o_ready {
				iservice.commit()
				pservice.commit()
				oservice.commit()
				all := true
				if all {
					pservice.rollback()
					iservice.rollback()
					oservice.rollback()
				}
			} else {
				pservice.rollback()
				iservice.rollback()
			}
		} else {
			iservice.rollback()
		}
	}
}
