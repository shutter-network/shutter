package observe

// World describes the observable outside world, i.e. the shutter and main chain instance.
type World struct {
	Shutter   *Shutter
	MainChain *MainChain
}
