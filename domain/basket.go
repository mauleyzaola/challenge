package domain

import "fmt"

type Basket struct {
	Id     string
	Items  BasketItems
	Amount float64
}

func (this *Basket) Print() {
	fmt.Printf("id:%s\tAmount:%v\n", this.Id, this.Amount)
	if len(this.Items) != 0 {
		fmt.Printf("code\tqty\tproduct\tprice\n")
		for _, v := range this.Items {
			fmt.Printf("%s\t%d\t%s\t%v\n", v.Product.Code, v.Quantity, v.Product.Name, v.Product.Price)
		}
	}
}
