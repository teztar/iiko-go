package iiko

import (
	"encoding/json"
	"testing"
)

// Ответ iiko по категорийной скидке. Раньше ProductCategoryDiscount была пустой
// структурой, и весь этот блок молча терялся при разборе.
const categorisedDiscountResponse = `{
  "correlationId": "5b1b0e36-4e9c-4e51-8d3a-9a2b1c0d0000",
  "discounts": [
    {
      "organizationId": "0875c9eb-3800-4f1d-8500-5c493a97ba1d",
      "items": [
        {
          "id": "11111111-1111-1111-1111-111111111111",
          "name": "Happy Hours",
          "percent": 0,
          "isCategorisedDiscount": true,
          "productCategoryDiscounts": [
            {
              "categoryId": "22222222-2222-2222-2222-222222222222",
              "categoryName": "Десерты",
              "percent": 20
            },
            {
              "categoryId": "33333333-3333-3333-3333-333333333333",
              "categoryName": "Напитки",
              "percent": 10
            }
          ],
          "mode": "Percent",
          "isAutomatic": true,
          "isDeleted": false
        }
      ]
    }
  ]
}`

func TestCategorisedDiscountIsParsed(t *testing.T) {
	var response DiscountsResponse

	if err := json.Unmarshal([]byte(categorisedDiscountResponse), &response); err != nil {
		t.Fatalf("разбор ответа: %v", err)
	}

	if len(response.Discounts) != 1 {
		t.Fatalf("ожидали одну организацию, получили %d", len(response.Discounts))
	}

	items := response.Discounts[0].Items
	if len(items) != 1 {
		t.Fatalf("ожидали одну скидку, получили %d", len(items))
	}

	discount := items[0]
	if !discount.IsCategorisedDiscount {
		t.Fatal("скидка должна быть категорийной")
	}

	if len(discount.ProductCategoryDiscounts) != 2 {
		t.Fatalf("ожидали две категории, получили %d", len(discount.ProductCategoryDiscounts))
	}

	desserts := discount.ProductCategoryDiscounts[0]
	if desserts.CategoryName != "Десерты" {
		t.Fatalf("ожидали «Десерты», получили %q", desserts.CategoryName)
	}

	if desserts.Percent != 20 {
		t.Fatalf("ожидали 20 процентов, получили %v", desserts.Percent)
	}

	if desserts.CategoryID.String() != "22222222-2222-2222-2222-222222222222" {
		t.Fatalf("идентификатор категории потерян: %s", desserts.CategoryID)
	}
}

// У категорийной скидки общий percent игнорируется — это прямо написано в
// документации iiko, и путать одно с другим значит показать гостю не ту цену.
func TestCategorisedDiscountIgnoresTopLevelPercent(t *testing.T) {
	var response DiscountsResponse

	if err := json.Unmarshal([]byte(categorisedDiscountResponse), &response); err != nil {
		t.Fatalf("разбор ответа: %v", err)
	}

	discount := response.Discounts[0].Items[0]
	if discount.Percent != 0 {
		t.Fatalf("общий процент здесь нулевой: %v", discount.Percent)
	}

	for _, category := range discount.ProductCategoryDiscounts {
		if category.Percent == 0 {
			t.Fatalf("категория %q осталась без процента", category.CategoryName)
		}
	}
}
