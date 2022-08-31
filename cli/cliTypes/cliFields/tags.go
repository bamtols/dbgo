package cliFields

type (
	Tag     string
	TagList []Tag
	TagMap  map[Tag]bool
)

const (
	TagPrimaryKey Tag = "primaryKey"
	TagOmit       Tag = "omit"
	TagUnique     Tag = "unique"
	TagNullable   Tag = "nullable"
	TagCreatedAt  Tag = "createdAt"
	TagUpdatedAt  Tag = "updatedAt"
	TagDeletedAt  Tag = "deletedAt"
)

var (
	TagListAll = TagList{
		TagPrimaryKey,
		TagOmit,
		TagUnique,
		TagNullable,
		TagCreatedAt,
		TagUpdatedAt,
		TagDeletedAt,
	}
)

// Valid 유효한 값인지 확인
func (x *Tag) Valid() bool {
	for _, tag := range TagListAll {
		if *x == tag {
			return true
		}
	}
	return false
}

// Validate 중복제거 및 유효한 값인지 확인
func (x *TagList) Validate() TagList {
	m := make(TagMap)
	for _, tag := range *x {
		if !tag.Valid() {
			continue
		}
		m[tag] = true
	}

	*x = make(TagList, 0)
	for tag := range m {
		*x = append(*x, tag)
	}

	return *x
}
