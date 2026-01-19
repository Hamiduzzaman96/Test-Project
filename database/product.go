package database

type Products struct {
	ID          int
	Title       string
	Description string
	price       float64
	ImageURL    string
}

var productList []Products //empty slice

func Store(p Products) Products {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Products {
	return productList
}

func Get(productID int) *Products {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Products) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Products = make([]Products, 0)

	for idx, p := range productList {
		if p.ID != productID {
			tempList[idx] = p
		}
	}
	productList = tempList
}

func init() {
	prd1 := Products{
		ID:          1,
		Title:       "Orange",
		Description: "This is orange and i love it",
		price:       20.5,
		ImageURL:    "https://www.istockphoto.com/photos/orange",
	}
	prd2 := Products{
		ID:          2,
		Title:       "Apple",
		Description: "This is apple and i love it",
		price:       100,
		ImageURL:    "https://www.istockphoto.com/photos/green-apple",
	}
	prd3 := Products{
		ID:          3,
		Title:       "Grape",
		Description: "This is Grape and i love it",
		price:       60,
		ImageURL:    "https://www.istockphoto.com/photos/red-grape",
	}
	prd4 := Products{
		ID:          4,
		Title:       "Banana",
		Description: "This is Banana and i love it",
		price:       70,
		ImageURL:    "https://stock.adobe.com/search?k=picture+of+banana",
	}

	productList = append(productList, prd1, prd2, prd3, prd4) //empty slice a append kora hoyece
}
