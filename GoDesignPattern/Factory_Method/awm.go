package factorymethod

type awm struct {
	Gun
}

func newAwm() iGun {
	return &awm{
		Gun: Gun{
			name:  "awm",
			power: 10,
		},
	}
}
