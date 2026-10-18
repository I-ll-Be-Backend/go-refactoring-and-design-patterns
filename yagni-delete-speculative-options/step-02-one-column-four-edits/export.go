package lesson

/* config.yaml
export:
  format: csv
  delimiter: ","
  header: true
*/
import (
	"fmt"
	"strings"
)

type ExportConfig struct {
	Format    string
	Delimiter string
	Header    bool
}

type OrderExport struct {
	format    string
	delimiter string
	header    bool
}

func NewOrderExport(c ExportConfig) OrderExport {
	return OrderExport{c.Format, c.Delimiter, c.Header}
}

func (e OrderExport) Export(orders []Order) (string, error) {
	var out strings.Builder
	if e.header && e.format == "csv" {
		out.WriteString("id" + e.delimiter + "total" +
			e.delimiter + "postalCode\n")
	}
	for _, o := range orders {
		total := fmt.Sprintf("%.2f", o.Total)
		var row string
		switch e.format {
		case "csv":
			row = o.ID + e.delimiter + total + e.delimiter + o.PostalCode
		case "json":
			row = fmt.Sprintf(`{"id":%q,"total":%s,"postalCode":%q}`,
				o.ID, total, o.PostalCode)
		case "xml":
			row = fmt.Sprintf(`<order id="%s" total="%s" postalCode="%s"/>`,
				o.ID, total, o.PostalCode)
		default:
			return "", fmt.Errorf("bad format %q", e.format)
		}
		out.WriteString(row + "\n")
	}
	return out.String(), nil
}

type NightlyJob struct {
	config    ExportConfig
	orders    OrderRepo
	warehouse Warehouse
}

func (j NightlyJob) Run() error {
	today := j.orders.Today()
	csv, err := NewOrderExport(j.config).Export(today)
	if err != nil {
		return err
	}
	return j.warehouse.UploadCSV(csv)
}
type Order struct {
	ID         string
	Total      float64
	PostalCode string
}

type OrderRepo interface{ Today() []Order }

type Warehouse interface{ UploadCSV(csv string) error }
