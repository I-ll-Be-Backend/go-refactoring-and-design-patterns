package lesson

/* config.yaml
export:
  header: true
*/
import (
	"fmt"
	"strings"
)

type ExportConfig struct {
	Header bool
}

type OrderExport struct {
	header bool
}

func NewOrderExport(c ExportConfig) OrderExport {
	return OrderExport{c.Header}
}

func (e OrderExport) Export(orders []Order) string {
	var out strings.Builder
	if e.header {
		out.WriteString("id,total\n")
	}
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
