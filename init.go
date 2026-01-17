package main

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
