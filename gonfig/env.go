package gonfig

var EnvSets = map[string]func(string, string){
	"sdev0": SetSdev0,
	"sdev":  SetSdev,
	"dev":   SetDev,
	"dev2":  SetDev2,
	"ops":   SetOps,
}

func Initialize(env, appName string) {

	fn, ok := EnvSets[env]
	if !ok {
		return
	}

	fn(env, appName)

}
