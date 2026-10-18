package lesson

/* config.yaml
*/
import (
	"fmt"
	"strings"
)

func ExportOrders(orders []Order) string {
	var out strings.Builder
	out.WriteString("id,total,postalCode\n")
	for _, o := range orders {
		total := fmt.Sprintf("%.2f", o.Total)
		row := o.ID + "," + total + "," + o.PostalCode
		out.WriteString(row + "\n")
	}
	return out.String()
}

type NightlyJob struct {
	orders    OrderRepo
	warehouse Warehouse
}

func (j NightlyJob) Run() error {
	today := j.orders.Today()
	csv := ExportOrders(today)
	return j.warehouse.UploadCSV(csv)
}
type Order struct {
	ID         string
	Total      float64
	PostalCode string
}

type OrderRepo interface{ Today() []Order }

type Warehouse interface{ UploadCSV(csv string) error }
