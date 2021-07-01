package gonfig

var EnvSets = map[string]func(){
	"sdev0": SetSdev0,
	"ops":   SetOps,
}

func Initialize(env string) {

	fn, ok := EnvSets[env]
	if !ok {
		return
	}

	fn()

}
