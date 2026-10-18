package lesson

/* config.yaml
export:
*/
import (
	"fmt"
	"strings"
)

type ExportConfig struct{}

type OrderExport struct{}

func NewOrderExport(c ExportConfig) OrderExport {
	return OrderExport{}
}

func (OrderExport) Export(orders []Order) string {
	var out strings.Builder
	out.WriteString("id,total\n")
	for _, o := range orders {
		total := fmt.Sprintf("%.2f", o.Total)
		row := o.ID + "," + total
		out.WriteString(row + "\n")
	}
	return out.String()
}

type NightlyJob struct {
	config    ExportConfig
	orders    OrderRepo
	warehouse Warehouse
}

func (j NightlyJob) Run() error {
	today := j.orders.Today()
	csv := NewOrderExport(j.config).Export(today)
	return j.warehouse.UploadCSV(csv)
}
type Order struct {
	ID         string
	Total      float64
	PostalCode string
}

type OrderRepo interface{ Today() []Order }

type Warehouse interface{ UploadCSV(csv string) error }
