package edmx

type Edmx struct {
	Version string `xml:"Version,attr"`
	DataServices DataServices `xml:"DataServices"`
}

type DataServices struct {
	Schemas[]Schema `xml:"Schema"`
}

type Schema struct {
	Namespace string `xml:"Namespace,attr"`
	EntityTypes[]EntityType `xml:"EntityType"`
}

type EntityType struct {
	Name string `xml:"Name,attr"`
	Key Key `xml:"Key"`
	Properties Property `xml:"Property"`
}

type Key struct {
	PropertyRefs []PropertyRef `xml:"PropertyRef"`
}

type Property struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
	IsLanguageDependent bool `xml:"IsLanguageDependent"`
	Nullable bool `xml:"Nullable"`
}

type PropertyRef struct {
	Name string `xml:"Name,attr"`
}

type NavigationProperty struct {
	Name string `xml:"Name,attr"`
	Relationship string `xml:"Relationship,attr"`
	FromRole string `xml:"FromRole,attr"`
	ToRole string `xml:"ToRole,attr"`
	IsUnique bool `xml:"IsUnique"`
	Nullable bool `xml:"Nullable"`
}
