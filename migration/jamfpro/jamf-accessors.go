package jamfpro

/* Departments */
func (d *Department) SetId(v string) {
	d.Id = &v
}

func (d *Department) SetName(v string) {
	d.Name = &v
}

func (d *Department) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}

func (d *Department) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

/* Categories */
func (d *Category) SetId(v string) {
	d.Id = &v
}

func (d *Category) SetName(v string) {
	d.Name = &v
}

func (d *Category) SetPriority(v int) {
	d.Priority = &v
}

func (d *Category) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}

func (d *Category) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

func (d *Category) GetPriority() int {
	if d == nil || d.Priority == nil {
		return 0
	}
	return *d.Priority
}

/* Buildings */
func (d *Building) SetId(v string) {
	d.Id = &v
}

func (d *Building) SetName(v string) {
	d.Name = &v
}

func (d *Building) SetStreetAddress1(v string) {
	d.StreetAddress1 = &v
}

func (d *Building) SetStreetAddress2(v string) {
	d.StreetAddress2 = &v
}

func (d *Building) SetCity(v string) {
	d.City = &v
}

func (d *Building) SetStateProvince(v string) {
	d.StateProvince = &v
}

func (d *Building) SetZipPostalCode(v string) {
	d.ZipPostalCode = &v
}

func (d *Building) SetCountry(v string) {
	d.Country = &v
}

func (d *Building) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}

func (d *Building) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

func (d *Building) GetStreetAddress1() string {
	if d == nil || d.StreetAddress1 == nil {
		return ""
	}
	return *d.StreetAddress1
}

func (d *Building) GetStreetAddress2() string {
	if d == nil || d.StreetAddress2 == nil {
		return ""
	}
	return *d.StreetAddress2
}

func (d *Building) GetCity() string {
	if d == nil || d.City == nil {
		return ""
	}
	return *d.City
}

func (d *Building) GetStateProvince() string {
	if d == nil || d.StateProvince == nil {
		return ""
	}
	return *d.StateProvince
}

func (d *Building) GetZipPostalCode() string {
	if d == nil || d.ZipPostalCode == nil {
		return ""
	}
	return *d.ZipPostalCode
}

func (d *Building) GetCountry() string {
	if d == nil || d.Country == nil {
		return ""
	}
	return *d.Country
}
