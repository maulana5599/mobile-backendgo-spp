package entity

type TahunAjaran struct {
	IdTahunAjaran int
	TAjar         string
	Status        string
	Tampilan      string
}

func (TahunAjaran) TableName() string {
	return "tahun_ajaran"
}
