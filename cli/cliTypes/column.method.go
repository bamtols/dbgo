package cliTypes

func (x *Column) HasTag(v ColumnTag) bool {
	if !x.hasTagData() {
		return false
	}

	for _, tag := range x.Tags {
		if tag == v {
			return true
		}
	}
	return false
}

func (x *Column) HasTagOne(v ...ColumnTag) bool {
	if !x.hasTagData() {
		return false
	}

	for _, tag := range x.Tags {
		for _, hasTag := range v {
			if tag == hasTag {
				return true
			}
		}
	}

	return false
}

func (x *Column) HasTagAll(v ...ColumnTag) bool {
	if !x.hasTagData() {
		return false
	}

	for _, hasTag := range v {
		if !x.HasTag(hasTag) {
			return false
		}
	}

	return true
}

func (x *Column) hasTagData() bool {
	if x.Tags == nil || len(x.Tags) == 0 {
		return false
	}
	return true
}
