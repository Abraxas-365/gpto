package models

// import "strings"

/*Producto como se ve en Magento*/
type ProductMagento struct {
	Id                  *int                     `json:"id,omitempty"`
	Name                string                   `json:"name"`
	Visibility          int                      `json:"visibility"`
	Sku                 string                   `json:"sku"`
	Price               int                      `json:"price"`
	Status              StatusInt                `json:"status"`
	FamilyId            int                      `json:"attribute_set_id"`
	CustomAttributes    ProductAttributesMagento `json:"custom_attributes"`
	ExtensionAttributes ExtensionAttributes      `json:"extension_attributes"`
}
type ProductsMagento []ProductMagento

type ProductAttributeMagento struct {
	AttributeCode string `json:"attribute_code"`
	Value         any    `json:"value"`
}
type ProductAttributesMagento []ProductAttributeMagento

type ProductAttributeMagentoBuilder interface {
	AddAttributeCode(string) ProductAttributeMagentoBuilder
	AddValue(any) ProductAttributeMagentoBuilder
	Build() ProductAttributeMagento
}

func (pam *ProductAttributeMagento) AddAttributeCode(code string) ProductAttributeMagentoBuilder {
	pam.AttributeCode = code
	return pam
}

func (pam *ProductAttributeMagento) AddValue(value any) ProductAttributeMagentoBuilder {
	pam.Value = value
	return pam
}

func (pam *ProductAttributeMagento) Build() ProductAttributeMagento {
	return ProductAttributeMagento{
		AttributeCode: pam.AttributeCode,
		Value:         pam.Value,
	}
}

func NewProductAttributeMagentoBuilder() ProductAttributeMagentoBuilder {
	return &ProductAttributeMagento{}
}

type ExtensionAttributes struct {
	CategoryLinks CategoryLinks `json:"category_links"`
	StockItem     struct {
		IsInStock StatusBool `json:"is_in_stock"`
	} `json:"stock_item"`
}

type ProductMagentoBuilder interface {
	//TODO ADD ISINSTOCK
	AddSku(string) ProductMagentoBuilder
	AddName(string) ProductMagentoBuilder
	AddAttributeSet(int) ProductMagentoBuilder //en akeneo esto es la familia
	AddPrice(int) ProductMagentoBuilder
	AddStatus(int) ProductMagentoBuilder
	AddVisibility(int) ProductMagentoBuilder //en akeneo esto es un string
	AddAttributes(ProductAttributeMagento) ProductMagentoBuilder
	Build() ProductMagento
}

func (pm *ProductMagento) AddSku(sku string) ProductMagentoBuilder {
	pm.Sku = sku
	return pm
}

func (pm *ProductMagento) AddName(name string) ProductMagentoBuilder {
	pm.Name = name
	return pm
}

func (pm *ProductMagento) AddPrice(price int) ProductMagentoBuilder {
	pm.Price = price
	return pm
}

func (pm *ProductMagento) AddAttributeSet(set int) ProductMagentoBuilder {
	pm.FamilyId = set
	return pm
}

func (pm *ProductMagento) AddStatus(status int) ProductMagentoBuilder {
	pm.Status = StatusInt(status)
	return pm
}

func (pm *ProductMagento) AddVisibility(visibility int) ProductMagentoBuilder {
	pm.Visibility = visibility
	return pm
}

func (pm *ProductMagento) AddAttributes(pam ProductAttributeMagento) ProductMagentoBuilder {
	pm.CustomAttributes = append(pm.CustomAttributes, pam)
	return pm
}

func (pm *ProductMagento) Build() ProductMagento {
	return ProductMagento{
		Name:                pm.Name,
		Visibility:          pm.Visibility,
		Sku:                 pm.Sku,
		Price:               pm.Price,
		Status:              pm.Status,
		FamilyId:            pm.FamilyId,
		ExtensionAttributes: pm.ExtensionAttributes,
		CustomAttributes:    pm.CustomAttributes,
	}
}

func NewProductMagentoBuilder() ProductMagentoBuilder {
	return &ProductMagento{
		CustomAttributes: ProductAttributesMagento{},
	}
}

/*Producto como se ve en Akeneo*/
type ProductAkeneo struct {
	Identifier string            `json:"identifier"`
	Categories []string          `json:"categories"`
	Enabled    StatusBool        `json:"enabled"`
	Family     string            `json:"family"`
	Values     map[string]Values `json:"values"`
	Locale     *string           `json:"-"`
	Scope      *string           `json:"-"`
}
type ProductsAkeneo []ProductAkeneo

type ProductAkeneoBuilder interface {
	AddLocale(*string) ProductAkeneoBuilder
	AddScope(*string) ProductAkeneoBuilder
	AddIdentifier(string) ProductAkeneoBuilder
	AddName(string) ProductAkeneoBuilder
	AddIsInStock(string) ProductAkeneoBuilder
	AddSku(string) ProductAkeneoBuilder
	AddPrice(int) ProductAkeneoBuilder
	AddStatus(bool) ProductAkeneoBuilder
	AddFamily(string) ProductAkeneoBuilder
	AddValues(attributesCode string, value interface{}, akeneoType string) ProductAkeneoBuilder
	AddCategories([]string) ProductAkeneoBuilder
	AddVisibility(string) ProductAkeneoBuilder
	Build() ProductAkeneo
}

func (pb *ProductAkeneo) AddLocale(locale *string) ProductAkeneoBuilder {
	pb.Locale = locale
	return pb
}

func (pb *ProductAkeneo) AddScope(scope *string) ProductAkeneoBuilder {
	pb.Scope = scope
	return pb
}

func (pb *ProductAkeneo) AddIdentifier(identifier string) ProductAkeneoBuilder {
	pb.Identifier = identifier
	return pb
}

func (pb *ProductAkeneo) AddName(name string) ProductAkeneoBuilder {
	pb.Values["name"] = Values{NewValueBuilder().AddLocale(nil).AddScope(pb.Scope).AddData(name, "value").Build()}
	return pb
}

func (pb *ProductAkeneo) AddSku(sku string) ProductAkeneoBuilder {
	pb.Values["sku"] = Values{NewValueBuilder().AddLocale(pb.Locale).AddScope(pb.Scope).AddData(sku, "identifier").Build()}
	return pb
}

func (pb *ProductAkeneo) AddPrice(price int) ProductAkeneoBuilder {
	pb.Values["price"] = Values{NewValueBuilder().AddLocale(nil).AddScope(pb.Scope).AddData(price, "pim_catalog_price").Build()}
	return pb
}

func (pb *ProductAkeneo) AddStatus(status bool) ProductAkeneoBuilder {
	pb.Enabled = StatusBool(status)
	return pb
}

func (pb *ProductAkeneo) AddVisibility(visibility string) ProductAkeneoBuilder {
	pb.Values["visibility"] = Values{NewValueBuilder().AddLocale(nil).AddScope(pb.Scope).AddData(visibility, "pim_catalog_simpleselect").Build()}
	return pb
}

func (pb *ProductAkeneo) AddIsInStock(status string) ProductAkeneoBuilder {
	pb.Values["quantity_and_stock_status"] = Values{NewValueBuilder().AddLocale(nil).AddScope(pb.Scope).AddData(status, "pim_catalog_simpleselect").Build()}
	return pb
}

func (pb *ProductAkeneo) AddValues(attributeCode string, value interface{}, akeneoType string) ProductAkeneoBuilder {
	pb.Values[attributeCode] = Values{NewValueBuilder().AddLocale(nil).AddScope(pb.Scope).AddData(value, akeneoType).Build()}
	return pb
}

func (pb *ProductAkeneo) AddCategories(categories []string) ProductAkeneoBuilder {
	pb.Categories = categories
	return pb
}
func (pb *ProductAkeneo) AddFamily(family string) ProductAkeneoBuilder {
	pb.Family = family
	return pb
}

func (pb *ProductAkeneo) Build() ProductAkeneo {
	return ProductAkeneo{
		Identifier: pb.Identifier,
		Categories: pb.Categories,
		Enabled:    pb.Enabled,
		Family:     pb.Family,
		Values:     pb.Values,
		Locale:     pb.Locale,
		Scope:      pb.Scope,
	}
}

func NewProductAkeneoBuilder() ProductAkeneoBuilder {
	return &ProductAkeneo{
		Values: map[string]Values{},
	}
}

/*Value*/
type IValueBuilder interface {
	AddLocale(*string) IValueBuilder
	AddScope(*string) IValueBuilder
	AddData(any, string) IValueBuilder
	Build() Value
}
type Value struct {
	Locale *string `json:"locale"`
	Scope  *string `json:"scope"`
	Data   any     `json:"data"`
}
type ValuePimPrice struct {
	Amount   any    `json:"amount"`
	Currency string `json:"currency"`
}

type Values []Value

func (vb *Value) AddLocale(locale *string) IValueBuilder {
	vb.Locale = locale
	return vb
}

func (vb *Value) AddScope(scope *string) IValueBuilder {
	vb.Scope = scope
	return vb
}

func (vb *Value) AddData(data any, typeData string) IValueBuilder {
	switch {
	case typeData == "pim_catalog_price" || typeData == "pim_catalog_price_collection":
		vb.Data = []ValuePimPrice{{Amount: data, Currency: "CLP"}}
	case typeData == "pim_catalog_image":
		vb.Data = nil

	case typeData == "pim_catalog_boolean":
		if data == "0" {
			vb.Data = false
		} else {
			vb.Data = true
		}

	default:
		vb.Data = data

	}

	return vb
}
func (vb *Value) Build() Value {
	return Value{
		Locale: vb.Locale,
		Scope:  vb.Scope,
		Data:   vb.Data,
	}
}
func NewValueBuilder() IValueBuilder {
	return &Value{}
}
