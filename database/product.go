package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price,string"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product //= make([]Product, 0)

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}
func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       100,
		ImgUrl:      "https://i.chaldn.com/_mpimage/komola-orange-imported-50-gm-1-kg?src=https%3A%2F%2Feggyolk.chaldal.com%2Fapi%2FPicture%2FRaw%3FpictureId%3D64292&q=best&v=1&m=1600",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Green",
		Price:       150,
		ImgUrl:      "https://assets.clevelandclinic.org/transform/LargeFeatureImage/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yellow",
		Price:       50,
		ImgUrl:      "https://www.orchardfood.co.za/cdn/shop/files/Bananas-2.png?v=1753697464&width=1946",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Lichu",
		Description: "Lichu is red",
		Price:       10,
		ImgUrl:      "https://media.gettyimages.com/id/566454679/photo/lychee-fruits.jpg?s=1024x1024&w=gi&k=20&c=mb13OLKoGK9QZYHi3j5nMgBfkuhSMue19wMRCrLWy5g=",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)

}
