package models

type Product struct {
	Id        string   `bson:"id" json:"id" example:"coles-muslie-almond"`
	Name      string   `bson:"name" json:"name" example:"Muesli (Almond)"`
	Tags      []string `bson:"tags" json:"tags" example:"[grain, carb]"`
	Nutrition `bson:"nutrition-per-100g" json:"nutrition"`
}
type Nutrition struct {
	Energy             int64   `bson:"energy" json:"energy" example:"1560"`
	Protein            float64 `bson:"protein" json:"protein" example:"12.3"`
	Fat                float64 `bson:"fat" json:"fat" example:"9.9"`
	SaturatedFat       float64 `bson:"saturated-fat" json:"saturatedFat" example:"2.8"`
	TransFat           float64 `bson:"trans-fat" json:"transFat" example:"2.8"`
	PolyunsaturatedFat float64 `bson:"polyunsaturated-fat" json:"polyunsaturatedFat" example:"2.8"`
	MonounsaturatedFat float64 `bson:"monounsaturated-fat" json:"monounsaturatedFat" example:"2.8"`
	Carbohydrate       float64 `bson:"carbohydrate" json:"carbohydrate" example:"2.8"`
	Sugars             float64 `bson:"sugars" json:"sugars" example:"2.8"`
	DietaryFibre       float64 `bson:"dietary-fibre " json:"dietaryFibre " example:"2.8"`
	Sodium             float64 `bson:"sodium" json:"sodium " example:"2"`
	Potassium          float64 `bson:"potassium" json:"potassium " example:"2"`
	Calcium            float64 `bson:"calcium" json:"calcium " example:"2"`
	VitaminE           float64 `bson:"vitamin-e" json:"vitaminE " example:"2"`
}

type ProductSearch struct {
	Product string `form:"product" bson:"-" json:"product" example:"Muesli (Almond)"`
	Number  int    `form:"number" bson:"-" json:"number" example:"1"`
}

/*
{
    "_id": {
        "$oid": "6102e72191e927acf869ae87"
    },
    "id": "coles-muslie-almond",
    "name": "Muesli (Almond)",
    "nutrition-per-100g": {
        "energy": 1560,
        "protein": 12.3,
        "fat": 9.9,
        "saturated-fat": 2.8,
        "carbohydrate": 51.7,
        "sugars": 19.7,
        "dietary-fibre": 13,
        "sodium": 6
    },
    "tags": ["grain", "carb"]
}
*/
